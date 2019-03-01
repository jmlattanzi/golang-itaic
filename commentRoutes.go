package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// GetCommentsHandler ...Return all comments
func GetCommentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	comments := GetCommentsSlice()
	err := json.NewEncoder(w).Encode(comments)
	if err != nil {
		log.Fatal("Error encoding comments data: ", err)
	}
}

// GetCommentHandler ...Return single comment
func GetCommentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	comment := GetSingleComment(params["id"])
	err := json.NewEncoder(w).Encode(comment)
	if err != nil {
		log.Fatal("Error encoding comment data: ", err)
	}
}

// GetPostCommentsHandler ...Returns all comments for a specific post
func GetPostCommentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	comments := GetCommentsForPost(params["id"])
	err := json.NewEncoder(w).Encode(comments)
	if err != nil {
		log.Fatal("Error encoding post comments: ", err)
	}
}
