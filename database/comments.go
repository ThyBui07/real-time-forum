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
		ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
		Username STRING NOT NULL,
		PostID INTEGER NOT NULL,
		Content BLOB NOT NULL,
		FOREIGN KEY(Username) REFERENCES Users(Username) ON DELETE CASCADE,
		FOREIGN KEY(PostID) REFERENCES Posts(ID) ON DELETE CASCADE);`

	query, err := db.Prepare(commentsTable)
	u.CheckErr(err)
	query.Exec()
	fmt.Println("Comments created successfully!")
}
