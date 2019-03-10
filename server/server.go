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
	router *mux.Router
	config *Infra.ServerConfiguration
	auth   *Auth
}

// CreateServer create a new http server
func CreateServer(c *Infra.Configuration) *Server {
	s := Server{router: mux.NewRouter(), config: c.Server}
	s.auth = &Auth{c.Auth.Secret}
	return &s
}

// Start start http server
func (s *Server) Start() {
	log.Printf("Listening on port %d", s.config.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", s.config.Port), handlers.LoggingHandler(os.Stdout, s.router))
	if err != nil {
		log.Fatal("http.ListenAndServe: ", err)
	}
}

// NewSubrouter create a new subrouter
func (s *Server) NewSubrouter(path string) *mux.Router {
	return s.router.PathPrefix(path).Subrouter()
}

func (s *Server) GetAuth() *Auth {
	return s.auth
}
