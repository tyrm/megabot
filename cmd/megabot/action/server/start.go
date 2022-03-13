package server

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tyrm/go-util"
	"github.com/tyrm/megabot/cmd/megabot/action"
	"github.com/tyrm/megabot/internal/config"
	"github.com/tyrm/megabot/internal/db/bun"
	"github.com/tyrm/megabot/internal/graphql"
	"github.com/tyrm/megabot/internal/jwt"
	"github.com/tyrm/megabot/internal/kv/redis"
	"github.com/tyrm/megabot/internal/language"
	"github.com/tyrm/megabot/internal/web"
	"github.com/tyrm/megabot/internal/webapp"
	"os"
	"os/signal"
	"reflect"
	"syscall"
)

// Start starts the server
var Start action.Action = func(ctx context.Context) error {
	logrus.Infof("starting")
	dbClient, err := bun.New(ctx)
	if err != nil {
		logrus.Errorf("db: %s", err.Error())
		return err
	}
	defer func() {
		err := dbClient.Close(ctx)
		if err != nil {
			logrus.Errorf("closing db: %s", err.Error())
		}
	}()

	redisClient, err := redis.New(ctx)
	if err != nil {
		logrus.Errorf("redis: %s", err.Error())
		return err
	}
	defer func() {
		err := redisClient.Close(ctx)
		if err != nil {
			logrus.Errorf("closing redis: %s", err.Error())
		}
	}()

	jwtModule, err := jwt.New(dbClient, redisClient)
	if err != nil {
		logrus.Errorf("jwt: %s", err.Error())
		return err
	}
	defer func() {
		err := jwtModule.Close()
		if err != nil {
			logrus.Errorf("closing jwt: %s", err.Error())
		}
	}()

	languageMod, err := language.New()
	if err != nil {
		logrus.Errorf("language: %s", err.Error())
		return err
	}

	var webServers []web.Server
	if viper.GetBool(config.Keys.ServerHTTP2) {
		logrus.Debugf("creating http2 server")
		server2, err := web.New2(ctx, dbClient)
		if err != nil {
			logrus.Errorf("http2 server: %s", err.Error())
			return err
		}
		webServers = append(webServers, server2)
	}
	logrus.Infof("%v", viper.GetBool(config.Keys.ServerHTTP3))
	if viper.GetBool(config.Keys.ServerHTTP3) {
		logrus.Debugf("creating http3 server")
		server3, err := web.New3(ctx, dbClient)
		if err != nil {
			logrus.Errorf("http3 server: %s", err.Error())
			return err
		}
		webServers = append(webServers, server3)
	}

	var webModules []web.Module
	if util.ContainsString(viper.GetStringSlice(config.Keys.ServerRoles), config.ServerRoleGraphQL) {
		logrus.Infof("adding graphql module")
		webMod := graphql.New(dbClient, jwtModule)
		webModules = append(webModules, webMod)
	}
	if util.ContainsString(viper.GetStringSlice(config.Keys.ServerRoles), config.ServerRoleWebapp) {
		logrus.Infof("adding webapp module")
		webMod, err := webapp.New(ctx, dbClient, redisClient, languageMod)
		if err != nil {
			logrus.Errorf("webapp module: %s", err.Error())
			return err
		}
		webModules = append(webModules, webMod)
	}

	for _, server := range webServers {
		for _, mod := range webModules {
			err := mod.Route(server)
			if err != nil {
				logrus.Errorf("loading %s module: %s", mod.Name(), err.Error())
				return err
			}
		}
	}

	// ** start application **
	errChan := make(chan error)

	// start web server
	logrus.Infof("starting web server")
	for _, server := range webServers {
		go func(s web.Server, errChan chan error) {
			logrus.Debugf("starting %s", reflect.TypeOf(s).String())
			err := s.Start()
			if err != nil {
				errChan <- fmt.Errorf("web server: %s", err.Error())
			}
		}(server, errChan)
	}
	defer func() {
		for _, server := range webServers {
			err := server.Stop(ctx)
			if err != nil {
				logrus.Errorf("stopping web server: %s", err.Error())
			}
		}
	}()

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	nch := make(chan os.Signal)
	signal.Notify(nch, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-nch:
		logrus.Infof("got sig: %s", sig)
	case err := <-errChan:
		logrus.Fatal(err.Error())
	}

	logrus.Infof("done")
	return nil
}
