package routes

import (
	"encoding/json"
	"errors"
	"net/http"

	DTOs "../dataobjects"
	Entities "../entities"
	Server "../server"
	Services "../services"
	"github.com/gorilla/mux"
)

type userController struct {
	userService *Services.UserService
	auth        *Server.Auth
}

func CreateUserController(us *Services.UserService, router *mux.Router, a *Server.Auth) *mux.Router {
	userController := userController{us, a}

	router.HandleFunc("/", userController.createUserHandler).Methods("PUT")
	router.HandleFunc("/{username}", userController.getUserHandler).Methods("GET")

	return router
}

func (ur *userController) createUserHandler(w http.ResponseWriter, r *http.Request) {
	err, user := decodeUser(r)
	if err != nil {
		Server.Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = ur.userService.CreateUser(&user)
	if err != nil {
		Server.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	Server.JSON(w, http.StatusOK, err)
}

func (ur *userController) profileHandler(w http.ResponseWriter, r *http.Request) {
	claim, ok := r.Context().Value(Server.ContextKeyAuthtoken).(Server.Claims)
	if !ok {
		Server.Error(w, http.StatusBadRequest, "no context")
		return
	}
	username := claim.Username

	user, err := ur.userService.GetByUsername(username)
	if err != nil {
		Server.Error(w, http.StatusNotFound, err.Error())
		return
	}

	Server.JSON(w, http.StatusOK, user)
}

func (ur *userController) getUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	user, err := ur.userService.GetByUsername(username)
	if err != nil {
		Server.Error(w, http.StatusNotFound, err.Error())
		return
	}

	Server.JSON(w, http.StatusOK, user)
}

func (ur *userController) loginHandler(w http.ResponseWriter, r *http.Request) {
	err, credentials := decodeCredentials(r)
	if err != nil {
		Server.Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	var user Entities.User
	user, err = ur.userService.Login(credentials)
	if err == nil {
		cookie := ur.auth.NewCookie(user)
		Server.JSONWithCookie(w, http.StatusOK, user, cookie)
	} else {
		Server.Error(w, http.StatusInternalServerError, "Incorrect password")
	}
}

func decodeUser(r *http.Request) (error, Entities.User) {
	var u Entities.User
	if r.Body == nil {
		return errors.New("no request body"), u
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&u)
	return err, u
}

func decodeCredentials(r *http.Request) (error, DTOs.UserCredentials) {
	var c DTOs.UserCredentials
	if r.Body == nil {
		return errors.New("no request body"), c
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&c)
	return err, c
}
