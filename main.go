package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	a "server/authentication"
	comments "server/comments"
	d "server/database"
	likes "server/likes"
	posts "server/posts"
	users "server/users"

	u "server/utils"
	ws "server/websocket"
)

// func aboutHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "About Page")
// }

var Database *sql.DB

func init() {
	fmt.Println("Server init")
	Database = d.GetDB()
}

func main() {
	err := Start()
	u.CheckErr(err)
}

func Start() error {
	fs := http.FileServer(http.Dir("./assets"))
	//read the file from the assets folder
	http.Handle("/", fs)
	//handle ws connections
	manager := ws.NewManager()
	http.HandleFunc("/ws", manager.ServeWS)
	//handle login- logout
	http.HandleFunc("/login", a.LogIn)
	http.HandleFunc("/logout", a.LogOut)
	http.HandleFunc("/posts", posts.GetPostsHandler)
	http.HandleFunc("/posts/create", posts.CreatePostHandler)
	http.HandleFunc("/comments/create", comments.CreateCommentToPost)
	http.HandleFunc("/likes/create", likes.AddLiketoPostHandler)
	http.HandleFunc("/users/create", users.CreateUserHandler)

	//http.HandleFunc("/homepage", a.SessionCheck)

	//fire up the server
	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	u.CheckErr(err)
	return nil
}
