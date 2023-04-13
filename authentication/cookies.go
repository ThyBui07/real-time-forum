package authentication

import (
	"fmt"
	"net/http"
)

func CookiesCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CookiesCheck")
	cookie, err := r.Cookie("session")
	fmt.Println("cookie: ", cookie)
	if err != nil {
		fmt.Println("cookie not found")
		cookie = &http.Cookie{
			Name:     "session",
			Value:    "some value",
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)

	}

}
