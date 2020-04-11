package db

import (
	"database/sql"
	"fmt"
	"log"
)

// Name Struct
type Name struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// User Struct
type User struct {
	ID     int     `json:"userid"`
	Name   Name    `json:"name"`
	Age    int     `json:"age"`
	Weight float32 `json:"weight"`
}

// Users struct
type Users struct {
	Users []User `json:"users"`
}

// GetUsers Retrieves all Users from the database
func GetUsers(db *sql.DB) (users Users, err error) {
	rows, err := runQuery(db, "SELECT * FROM Users")
	if err != nil {
		log.Fatal("ERROR: SELECT Failed to Retrieve Users")
	}

	collection := Users{}
	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.ID, &user.Name.FirstName, &user.Name.LastName, &user.Age, &user.Weight); err != nil {
			log.Fatal("ERROR: Failed to Parse User Data")
		}
		collection.Users = append(collection.Users, user)
	}

	return collection, err
}

// AddUser Inserts user data to Users table
func AddUser(user User, db *sql.DB) {
	query := fmt.Sprintf(`INSERT INTO Users (user_id, first_name, last_name, age, weight) VALUES (%d, "%s", "%s", %d, %f)`,
		user.ID, user.Name.FirstName, user.Name.LastName, user.Age, user.Weight)
	db.Exec(query)
}

// GetUserByUserID Retrieves user data for user with given 'id'
func GetUserByUserID(id string, db *sql.DB) (user User, err error) {
	resUser := User{}
	query := fmt.Sprintf(`SELECT * FROM Users WHERE user_id = "%s"`, id)
	err = db.QueryRow(query).Scan(&resUser.ID, &resUser.Name.FirstName, &resUser.Name.LastName, &resUser.Age, &resUser.Weight)
	return resUser, err
}

// DeleteUserByUserID Removes user data from Users table for user with matching 'id'
func DeleteUserByUserID(id string, db *sql.DB) {
	query := fmt.Sprintf("DELETE FROM Users WHERE user_id = %s", id)
	db.Exec(query)
}
