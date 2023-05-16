package database

import (
	"database/sql"
	"fmt"
	u "server/utils"
)

type Post struct {
	ID           int
	AuthorID     int
	Title        string
	Content      string
	Date         string
	Categories   string
	CommentCount int
}

//create posts table
func CreatePostsTable(db *sql.DB) {
	postsTable := `CREATE TABLE IF NOT EXISTS Posts (
		ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
		AuthorID INTEGER NOT NULL,
		Title TEXT NOT NULL,
		Content TEXT NOT NULL,
		Date TEXT NOT NULL,
		Categories TEXT NOT NULL,
		CommentCount INTEGER NOT NULL DEFAULT 0,
		FOREIGN KEY(AuthorID) REFERENCES Users(ID) ON DELETE CASCADE);`
	query, err := db.Prepare(postsTable)
	u.CheckErr(err)
	query.Exec()
	fmt.Println("Posts table created successfully!")
}

//add posts to posts table
func AddPost(db *sql.DB, authorID int, title string, content string, date string, categories string) error {
	insertQuery := `INSERT INTO Posts (AuthorID, Title, Content, Date, Categories) VALUES ( ?, ?, ?, ?, ?)`
	_, err := db.Exec(insertQuery, authorID, title, content, date, categories)
	if err != nil {
		return err
	}
	return nil
}

func GetPosts() ([]Post, error) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM Posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.AuthorID, &post.Title, &post.Content, &post.Date, &post.Categories, &post.CommentCount)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}

// And here's an example of how you can update the comment count for a post after adding a new comment:

// go
// Copy code
// _, err = db.Exec("UPDATE Posts SET CommentCount = (SELECT COUNT(*) FROM Comments WHERE PostID = ?) WHERE ID = ?", postID, postID)
// if err != nil {
//     return err
// }
