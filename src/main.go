package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, "Home")
}

func main() {
	// Create routes
	InitDB()
	router := mux.NewRouter()

	// User routes
	router.HandleFunc("/", getHome).Methods("GET")
	router.HandleFunc("/users", GetUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", GetUserHandler).Methods("GET")
	router.HandleFunc("/users", CreateUserHandler).Methods("POST")

	// Post routes
	router.HandleFunc("/posts", GetPostsHandler).Methods("GET")
	router.HandleFunc("/posts/{id}", GetSinglePostHandler).Methods("GET")
	router.HandleFunc("/posts/user/{id}", GetUserPostsHandler).Methods("GET")
	router.HandleFunc("/posts", CreatePostHandler).Methods("POST")
	// upload
	// update
	// delete

	// Comment routes
	router.HandleFunc("/comments", GetCommentsHandler).Methods("GET")
	router.HandleFunc("/comments/{id}", GetCommentHandler).Methods("GET")
	router.HandleFunc("/comments/posts/{id}", GetPostCommentsHandler).Methods("GET")
	router.HandleFunc("/comments", CreateNewCommentHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
