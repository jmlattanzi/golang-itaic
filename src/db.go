package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB
var users []User
var posts []Post
var userPosts []Post
var comments []Comment
var config Config

/* * * * * * * * * * * * * * *
 *		  INIT METHOD		 *
 * * * * * * * * * * * * * * */

//InitDB ...Open a connection to the database and make sure it's connected
func InitDB() {
	// Load the config
	config := LoadConfigurationFile("src/config.json")

	db, err := sql.Open("postgres", config.DatabaseURI)
	if err != nil {
		log.Fatal("[!] Error while running sql.Open(): ", err)
	}
	defer db.Close() // Hold off on closing the connection until after the function has run

	// So we ping it to make sure we're in
	if err = db.Ping(); err != nil {
		log.Fatal("[!] Error pinging DB", err)
	}

	// Not sure if concurrency here is needed but oh well
	GetUsers(db)    // Populate the users slice
	GetPosts(db)    // Populate the posts slice
	GetComments(db) // Populate the comments slice
}

/* * * * * * * * * * * * * * *
 *		  SETUP METHODS		 *
 * * * * * * * * * * * * * * */

// LoadConfigurationFile ...Loads config data stored in config.json
func LoadConfigurationFile(filename string) Config {
	fmt.Println("[-] Loading configuration....")
	// Open the file and defer closing it until the function is done
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		log.Fatal("[!] Error loading configuration: ", err)
	}

	// decode the json and store it in config
	json.NewDecoder(configFile).Decode(&config)
	return config
}

// GetUsers ...Get all users from the database
func GetUsers(db *sql.DB) []User {
	rows, err := db.Query("SELECT * FROM go_users")
	if err != nil {
		log.Fatal("[!] Error selecting all users:  ", err)
	}

	fmt.Println("[+] Getting all users....")
	for rows.Next() {
		usr := User{}                                                                                             // setup a temp user
		err := rows.Scan(&usr.ID, &usr.Username, &usr.Email, &usr.Hash, &usr.Bio, &usr.AvatarURL, &usr.Followers) // scan the current row for this info
		if err != nil {
			log.Fatal("[!] Error scanning rows: ", err)
		}

		users = append(users, usr) // append the row info to the slice we made earlier
	}

	if err = rows.Err(); err != nil {
		log.Fatal("[!] Error while scanning rows: ", err)
	}

	return users
}

// GetPosts ...Select all posts from the DB
func GetPosts(db *sql.DB) []Post {
	rows, err := db.Query("SELECT * FROM go_posts")
	if err != nil {
		log.Fatal("[!] Error selecting all posts: ", err)
	}

	fmt.Println("[+] Getting all posts....")
	for rows.Next() {
		post := Post{}
		err := rows.Scan(&post.ID, &post.UserID, &post.ImageURL, &post.Caption, &post.Time, &post.LikeCount, &post.Edited)
		if err != nil {
			log.Fatal("[!] Error scanning rows in posts table: ", err)
		}

		posts = append(posts, post)
	}

	return posts
}

// GetComments ...Select all comments from the DB
func GetComments(db *sql.DB) []Comment {
	rows, err := db.Query("SELECT * FROM go_comments")
	if err != nil {
		log.Fatal("[!] Error selecting all comments: ", err)
	}

	fmt.Println("[+] Getting all comments....")
	for rows.Next() {
		comment := Comment{}
		err := rows.Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Comment, &comment.Likes)
		if err != nil {
			log.Fatal("[!] Error scanning rows in users table: ", err)
		}

		comments = append(comments, comment)
	}

	return comments
}
