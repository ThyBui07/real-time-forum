package database

import (
	"database/sql"
	"errors"
	"fmt"
	u "server/utils"
)

func CreatePostLikesTable(db *sql.DB) {
	postLikesTable := `CREATE TABLE IF NOT EXISTS PostLikes (
		PostLikesID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
		UserID INTEGER NOT NULL,
		PostID INTEGER NOT NULL,
		FOREIGN KEY(UserID) REFERENCES Users(UserID) ON DELETE CASCADE,
		FOREIGN KEY(PostID) REFERENCES Posts(PostID) ON DELETE CASCADE);`
	query, err := db.Prepare(postLikesTable)
	u.CheckErr(err)
	query.Exec()
	fmt.Println("Post Likes Table created successfully!")
}

func AddLikeToPost(db *sql.DB, userID, postID int) error {
	// Insert the like into the Likes table
	insertQuery := "INSERT INTO PostLikes(UserID, PostID) VALUES(?, ?)"
	_, err := db.Exec(insertQuery, userID, postID)
	if err != nil {
		return err
	}
	// Check if the user has already liked the post
	checkQuery := "SELECT COUNT(*) FROM PostLikes WHERE UserID = ? AND PostID = ?"
	var count int
	err = db.QueryRow(checkQuery, userID, postID).Scan(&count)
	if err != nil {
		return err
	}

	// If the user has already liked the post, return an error
	if count > 0 {
		return errors.New("User has already liked the post")
	}

	// Update the LikesCount in the Posts table
	updateQuery := "UPDATE Posts SET PostLikesCount = PostLikesCount + 1 WHERE PostID = ?"
	_, err = db.Exec(updateQuery, postID)
	if err != nil {
		return err
	}

	return nil
}

func RemoveLikeFromPost(db *sql.DB, userID, postID int) error {
	// Remove the like from the Likes table
	deleteQuery := "DELETE FROM PostLikes WHERE UserID = ? AND PostID = ?"
	_, err := db.Exec(deleteQuery, userID, postID)
	if err != nil {
		return err
	}

	// Update the LikesCount in the Posts table
	updateQuery := "UPDATE Posts SET LikesCount = LikesCount - 1 WHERE PostID = ?"
	_, err = db.Exec(updateQuery, postID)
	if err != nil {
		return err
	}

	return nil
}
