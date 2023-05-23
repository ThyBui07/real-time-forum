package database

import (
	"database/sql"
	"errors"
	"fmt"
	u "server/utils"
)

func CreateCommentLikesTable(db *sql.DB) {
	commentLikesTable := `CREATE TABLE IF NOT EXISTS CommentLikes (
		CommentLikesID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
		UserID INTEGER NOT NULL,
		CommentID INTEGER NOT NULL,
		FOREIGN KEY(UserID) REFERENCES Users(UserID) ON DELETE CASCADE,
		FOREIGN KEY(CommentID) REFERENCES Comments(CommentID) ON DELETE CASCADE);`
	query, err := db.Prepare(commentLikesTable)
	u.CheckErr(err)
	query.Exec()
	fmt.Println("Comment Likes Table created successfully!")
}

func AddLikeToComment(db *sql.DB, userID, commentID int) error {
	// Insert the like into the Likes table
	insertQuery := "INSERT INTO CommentLikes(UserID, CommentID) VALUES(?, ?)"
	_, err := db.Exec(insertQuery, userID, commentID)
	if err != nil {
		return err
	}
	// Check if the user has already liked the post
	checkQuery := "SELECT COUNT(*) FROM CommentLikes WHERE UserID = ? AND CommentID = ?"
	var count int
	err = db.QueryRow(checkQuery, userID, commentID).Scan(&count)
	if err != nil {
		return err
	}

	// If the user has already liked the post, return an error
	if count > 0 {
		return errors.New("User has already liked the comment")
	}

	// Update the LikesCount in the Comment table
	updateQuery := "UPDATE Comments SET CommentLikesCount = CommentLikesCount + 1 WHERE CommentID = ?"
	_, err = db.Exec(updateQuery, commentID)
	if err != nil {
		return err
	}

	return nil
}

func RemoveLikeFromComment(db *sql.DB, userID, commentID int) error {
	// Remove the like from the Likes table
	deleteQuery := "DELETE FROM CommentLikes WHERE UserID = ? AND CommentID = ?"
	_, err := db.Exec(deleteQuery, userID, commentID)
	if err != nil {
		return err
	}

	// Update the LikesCount in the Posts table
	updateQuery := "UPDATE Comments SET CommentLikesCount = CommentLikesCount - 1 WHERE CommentID = ?"
	_, err = db.Exec(updateQuery, commentID)
	if err != nil {
		return err
	}

	return nil
}
