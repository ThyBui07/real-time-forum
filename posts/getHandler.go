package posts

import (
	"encoding/json"
	"net/http"
	a "server/authentication"
	d "server/database"
)

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if a.SessionCheck(w, r) {
		posts, err := d.GetPosts()
		if err != nil {
			// handle the error, e.g. log it or return a specific error response to the client
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		// Write the response as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(posts)
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}
