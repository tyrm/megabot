package web

import (
	"context"
	"testing"
	"time"
)

func TestServer3_StartStop(t *testing.T) {
	server, err := New3(context.Background())
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
