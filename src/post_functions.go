/*
 * TODO:
 * [ ] Check sort order
 * [ ] Add:
 *		[ ] Update
 *		[ ] Upload
 *		[ ] Delete
 *		[ ] Like
 *		[x] User posts
 */

package main

import (
	"database/sql"
	"fmt"
	"log"
	"sort"
	"strconv"
)

/* * * * * * * * * * * * * * *
 *		  POST METHODS		 *
 * * * * * * * * * * * * * * */

// GetPostsSlice ...Returns the Post slice
func GetPostsSlice() []Post {
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].ID < posts[j].ID
	})

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
		log.Fatal("[!] Error while running sql.Open(): ", err)
	}
	defer db.Close()

	// Begin a database transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatal("[!] Error in db.Being(): ", err)
	}

	// Prepare the statement
	stmt, err := tx.Prepare("INSERT INTO go_posts (user_id, image_url, caption, time, like_count, edited) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		log.Fatal("[!] Error preparing statement: ", err)
	}
	defer stmt.Close()

	// Execute the statement
	_, err = stmt.Exec(&newPost.UserID, &newPost.ImageURL, &newPost.Caption, &newPost.Time, &newPost.LikeCount, &newPost.Edited)
	if err != nil {
		log.Fatal("[!] Error executing statement: ", err)
	}

	// Commit the transaction
	tx.Commit()

	// Tell the User
	fmt.Println("[!] New user added to database....")

	// Update the Posts slice
	GetPosts(db)
}

// GetUserPosts ...Get all of a specific users posts
func GetUserPosts(UserID string) []Post {
	db, err := sql.Open("postgres", config.DatabaseURI)
	if err != nil {
		log.Fatal("[!] Error while running sql.Open(): ", err)
	}
	defer db.Close()

	intUserID, err := strconv.Atoi(UserID)
	if err != nil {
		log.Fatal("[!] Error converting ID to int")
	}
	rows, err := db.Query("SELECT * FROM go_posts WHERE user_id = $1", intUserID)
	if err != nil {
		log.Fatal("[!] Error during query")
	}

	for rows.Next() {
		up := Post{}
		err := rows.Scan(&up.ID, &up.UserID, &up.ImageURL, &up.Caption, &up.Time, &up.LikeCount, &up.Edited)
		if err != nil {
			log.Fatal("[!] Error scanning rows")
		}

		userPosts = append(userPosts, up)
	}

	return userPosts
}
