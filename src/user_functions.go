/*
 * TODO:
 * [] Add:
 *		[] User account
 *		[] Update bio
 *		[] Delete user
 */

package main

import (
	"database/sql"
	"fmt"
	"log"
)

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
		log.Fatal("[!] Error while running sql.Open(): ", err)
	}
	defer db.Close()

	// Begin a database transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatal("[!] Error in db.Being(): ", err)
	}

	// Prepare the statement
	stmt, err := tx.Prepare("INSERT INTO go_users (username, email, hash, bio, avatar_url, followers) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		log.Fatal("[!] Error preparing statement: ", err)
	}
	defer stmt.Close()

	// Execute the statement
	_, err = stmt.Exec(&newUser.Username, &newUser.Email, &newUser.Hash, &newUser.Bio, &newUser.AvatarURL, &newUser.Followers)
	if err != nil {
		log.Fatal("[!] Error executing statement: ", err)
	}

	// Commit the transaction
	tx.Commit()

	// Tell the User
	fmt.Println("[!] New user added to database....")

	// Update the Users slice
	GetUsers(db)
}

// GetUserAccount ...Returns the users account
func GetUserAccount(id string) {

}
