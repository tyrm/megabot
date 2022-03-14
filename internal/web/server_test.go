package web

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tables := []struct {
		new      func(ctx context.Context) (Server, error)
		t        interface{}
		validate func(t *testing.T, s Server)
	}{
		{New2, &Server2{}, testNewValidateServer2},
		{New3, &Server3{}, testNewValidateServer3},
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
