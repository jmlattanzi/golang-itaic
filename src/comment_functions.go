/*
 * TODO:
 * [] Check sort order
 * [] Add delete method
 * [] Add update method
 *
 */

package main

import (
	"database/sql"
	"fmt"
	"log"
)

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

// CreateNewComment ...Add a new post to the database
func CreateNewComment(newComment Comment) {
	config := LoadConfigurationFile("config.json")

	db, err := sql.Open("postgres", config.DatabaseURI)
	if err != nil {
		log.Fatal("[!] Error while running sql.Open(): ", err)
	}
	defer db.Close()

	// Begin a database transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatal("[!] Error in db.Being(): ", err)
	}

	// Prepare the statement
	stmt, err := tx.Prepare("INSERT INTO go_comments (post_id, user_id, comment, likes) VALUES ($1, $2, $3, $4)")
	if err != nil {
		log.Fatal("[!] Error preparing statement: ", err)
	}
	defer stmt.Close()

	// Execute the statement
	_, err = stmt.Exec(&newComment.PostID, &newComment.UserID, &newComment.Comment, &newComment.Likes)
	if err != nil {
		log.Fatal("[!] Error executing statement: ", err)
	}

	// Commit the transaction
	tx.Commit()

	// Tell the User
	fmt.Println("[!] New user added to database....")

	// Update the Comments slice
	GetComments(db)
}
