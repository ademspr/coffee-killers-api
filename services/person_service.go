package services

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	Entities "../entities"
)

type PersonService struct {
	people []Entities.Person
}

func NewPersonService(people []Entities.Person) PersonService {
	s := PersonService{people}
	return s
}

func (s *PersonService) GetPersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range s.people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Entities.Person{})
}
func (s *PersonService) GetPeopleEndpoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(s.people)
}
func (s *PersonService) CreatePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Entities.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	s.people = append(s.people, person)
	json.NewEncoder(w).Encode(s.people)
}
func (s *PersonService) DeletePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range s.people {
		if item.ID == params["id"] {
			s.people = append(s.people[:index], s.people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(s.people)
	}
}
