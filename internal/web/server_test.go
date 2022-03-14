package web

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/lucas-clemente/quic-go"
	"github.com/lucas-clemente/quic-go/http3"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/tyrm/megabot/internal/config"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
	"time"
)

const (
	testPath         = "/test"
	testPathPrefix   = "/prefix"
	testResponseBody = "hello world!"
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
				t.Errorf("unexpected error initializing %s server: %s", expectedType.String(), err.Error())
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
	server := s.(*Server2)

	if server.router == nil {
		t.Errorf("router is nil")
	}
	if server.srv == nil {
		t.Errorf("server is nil")
	}
}

func testNewValidateServer3(t *testing.T, s Server) {
	server := s.(*Server3)

	if server.router == nil {
		t.Errorf("router is nil")
	}
	if server.srv == nil {
		t.Errorf("server is nil")
	}
}

func TestServer_HandleFunc(t *testing.T) {
	viper.Reset()
	err := config.Init(&pflag.FlagSet{})
	if err != nil {
		t.Errorf("init: %s", err.Error())
		return
	}

	tables := []struct {
		t        interface{}
		new      func(ctx context.Context) (Server, error)
		validate func(t *testing.T, path, body string)
	}{
		{&Server2{}, New2, testServerValidateServer2},
		{&Server3{}, New3, testServerValidateServer3},
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

			server.HandleFunc(testPath, testHTTPHandler).Methods("GET")

			go func(s Server, st string) {
				t.Logf("starting %s server", st)
				err := s.Start()
				t.Logf("%s server stopped: %s", st, err)
			}(server, expectedType.String())
			time.Sleep(100 * time.Millisecond)

			table.validate(t, testPath, testResponseBody)

			err = server.Stop(context.Background())
			if err != nil {
				t.Errorf("unexpected error stopping %s server: %s", expectedType.String(), err.Error())
			}
		})
	}
}

func TestServer_PathPrefix(t *testing.T) {
	viper.Reset()
	err := config.Init(&pflag.FlagSet{})
	if err != nil {
		t.Errorf("init: %s", err.Error())
		return
	}

	tables := []struct {
		t        interface{}
		new      func(ctx context.Context) (Server, error)
		validate func(t *testing.T, path, body string)
	}{
		{&Server2{}, New2, testServerValidateServer2},
		{&Server3{}, New3, testServerValidateServer3},
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

			server.PathPrefix(testPathPrefix).Subrouter().HandleFunc(testPath, testHTTPHandler).Methods("GET")

			go func(s Server, st string) {
				t.Logf("starting %s server", st)
				err := s.Start()
				t.Logf("%s server stopped: %s", st, err)
			}(server, expectedType.String())
			time.Sleep(100 * time.Millisecond)

			table.validate(t, testPathPrefix+testPath, testResponseBody)

			err = server.Stop(context.Background())
			if err != nil {
				t.Errorf("unexpected error stopping %s server: %s", expectedType.String(), err.Error())
			}
		})
	}
}

func testServerValidateServer2(t *testing.T, path, body string) {
	resp, err := http.Get("http://localhost:5000" + path)
	if err != nil {
		t.Errorf("http 2 request error: %s", err.Error())
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("http 2 request error: %s", err.Error())
		return
	}
	if string(b) != body {
		t.Errorf("http 2 request got wrong body, got: '%s' , want: '%s'", string(b), body)
		return
	}
}

func testServerValidateServer3(t *testing.T, path, body string) {
	client := &http.Client{
		Transport: &http3.RoundTripper{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			QuicConfig: &quic.Config{},
		},
	}

	resp, err := client.Get("https://localhost:5000" + path)
	if err != nil {
		t.Errorf("http 3 request error: %s", err.Error())
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("http 3 request error: %s", err.Error())
		return
	}
	if string(b) != body {
		t.Errorf("http 3 request got wrong body, got: '%s' , want: '%s'", string(b), body)
		return
	}
}

func TestServer_StartStop(t *testing.T) {
	viper.Reset()
	err := config.Init(&pflag.FlagSet{})
	if err != nil {
		t.Errorf("init: %s", err.Error())
		return
	}

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
			go func(s Server, e chan error) {
				t.Logf("starting %s server", expectedType.String())
				err := server.Start()
				t.Logf("%s server stopped", expectedType.String())
				e <- err
			}(server, errChan)

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

func testHTTPHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}
