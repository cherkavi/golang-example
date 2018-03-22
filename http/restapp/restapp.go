package restapp

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"../../os/exec"
	"github.com/gorilla/mux"
)

const (
	PersonID = "id"
)

// The person Type (more like an object)
type Person struct {
	ID      string   `json:"id,omitempty"`
	Name    string   `json:"name,omitempty"`
	Address *Address `json:"address,omitempty"`
}

func findPersonById(id string) (Person, error) {
	for _, each := range peopleDatabase {
		if each.ID == id {
			return each, nil
		}
	}
	return Person{}, errors.New(fmt.Sprintf("person with ID [%v] not found ", id))
}

func removePersonById(id string) {
	for index, item := range peopleDatabase {
		if item.ID == id {
			peopleDatabase = append(peopleDatabase[:index], peopleDatabase[index+1:]...)
			goto End
		}
	}
End:
}

type Address struct {
	City       string `json:"city,omitempty"`
	PostalCode string `json:"postal-code,omitempty"`
}

// data storage for application
var peopleDatabase []Person

func GetAll(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode(peopleDatabase)
}

func Get(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	person, error := findPersonById(params[PersonID])
	if error != nil {
		response.WriteHeader(404) // not found
	} else {
		json.NewEncoder(response).Encode(person)
	}
}

// create a new item
func Create(response http.ResponseWriter, request *http.Request) {
	requestParams := mux.Vars(request)
	var person Person
	err := json.NewDecoder(request.Body).Decode(&person)
	if err != nil {
		response.WriteHeader(500) // internal server error
		return
	}
	person.ID = requestParams[PersonID]
	peopleDatabase = append(peopleDatabase, person)
	json.NewEncoder(response).Encode(person)
}

// Delete an item
func Delete(response http.ResponseWriter, request *http.Request) {
	inputParameters := mux.Vars(request)
	removePersonById(inputParameters["id"])
	response.WriteHeader(201)
}

func initData() {
	peopleDatabase = append(peopleDatabase,
		Person{ID: "1", Name: "Luck Skywalker", Address: &Address{City: "City X", PostalCode: "80222"}},
		Person{ID: "2", Name: "Hans Solo", Address: &Address{City: "City Z", PostalCode: "80333"}})
}

// main function to boot up everything
func Example() {
	fmt.Println("---- Start REST application ---- ")
	router := mux.NewRouter()
	initData()
	router.HandleFunc("/person", GetAll).Methods("GET")
	router.HandleFunc("/person/{id}", Get).Methods("GET")
	router.HandleFunc("/person/{id}", Create).Methods("POST")
	router.HandleFunc("/person/{id}", Delete).Methods("DELETE")
	go http.ListenAndServe(":8002", router)
	time.Sleep(time.Second * 2)
	url := "http://127.0.0.1:8002/person"
	fmt.Println("---- open url:  ", url)
	exec.Openbrowser(url)
}
