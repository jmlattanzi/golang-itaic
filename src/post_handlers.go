package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// GetPostsHandler ...Return all posts
func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	posts := GetPostsSlice()

	err := json.NewEncoder(w).Encode(posts)
	if err != nil {
		log.Fatal("[!] Error encoding post data: ", err)
	}
}

// GetSinglePostHandler ...Returns a single post
func GetSinglePostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	post := GetSinglePost(params["id"])
	err := json.NewEncoder(w).Encode(post)
	if err != nil {
		log.Fatal("[!] Error encoding post data: ", err)
	}
}

// CreatePostHandler ...Add a new psot
func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	newPost := Post{}
	err := json.NewDecoder(r.Body).Decode(&newPost)
	if err != nil {
		log.Fatal("[!] Error decoding data in request body (CreatePostHandler): ", err)
	}

	CreateNewPost(newPost)
}

// GetUserPostsHandler ...Get all of a users posts
func GetUserPostsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	userPosts := GetUserPosts(params["id"])
	err := json.NewEncoder(w).Encode(userPosts)
	if err != nil {
		log.Fatal("[!] Error encoding data")
	}
}
