package main

import (
	"encoding/json"
	"log"
	"net/http"

	Entities "./entities"
	Infra "./infra"
	Services "./services"
	"github.com/gorilla/mux"
)

var people []Entities.Person

func main() {
	router := mux.NewRouter()
	people = append(people, Entities.Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Entities.Address{City: "City X", State: "State X"}})
	people = append(people, Entities.Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Entities.Address{City: "City Z", State: "State Y"}})

	personService := Services.NewPersonService(people)

	router.HandleFunc("/config", GetConfigurations).Methods("GET")
	router.HandleFunc("/people", personService.GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", personService.GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", personService.CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", personService.DeletePersonEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetConfigurations(w http.ResponseWriter, r *http.Request) {
	config := Infra.GetConfigurations()
	json.NewEncoder(w).Encode(config)
}
