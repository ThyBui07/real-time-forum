package database

import (
	"database/sql"
	"fmt"
	u "server/utils"
)

type Comment struct {
}

func CreateCommentsTable(db *sql.DB) {
	commentsTable := `CREATE TABLE IF NOT EXISTS Comments (
		CommentID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
		UserID INTEGER NOT NULL,
		Username STRING NOT NULL,
		PostID INTEGER NOT NULL,
		Content TEXT NOT NULL,
		CommentLikesCount INTEGER NOT NULL DEFAULT 0,
		FOREIGN KEY(UserID) REFERENCES Users(UserID) ON DELETE CASCADE,
		FOREIGN KEY(Username) REFERENCES Users(Username) ON DELETE CASCADE,
		FOREIGN KEY(PostID) REFERENCES Posts(PostID) ON DELETE CASCADE);`

	query, err := db.Prepare(commentsTable)
	u.CheckErr(err)
	query.Exec()
	fmt.Println("Comments table created successfully!")
}

func AddComment(db *sql.DB, userID int, username string, postID int, content string) error {
	// Prepare the SQL statement to insert the comment
	statement, err := db.Prepare("INSERT INTO Comments (UserID, Username, PostID, Content) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	// Execute the SQL statement with the provided parameters
	_, err = statement.Exec(userID, username, postID, content)
	if err != nil {
		return err
	}

	// Update the CommentCount in the Posts table
	updateQuery := "UPDATE Posts SET CommentCount = CommentCount + 1 WHERE PostID = ?"
	_, err = db.Exec(updateQuery, postID)
	if err != nil {
		return err
	}

	return nil
}
