package server

import (
	"log"
	"net/http"
	"os"
)

const HostUrlKey string = "HOST_URL"
const HostPortKey string = "HOST_PORT"

type Server interface {
	Serve()
}

type server struct {
	routes map[string]func(w http.ResponseWriter, r *http.Request)
}

func NewServer(routes map[string]func(w http.ResponseWriter, r *http.Request)) Server {
	return server{routes}
}

func (s server) Serve() {
	for url, handler := range s.routes {
		http.HandleFunc(url, handler)
	}

	log.Fatal(http.ListenAndServe(":" + os.Getenv(HostPortKey), nil))
}
