package database

import (
	"database/sql"
	"fmt"
	u "server/utils"
)

func CreateLikesTable(db *sql.DB) {
	likesTable := `CREATE TABLE IF NOT EXISTS Likes (
		ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
		Username STRING NOT NULL,
		PostID INTEGER NOT NULL,
		FOREIGN KEY(Username) REFERENCES Users(Username) ON DELETE CASCADE,
		FOREIGN KEY(PostID) REFERENCES Posts(ID) ON DELETE CASCADE);`

}