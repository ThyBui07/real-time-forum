package database

import (
	"database/sql"
	"fmt"
	u "server/utils"
)

type Post struct {
}

//create posts table
func CreatePostsTable(db *sql.DB) {
	postsTable := `CREATE TABLE IF NOT EXISTS Posts (
		ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
		AuthorID INTEGER NOT NULL,
		Title TEXT NOT NULL,
		Content TEXT NOT NULL,
		Date TEXT NOT NULL,
		CategoryIDs TEXT NOT NULL,
		Categories TEXT NOT NULL,
		FOREIGN KEY(AuthorID) REFERENCES Users(ID) ON DELETE CASCADE);`
	query, err := db.Prepare(postsTable)
	u.CheckErr(err)
	query.Exec()
	fmt.Println("Posts table created successfully!")
}

//add users to users table
// func AddUsers(db *sql.DB, FirstName string, LastName string, UserName string, Email string, Password string, Age int, Gender string) {
// 	records := `INSERT INTO users(FirstName, LastName, UserName, Email, Password, Age, Gender) VALUES (?, ?, ?, ?, ?, ?, ?)`
// 	query, err := db.Prepare(records)
// 	u.CheckErr(err)
// 	_, err = query.Exec(FirstName, LastName, UserName, Email, Password, Age, Gender)
// 	u.CheckErr(err)
// }

//get user by email
//stmt is a prepared statement in the Go programming language that represents a SQL query that has been prepared and is ready to be executed against a database.
// func GetUserByEmail(email string) (*User, error) {
// 	db, err := sql.Open("sqlite3", "./forum.db")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer db.Close()

// 	stmt, err := db.Prepare("SELECT * FROM users WHERE email = ?")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer stmt.Close()

// 	var user User
// 	err = stmt.QueryRow(email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.UserName, &user.Email, &user.Password, &user.Age, &user.Gender)
// 	fmt.Println("err from GetUserByEmail: ", err)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, nil // user not found
// 		} else {
// 			return nil, err
// 		}
// 	}

// 	return &user, nil
// }
