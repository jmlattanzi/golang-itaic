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

	// Open the file and defer closing it until the function is done
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		log.Fatal("Error loading configuration: ", err)
	}

	// decode the json and store it in config
	json.NewDecoder(configFile).Decode(&config)
	return config
}

// GetUsers ...Get all users from the database
func GetUsers(db *sql.DB) []User {
	rows, err := db.Query("SELECT * FROM go_users")
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
	rows, err := db.Query("SELECT * FROM go_posts")
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
	rows, err := db.Query("SELECT * FROM go_comments")
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
func CreateNewUser(newUser User) {
	config := LoadConfigurationFile("config.json")

	db, err := sql.Open("postgres", config.DatabaseURI)
	if err != nil {
		log.Fatal("Error while running sql.Open(): ", err)
	}
	defer db.Close()

	// Begin a database transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatal("Error in db.Being(): ", err)
	}

	// Prepare the statement
	stmt, err := tx.Prepare("INSERT INTO go_users (username, email, hash, bio, avatar_url, followers) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		log.Fatal("Error preparing statement: ", err)
	}
	defer stmt.Close()

	// Execute the statement
	_, err = stmt.Exec(&newUser.Username, &newUser.Email, &newUser.Hash, &newUser.Bio, &newUser.AvatarURL, &newUser.Followers)
	if err != nil {
		log.Fatal("Error executing statement: ", err)
	}

	// Commit the transaction
	tx.Commit()

	// Tell the User
	fmt.Println("New user added to database....")

	// Update the Users slice
	GetUsers(db)
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

// CreateNewPost ...Add a new post to the database
func CreateNewPost(newPost Post) {
	config := LoadConfigurationFile("config.json")

	db, err := sql.Open("postgres", config.DatabaseURI)
	if err != nil {
		log.Fatal("Error while running sql.Open(): ", err)
	}
	defer db.Close()

	// Begin a database transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatal("Error in db.Being(): ", err)
	}

	// Prepare the statement
	stmt, err := tx.Prepare("INSERT INTO go_posts (user_id, image_url, caption, time, like_count, edited) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		log.Fatal("Error preparing statement: ", err)
	}
	defer stmt.Close()

	// Execute the statement
	_, err = stmt.Exec(&newPost.UserID, &newPost.ImageURL, &newPost.Caption, &newPost.Time, &newPost.LikeCount, &newPost.Edited)
	if err != nil {
		log.Fatal("Error executing statement: ", err)
	}

	// Commit the transaction
	tx.Commit()

	// Tell the User
	fmt.Println("New user added to database....")

	// Update the Posts slice
	GetPosts(db)
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
