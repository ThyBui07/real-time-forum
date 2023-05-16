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
	CreatePostsTable(database)
	CreateCommentsTable(database)
	//fake data for testing
	// err = AddUser(database, "Gin", "Phan", "GinP", "gin.phan@gritlab.ax", "abc123", 29, "Female")
	// err = AddUser(database, "Nafisah", "Rantasalmi", "NafisahR", "nafisah.rantasalmi@gritlab.ax", "abc123456", 39, "Female")

	// err = AddPost(database, 1, "My First Post", "Hello, world!", "2023-04-27", "Programming, Golang")
	// err = AddPost(database, 1, "My Favorite Books", "Here are some of my favorite books: 1984, To Kill a Mockingbird, and The Catcher in the Rye.", "2023-04-27", "Books, Literature")

	// err = AddComment(database, "GinP", 1, "This is a great post!")
	//add this every time add comment to update comment count => add postID to AddComment function
	// _, err = database.Exec("UPDATE Posts SET CommentCount = (SELECT COUNT(*) FROM Comments WHERE PostID = ?) WHERE ID = ?", 1, 1)
	// err = AddComment(database, "NafisahR", 1, "I agree, this post is very insightful.")
	// _, err = database.Exec("UPDATE Posts SET CommentCount = (SELECT COUNT(*) FROM Comments WHERE PostID = ?) WHERE ID = ?", 1, 1)
	// err = AddComment(database, "GinP", 2, "Thanks for sharing this information.")
	// _, err = database.Exec("UPDATE Posts SET CommentCount = (SELECT COUNT(*) FROM Comments WHERE PostID = ?) WHERE ID = ?", 2, 2)

	u.CheckErr(err)
	fmt.Println("Fetching records...")
	return database
}
