package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	Infra "../infra"
)

// Server http server
type Server struct {
	Router *mux.Router
	config *Infra.ServerConfiguration
}

// CreateServer create a new http server
func CreateServer(c *Infra.ServerConfiguration) *Server {
	s := Server{Router: mux.NewRouter(), config: c}
	return &s
}

// Start start http server
func (s *Server) Start() {
	log.Printf("Listening on port %d", s.config.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", s.config.Port), handlers.LoggingHandler(os.Stdout, s.Router))
	if err != nil {
		log.Fatal("http.ListenAndServe: ", err)
	}
}

// NewSubrouter create a new subrouter
func (s *Server) NewSubrouter(path string) *mux.Router {
	return s.Router.PathPrefix(path).Subrouter()
}
