package database

import (
	"database/sql"
	"fmt"
	u "server/utils"

	_ "github.com/mattn/go-sqlite3"
)

func GetDB() *sql.DB {
	//open the database
	fmt.Println("Opening database...")
	database, err := sql.Open("sqlite3", "./forum.db")
	u.CheckErr(err)
	createTable(database)
	// addUsers(database, "Ankita", "Maudie", "Game Developer", 140000)
	// addUsers(database, "Emiliana", "Alfiya", "Bakend Developer", 120000)
	// addUsers(database, "Emmet", "Brian", "DevOps Developer", 110000)
	// addUsers(database, "Reidun", "Jorge", "Dtabase Developer", 140000)
	// addUsers(database, "Tyrone", "Silvia", "Front-End Developer", 109000)

	// fetch records:
	fmt.Println("Fetching records...")
	fmt.Println(database)
	//FetchUsers(database)
	return database
}

//create users table
func createTable(db *sql.DB) {
	users_table := `CREATE TABLE IF NOT EXISTS users (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        FirstName TEXT NOT NULL,
        LastName TEXT NOT NULL,
		UserName TEXT NOT NULL UNIQUE,
		Email TEXT NOT NULL UNIQUE,
		Password TEXT NOT NULL
        );`
	query, err := db.Prepare(users_table)
	u.CheckErr(err)
	query.Exec()
	fmt.Println("Table created successfully!")
}

//add users to users table
func addUsers(db *sql.DB, FirstName string, LastName string, Dept string, Salary int) {
	records := `INSERT INTO users(FirstName, LastName, Dept, Salary) VALUES (?, ?, ?, ?)`
	query, err := db.Prepare(records)
	u.CheckErr(err)
	_, err = query.Exec(FirstName, LastName, Dept, Salary)
	u.CheckErr(err)
}

//fetch users

func FetchUsers(db *sql.DB) {
	record, err := db.Query("SELECT * FROM users")
	u.CheckErr(err)
	defer record.Close()
	for record.Next() {
		var id int
		var FirstName string
		var LastName string
		var Dept string
		var Salary int
		record.Scan(&id, &FirstName, &LastName, &Dept, &Salary)
		fmt.Printf("User: %d %s %s %s %d", id, FirstName, LastName, Dept, Salary)
	}
}
