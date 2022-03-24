package server

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tyrm/go-util"
	"github.com/tyrm/megabot/cmd/megabot/action"
	"github.com/tyrm/megabot/internal/boostrapcheatsheet"
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
	l := logger.WithField("func", "Start")

	l.Infof("starting")
	dbClient, err := bun.New(ctx)
	if err != nil {
		l.Errorf("db: %s", err.Error())
		return err
	}
	defer func() {
		err := dbClient.Close(ctx)
		if err != nil {
			l.Errorf("closing db: %s", err.Error())
		}
	}()

	redisClient, err := redis.New(ctx)
	if err != nil {
		l.Errorf("redis: %s", err.Error())
		return err
	}
	defer func() {
		err := redisClient.Close(ctx)
		if err != nil {
			l.Errorf("closing redis: %s", err.Error())
		}
	}()

	jwtModule, err := jwt.New(dbClient, redisClient)
	if err != nil {
		l.Errorf("jwt: %s", err.Error())
		return err
	}
	defer func() {
		err := jwtModule.Close()
		if err != nil {
			l.Errorf("closing jwt: %s", err.Error())
		}
	}()

	languageMod, err := language.New()
	if err != nil {
		l.Errorf("language: %s", err.Error())
		return err
	}

	// create web servers
	var webServers []web.Server
	if viper.GetBool(config.Keys.ServerHTTP2) {
		l.Debugf("creating http2 server")
		server2, err := web.New2(ctx)
		if err != nil {
			l.Errorf("http2 server: %s", err.Error())
			return err
		}
		webServers = append(webServers, server2)
	}
	l.Infof("%v", viper.GetBool(config.Keys.ServerHTTP3))
	if viper.GetBool(config.Keys.ServerHTTP3) {
		l.Debugf("creating http3 server")
		server3, err := web.New3(ctx)
		if err != nil {
			l.Errorf("http3 server: %s", err.Error())
			return err
		}
		webServers = append(webServers, server3)
	}

	// create web modules
	var webModules []web.Module
	if util.ContainsString(viper.GetStringSlice(config.Keys.ServerRoles), config.ServerRoleBootstrap) {
		l.Infof("adding bootstrap test module")
		webMod, err := boostrapcheatsheet.New(ctx)
		if err != nil {
			logrus.Errorf("bootstrap test module: %s", err.Error())
			return err
		}
		webModules = append(webModules, webMod)
	}
	if util.ContainsString(viper.GetStringSlice(config.Keys.ServerRoles), config.ServerRoleGraphQL) {
		l.Infof("adding graphql module")
		webMod := graphql.New(dbClient, jwtModule)
		webModules = append(webModules, webMod)
	}
	if util.ContainsString(viper.GetStringSlice(config.Keys.ServerRoles), config.ServerRoleWebapp) {
		l.Infof("adding webapp module")
		webMod, err := webapp.New(ctx, dbClient, redisClient, languageMod)
		if err != nil {
			logrus.Errorf("webapp module: %s", err.Error())
			return err
		}
		webModules = append(webModules, webMod)
	}

	// add modules to servers
	for _, server := range webServers {
		for _, mod := range webModules {
			err := mod.Route(server)
			if err != nil {
				l.Errorf("loading %s module: %s", mod.Name(), err.Error())
				return err
			}
		}
	}

	// ** start application **
	errChan := make(chan error)

	// start web server
	l.Infof("starting web server")
	for _, server := range webServers {
		go func(s web.Server, errChan chan error) {
			l.Debugf("starting %s", reflect.TypeOf(s).String())
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
				l.Errorf("stopping web server: %s", err.Error())
			}
		}
	}()

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	nch := make(chan os.Signal)
	signal.Notify(nch, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-nch:
		l.Infof("got sig: %s", sig)
	case err := <-errChan:
		l.Fatal(err.Error())
	}

	l.Infof("done")
	return nil
}
