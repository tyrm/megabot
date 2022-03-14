package web

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	tables := []struct {
		t        interface{}
		new      func(ctx context.Context) (Server, error)
		validate func(t *testing.T, s Server)
	}{
		{&Server2{}, New2, testNewValidateServer2},
		{&Server3{}, New3, testNewValidateServer3},
	}

	for i, table := range tables {
		i := i
		table := table
		expectedType := reflect.TypeOf(table.t)

		name := fmt.Sprintf("[%d] Running new on %s", i, expectedType.String())
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			server, err := table.new(context.Background())
			if err != nil {
				t.Errorf("unexpected error initializing http 2 server: %s", err.Error())
				return
			}
			if server == nil {
				t.Errorf("server is nil")
				return
			}
			serverType := reflect.TypeOf(server)
			if serverType != expectedType {
				t.Errorf("server is wrong type, got %s, want %s", serverType, expectedType)
				return
			}

			table.validate(t, server)
		})
	}
}

func testNewValidateServer2(t *testing.T, s Server) {
	server2 := s.(*Server2)

	if server2.router == nil {
		t.Errorf("router is nil")
	}
	if server2.srv == nil {
		t.Errorf("server is nil")
	}
}

func testNewValidateServer3(t *testing.T, s Server) {
	server2 := s.(*Server3)

	if server2.router == nil {
		t.Errorf("router is nil")
	}
	if server2.srv == nil {
		t.Errorf("server is nil")
	}
}

func TestServer_StartStop(t *testing.T) {
	tables := []struct {
		t               interface{}
		new             func(ctx context.Context) (Server, error)
		stopSuccessResp string
	}{
		{&Server2{}, New2, "http: Server closed"},
		{&Server3{}, New3, "server closed"},
	}

	for i, table := range tables {
		i := i
		table := table
		expectedType := reflect.TypeOf(table.t)

		name := fmt.Sprintf("[%d] Running new on %s", i, expectedType.String())
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			server, err := table.new(context.Background())
			if err != nil {
				t.Errorf("unexpected error initializing %s server: %s", expectedType.String(), err.Error())
				return
			}

			errChan := make(chan error)
			doneChan := make(chan bool)
			go func(s Server, e chan error, d chan bool) {
				t.Logf("starting %s server", expectedType.String())
				err := server.Start()
				t.Logf("%s server stopped", expectedType.String())
				e <- err
			}(server, errChan, doneChan)

			time.Sleep(100 * time.Millisecond)

			err = server.Stop(context.Background())
			if err != nil {
				t.Errorf("unexpected error stopping %s server: %s", expectedType.String(), err.Error())
				return
			}

			select {
			case err := <-errChan:
				if err.Error() == table.stopSuccessResp {
					t.Logf("%s server close successfully", expectedType.String())
				} else {
					t.Errorf("error closing %s server: %s", expectedType.String(), err.Error())
				}
			case <-time.After(10 * time.Second):
				t.Errorf("timeout waiting for %s server to close", expectedType.String())
			}
		})
	}
}
