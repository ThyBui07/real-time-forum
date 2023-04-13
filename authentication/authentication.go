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
