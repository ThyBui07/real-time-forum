package authentication

import (
	"fmt"
	"net/http"
)

func LogIn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login")

}

func LogOut(w http.ResponseWriter, r *http.Request) {
	fmt.Println("logout")
}

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("register")
	/*
		1. check username criteria
		2. check password criteria
		3. check if username is already exists in database
		4. create bcrypt hash from password
		5. insert username and password hash in database
		(email validation will be in another video)
	*/
}
