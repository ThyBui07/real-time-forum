package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	a "server/authentication"
	d "server/database"
	u "server/utils"
)

// func aboutHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println(r.URL.Path)
// 	fmt.Fprint(w, "About Page")
// }

var Database *sql.DB

func init() {
	fmt.Println("Server init")
	Database = d.GetDB()
	fmt.Println("Database:")
}

func main() {

	err := Start()
	fmt.Println(err)
}

func Start() error {
	fs := http.FileServer(http.Dir("./assets"))
	//read the file from the assets folder
	http.Handle("/", fs)
	//handle ws connections
	manager := NewManager()
	http.HandleFunc("/ws", manager.serveWS)
	//handle login- logout
	http.HandleFunc("/login", a.LogIn)
	http.HandleFunc("/logout", a.LogOut)

	//fire up the server
	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	u.CheckErr(err)
	return nil
}
