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
	CreateUsersTable(database)
	//addUsers(database, "Gin", "Phan", "GinP", "gin.phan@gritlab.ax", "abc123", 29, "Female")
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
