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
	"github.com/tyrm/megabot/internal/web"
	"os"
	"os/signal"
	"syscall"
)

var Start action.Action = func(ctx context.Context) error {
	logrus.Infof("starting")
	dbClient, err := bun.NewClient(ctx)
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

	jwtModule, err := jwt.New(redisClient)
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

	webServer, err := web.NewServer(ctx, dbClient)
	if err != nil {
		logrus.Errorf("web server: %s", err.Error())
		return err
	}

	var modGraphQL web.Module
	if util.ContainsString(viper.GetStringSlice(config.Keys.ServerRoles), config.ServerRoleGraphql) {
		logrus.Infof("adding graphql")
		modGraphQL = graphql.New(dbClient, jwtModule)
		err := modGraphQL.Route(webServer)
		if err != nil {
			logrus.Errorf("graphql module: %s", err.Error())
			return err
		}
	}

	// ** start application **
	errChan := make(chan error)

	// start web server
	logrus.Infof("starting web server")
	go func(errChan chan error) {
		err := webServer.Start()
		if err != nil {
			errChan <- fmt.Errorf("web server: %s", err.Error())
		}
	}(errChan)
	defer func() {
		err := webServer.Stop(ctx)
		if err != nil {
			logrus.Errorf("stopping web server: %s", err.Error())
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
