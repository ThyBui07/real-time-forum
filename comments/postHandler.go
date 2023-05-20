package comments

import (
	"net/http"
	a "server/authentication"
)

func CreateCommentToPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if a.SessionCheck(w, r) {

	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}

}
