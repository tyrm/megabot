package web

import (
	"context"
	"github.com/spf13/viper"
	"github.com/tyrm/megabot/internal/config"
	"github.com/tyrm/megabot/internal/db"
	"github.com/tyrm/megabot/internal/db/bun"
	"testing"
	"time"
)

func TestNew2(t *testing.T) {
	dbClient, err := testNewTestDBClient()
	if err != nil {
		t.Errorf("init: %s", err.Error())
		return
	}

	server, err := New2(context.Background(), dbClient)
	if err != nil {
		t.Errorf("unexpected error initializing http 2 server: %s", err.Error())
		return
	}
	if server == nil {
		t.Errorf("http 2 server is nil")
		return
	}

	server2 := server.(*Server2)

	if server2.router == nil {
		t.Errorf("router is nil")
	}
	if server2.srv == nil {
		t.Errorf("server is nil")
	}
}

func TestServer2_StartStop(t *testing.T) {
	dbClient, err := testNewTestDBClient()
	if err != nil {
		t.Errorf("unexpected error initializing pg options: %s", err.Error())
		return
	}

	server, err := New2(context.Background(), dbClient)
	if err != nil {
		t.Errorf("unexpected error initializing http 2 server: %s", err.Error())
		return
	}

	errChan := make(chan error)
	doneChan := make(chan bool)
	go func(s Server, e chan error, d chan bool) {
		t.Logf("starting http 2 server")
		err := server.Start()
		e <- err
	}(server, errChan, doneChan)

	time.Sleep(100 * time.Millisecond)

	err = server.Stop(context.Background())
	if err != nil {
		t.Errorf("unexpected error stopping http 2 server: %s", err.Error())
		return
	}

	select {
	case err := <-errChan:
		if err.Error() == "http: Server closed" {
			t.Logf("http 2 server close successfully")
		} else {
			t.Errorf("http 2 server: %s", err.Error())
		}
	case <-time.After(3 * time.Second):
		t.Errorf("timeout waiting for http 2 server to close")
	}
}

func testNewTestDBClient() (db.DB, error) {
	viper.Reset()

	viper.Set(config.Keys.DbType, "sqlite")
	viper.Set(config.Keys.DbAddress, ":memory:")

	client, err := bun.New(context.Background())
	if err != nil {
		return nil, err
	}

	err = client.DoMigration(context.Background())
	if err != nil {
		return nil, err
	}

	err = client.LoadTestData(context.Background())
	if err != nil {
		return nil, err
	}

	return client, nil
}
