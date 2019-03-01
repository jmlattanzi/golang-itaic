package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// type User struct {
// 	id        int    `json:"id"`
// 	username  string `json:"username"`
// 	email     string `json:"email"`
// 	hash      string `json:"hash"`
// 	bio       string `json:"bio"`
// 	avatarURL string `json:"avatar_url"`
// 	followers int    `json:"followers"`
// }

// GetUsersHandler ... Gets all the users from the DB
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users := GetUsersSlice()

	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Fatal("Error encoding users: ", err)
	}
}

// GetUserHandler ...Returns a single user
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	user := GetSingleUser(params["id"])

	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Fatal("Error encoding user: ", err)
	}

}
