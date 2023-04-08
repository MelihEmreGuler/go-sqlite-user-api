package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type User struct {
	UserID       int
	UserName     string
	UserPassword string
}

var UserList []User

// Function to connect to the database and retrieve user data
func Database() {
	fmt.Println("database")

	var err error
	db, err = sql.Open("sqlite3", "../../dbrest.db")
	checkError(err)

	getUser() // Retrieve user data from the database
	fmt.Println(UserList)
}

// Function to delete a user from the database
func DeleteUser(userID int) {
	stmt, err := db.Prepare("delete from tblUser where userID=?")
	checkError(err)

	res, err := stmt.Exec(userID)
	checkError(err)

	_, err = res.RowsAffected()
	checkError(err)

	fmt.Println("data deleted.")

	getUser()
}

// Function to update a user's information in the database
func UpdateUser(userID int, userName, userPassword string) {
	stmt, err := db.Prepare("update tblUser set userName=?, userPassword=? where userID=?")
	checkError(err)

	res, err := stmt.Exec(userName, userPassword, userID)
	checkError(err)

	_, err = res.RowsAffected()
	checkError(err)

	getUser()
}

// Function to retrieve user data from the database
func getUser() {
	UserList = nil // Clear UserList before adding new data

	rows, err := db.Query("Select * from tblUser") // Query the database for all users

	if err == nil {
		for rows.Next() {
			var userID int
			var userName string
			var userPassword string

			err := rows.Scan(&userID, &userName, &userPassword)

			if err == nil {
				UserList = append(UserList, User{UserID: userID, UserName: userName, UserPassword: userPassword})
				// Add the retrieved user data to UserList
			} else {
				fmt.Println(err)
			}
		}
	} else {
		fmt.Println(err)
		return
	}

	rows.Close()
}

// Function to add a new user to the database
func AddUser(userName string, userPassword string) {

	stmt, err := db.Prepare("insert into tblUser(userName, userPassword) values(?,?)") // Prepare a SQL statement to insert new user data
	checkError(err)

	res, err := stmt.Exec(userName, userPassword) // Execute the SQL statement with the provided user data
	checkError(err)

	id, err := res.LastInsertId()
	checkError(err)

	fmt.Printf("save successful: %v \n", id)

	getUser() // Update the UserList with the newly added user data
}

// Function to check for errors and panic if found
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
