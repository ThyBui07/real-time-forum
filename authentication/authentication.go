package authentication

import (
	"encoding/json"
	"fmt"
	"net/http"
	d "server/database"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/gorilla/sessions"
)

// sessions is a Go package used for session management in web applications.
// NewCookieStore is a function provided by the sessions package that returns a new cookie-based session store.
// It takes a single argument, which is a byte slice representing a secret key used to sign and encrypt the session data.
// []byte("secret-key") creates a new byte slice containing the string "secret-key".
// This string is used as the secret key for signing and encrypting the session data.
// The resulting cookie store is assigned to the variable store.
var store = sessions.NewCookieStore([]byte("secret-key"))

type UserData struct {
	User string
	// Posts []d.Post
}

func LogIn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login")
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var data map[string]string
		err := decoder.Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		email := data["email"]
		password := data["password"]

		var storedPassword string
		user, err := d.GetUserByEmail(email)
		// need to improve here

		storedPassword = user.Password

		// Compare the stored password with the password received from the front-end
		if password != storedPassword {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		if user != nil && password == storedPassword {
			// Generate a new UUID
			id, err := uuid.NewV4()
			if err != nil {
				http.Error(w, "Failed to generate session ID", http.StatusInternalServerError)
				return
			}
			// Combine the UUID with the session name using a delimiter
			sessionName := "session-name-" + id.String()

			// Create a new session for the user with the combined session name
			session, _ := store.Get(r, sessionName)

			session.Values["authenticated"] = true
			// Saves all sessions used during the current request
			session.Save(r, w)

			w.WriteHeader(http.StatusOK)
			// Encode the user as JSON
			userJSON, err := json.Marshal(user)
			if err != nil {
				http.Error(w, "Failed to encode user", http.StatusInternalServerError)
				return
			}
			//ko nen, login chi tra ve user
			// posts, err := d.GetPosts()
			// data := UserData{
			// 	User:  string(userJSON),
			// 	Posts: posts,
			// }
			jsonData, err := json.Marshal(userJSON)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// write the JSON-encoded data to the response writer
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
		}
	}

}

func LogOut(w http.ResponseWriter, r *http.Request) {
	fmt.Println("logout")
	cookies := r.Cookies()
	var sessionCookie *http.Cookie
	for _, cookie := range cookies {
		if strings.HasPrefix(cookie.Name, "session-name-") {
			sessionCookie = cookie
			break
		}
	}
	// get the session from the session store
	session, _ := store.Get(r, sessionCookie.Name)
	// Set the authenticated value on the session to false
	session.Values["authenticated"] = false
	session.Save(r, w)
	w.Write([]byte("Logout Successful"))
}

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("register")
	w.Write([]byte("register"))
	/*
		1. check username criteria
		2. check password criteria
		3. check if username is already exists in database
		4. create bcrypt hash from password
		5. insert username and password hash in database
		(email validation will be in another video)
	*/
}
