package handler

import (
	"net/http"
)

// GoogleOAuthCallBack a
func GoogleOAuthCallBack(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/profile", http.StatusFound)
	return
}
