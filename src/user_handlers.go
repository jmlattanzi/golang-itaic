package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// GetUsersHandler ... Gets all the users from the DB
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users := GetUsersSlice()

	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Fatal("[!] Error encoding users: ", err)
	}
}

// GetUserHandler ...Returns a single user
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	user := GetSingleUser(params["id"])

	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Fatal("[!] Error encoding user: ", err)
	}

}

// CreateUserHandler ...Adds a new user and returns the updated user list
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Set newUser to be an empty user struct
	newUser := User{}

	// decode the request body and store it into newUser
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		log.Fatal("[!] Error decoding data in request body (CreateUserHandler): ", err)
	}

	CreateNewUser(newUser)
}
