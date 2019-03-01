package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// User ...Layout for each user row
type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Hash      string `json:"hash"`
	Bio       string `json:"bio"`
	AvatarURL string `json:"avatar_url"`
	Followers int    `json:"followers"`
}

// Post ...Post structure in DB
type Post struct {
	ID        string `json:"id"`
	UserID    int    `json:"user_id"`
	ImageURL  string `json:"image_url"`
	Caption   string `json:"caption"`
	Time      string `json:"time"`
	LikeCount int    `json:"like_count"`
	Edited    bool   `json:"edited"`
}

// Comment ... Comment structure in DB
type Comment struct {
	ID      string `json:"id"`
	PostID  string `json:"post_id"`
	UserID  int    `json:"user_id"`
	Comment string `json:"comment"`
	Likes   int    `json:"likes"`
}

// Config ...Defines the layout of our config file
type Config struct {
	DatabaseURI string `json:"databaseURI"`
}

var db *sql.DB
var users []User
var posts []Post
var comments []Comment

/* * * * * * * * * * * * * * *
 *		  INIT METHOD		 *
 * * * * * * * * * * * * * * */

//InitDB ...Open a connection to the database and make sure it's connected
func InitDB() {
	// Load the config
	config := LoadConfigurationFile("config.json")

	db, err := sql.Open("postgres", config.DatabaseURI)
	if err != nil {
		log.Fatal("Error while running sql.Open(): ", err)
	}
	defer db.Close() // Hold off on closing the connection until after the function has run

	// So we ping it to make sure we're in
	if err = db.Ping(); err != nil {
		log.Fatal("Error pinging DB", err)
	}

	// Not sure if the concurrency here is needed but oh well
	go GetUsers(db) // Populate the users slice
	GetPosts(db)    // Populate the posts slice
	GetComments(db) // Populate the comments slice
}

/* * * * * * * * * * * * * * *
 *		  SETUP METHODS		 *
 * * * * * * * * * * * * * * */

// LoadConfigurationFile ...Loads config data stored in config.json
func LoadConfigurationFile(filename string) Config {
	var config Config
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		log.Fatal("Error loading configuration: ", err)
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

// GetUsers ...Get all users from the database
func GetUsers(db *sql.DB) []User {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal("Error selecting all users:  ", err)
	}

	fmt.Println("Getting Users....")
	for rows.Next() {
		usr := User{}                                                                                             // setup a temp user
		err := rows.Scan(&usr.ID, &usr.Username, &usr.Email, &usr.Hash, &usr.Bio, &usr.AvatarURL, &usr.Followers) // scan the current row for this info
		if err != nil {
			log.Fatal("Error scanning rows: ", err)
		}

		users = append(users, usr) // append the row info to the slice we made earlier
	}

	if err = rows.Err(); err != nil {
		log.Fatal("Error while scanning rows: ", err)
	}

	return users
}

// GetPosts ...Select all posts from the DB
func GetPosts(db *sql.DB) []Post {
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		log.Fatal("Error selecting all posts: ", err)
	}

	fmt.Println("Getting all posts....")
	for rows.Next() {
		post := Post{}
		err := rows.Scan(&post.ID, &post.UserID, &post.ImageURL, &post.Caption, &post.Time, &post.LikeCount, &post.Edited)
		if err != nil {
			log.Fatal("Error scanning rows in posts table: ", err)
		}

		posts = append(posts, post)
	}

	return posts
}

// GetComments ...Select all comments from the DB
func GetComments(db *sql.DB) []Comment {
	rows, err := db.Query("SELECT * FROM comments")
	if err != nil {
		log.Fatal("Error selecting all comments: ", err)
	}

	fmt.Println("Getting all comments....")
	for rows.Next() {
		comment := Comment{}
		err := rows.Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Comment, &comment.Likes)
		if err != nil {
			log.Fatal("Error scanning rows in users table: ", err)
		}

		comments = append(comments, comment)
	}

	return comments
}

/* * * * * * * * * * * * * * *
 *		  USER METHODS		 *
 * * * * * * * * * * * * * * */

// GetUsersSlice ...Returns the User slice
func GetUsersSlice() []User {
	return users
}

// GetSingleUser ...Returns one user from the slice based on id
func GetSingleUser(id string) User {
	var foundUser User
	for _, user := range users {
		if user.ID == id {
			foundUser = user
		}
	}

	return foundUser
}

// CreateNewUser ...Add a new user to the database
func CreateNewUser() {

}

/* * * * * * * * * * * * * * *
 *		  POST METHODS		 *
 * * * * * * * * * * * * * * */

// GetPostsSlice ...Returns the Post slice
func GetPostsSlice() []Post {
	return posts
}

// GetSinglePost ...Returns single post from the slice based on ID
func GetSinglePost(id string) Post {
	var foundPost Post
	for _, post := range posts {
		if post.ID == id {
			foundPost = post
		}
	}

	return foundPost
}

/* * * * * * * * * * * * * * *
 *		COMMENT METHODS		 *
 * * * * * * * * * * * * * * */

//GetCommentsSlice ...Returns the entire comments slice
func GetCommentsSlice() []Comment {
	return comments
}

// GetSingleComment ...Returns a single comment
func GetSingleComment(id string) Comment {
	var foundComment Comment

	for _, comment := range comments {
		if comment.ID == id {
			foundComment = comment
		}
	}

	return foundComment
}

// GetCommentsForPost ...Returns all comments for a specific post
func GetCommentsForPost(postID string) []Comment {
	/*
		@TODO:
			+ Convert this to a database call
	*/
	var foundComments []Comment
	for _, comment := range comments {
		if comment.PostID == postID {
			foundComments = append(foundComments, comment)
		}
	}

	return foundComments
}
