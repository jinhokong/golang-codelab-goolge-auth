package handler

import (
	"net/http"
)

func GoogleOAuthCallBack(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/profile", http.StatusFound)
	return
}
