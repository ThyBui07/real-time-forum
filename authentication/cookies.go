package authentication

import (
	"net/http"
	"strings"
)

func SessionCheck(w http.ResponseWriter, r *http.Request) bool {
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
	authenticated := session.Values["authenticated"]
	if authenticated != nil && authenticated != false {
		return true
	} else {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return false
	}
}

// func GetPosts(w http.ResponseWriter, r *http.Request) {
//     if SessionCheck(w, r) {
//         // session is authenticated, proceed with processing the request
//         // do something here
//     } else {
//         // session is not authenticated, SessionCheck already returned an error response
//         return
//     }
// }
