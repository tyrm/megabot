package web

import (
	"context"
	"testing"
	"time"
)

func TestNew3(t *testing.T) {
	dbClient, err := testNewTestDBClient()
	if err != nil {
		t.Errorf("init: %s", err.Error())
		return
	}

	server, err := New3(context.Background(), dbClient)
	if err != nil {
		t.Errorf("unexpected error initializing http 3 server: %s", err.Error())
		return
	}
	if server == nil {
		t.Errorf("http 3 server is nil")
		return
	}

	server3 := server.(*Server3)

	if server3.router == nil {
		t.Errorf("router is nil")
	}
	if server3.srv == nil {
		t.Errorf("server is nil")
	}
}

func TestServer3_StartStop(t *testing.T) {
	dbClient, err := testNewTestDBClient()
	if err != nil {
		t.Errorf("unexpected error initializing pg options: %s", err.Error())
		return
	}

	server, err := New3(context.Background(), dbClient)
	if err != nil {
		t.Errorf("unexpected error initializing http 3 server: %s", err.Error())
		return
	}

	errChan := make(chan error)
	doneChan := make(chan bool)
	go func(s Server, e chan error, d chan bool) {
		t.Logf("starting http 3 server")
		err := server.Start()
		t.Logf("http 3 server stopped")
		e <- err
	}(server, errChan, doneChan)

	time.Sleep(100 * time.Millisecond)

	err = server.Stop(context.Background())
	if err != nil {
		t.Errorf("unexpected error stopping http 3 server: %s", err.Error())
		return
	}

	select {
	case err := <-errChan:
		if err.Error() == "server closed" {
			t.Logf("http 3 server close successfully")
		} else {
			t.Errorf("http 3 server: %s", err.Error())
		}
	case <-time.After(10 * time.Second):
		t.Errorf("timeout waiting for http 3 server to close")
	}
}
