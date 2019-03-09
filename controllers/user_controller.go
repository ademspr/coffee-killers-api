package routes

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	Entities "../entities"
	Server "../server"
	Services "../services"
	"github.com/gorilla/mux"
)

type userController struct {
	userService *Services.UserService
}

func CreateUserController(us *Services.UserService, router *mux.Router) *mux.Router {
	userController := userController{us}

	router.HandleFunc("/", userController.createUserHandler).Methods("PUT")
	router.HandleFunc("/{username}", userController.getUserHandler).Methods("GET")

	return router
}

func (ur *userController) createUserHandler(w http.ResponseWriter, r *http.Request) {
	user, err := decodeUser(r)
	if err != nil {
		Server.Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = ur.userService.Create(&user)
	if err != nil {
		Server.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	Server.JSON(w, http.StatusOK, err)
}

func (ur *userController) getUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars)
	username := vars["username"]

	user, err := ur.userService.GetByUsername(username)
	if err != nil {
		Server.Error(w, http.StatusNotFound, err.Error())
		return
	}

	Server.JSON(w, http.StatusOK, user)
}

func decodeUser(r *http.Request) (Entities.User, error) {
	var u Entities.User
	if r.Body == nil {
		return u, errors.New("no request body")
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&u)
	return u, err
}
