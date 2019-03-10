package main

import (
	"fmt"
	"log"

	Controllers "./controllers"
	Infra "./infra"
	Server "./server"
	Services "./services"
)

// App api
type App struct {
	server  *Server.Server
	session *Infra.Session
	config  *Infra.Configuration
}

// Init init application
func (a *App) Init() {
	a.config = Infra.GetConfigurations()
	var err error
	a.session, err = Infra.NewSession(a.config.Mongo)
	if err != nil {
		log.Fatalf("unable to connect to mongodb, %s", err)
	}
	a.server = Server.CreateServer(a.config)

	us := Services.NewUserService(a.session.Copy())

	Controllers.CreateUserController(us, a.server.NewSubrouter("api/user"), a.server.GetAuth())
}

// Run run app
func (a *App) Run() {
	fmt.Println("Running app...")
	defer a.session.Close()
	a.server.Start()
}
