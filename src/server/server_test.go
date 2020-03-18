package server_test

import (
	"dockerWebhook/src/server"
	"net/http"
	"testing"
)

func TestNewServer(t *testing.T) {
	routes := map[string]func(http.ResponseWriter, *http.Request){
		"/": func(w http.ResponseWriter, r *http.Request) {},
	}

	assertServerInterface(server.NewServer(routes), t)
}

func assertServerInterface(s server.Server, t *testing.T)  {
	_, ok := s.(server.Server)

	if ok != true {
		t.Error("expecting Server interface")
	}
}
