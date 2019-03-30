package handler

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/golangkorea/codelab/model"
	"github.com/golangkorea/codelab/oauth"
	"github.com/gorilla/sessions"
)

func GoogleOAuthCallBack(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "auth")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	state := session.Values["state"]
	delete(session.Values, "state")

	if state != r.FormValue("state") {
		http.Error(w, "invalid session state", http.StatusUnauthorized)
		return
	}

	token, err := oauth.GoogleOAuthConf.Exchange(context.Background(), r.FormValue("code"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client := oauth.GoogleOAuthConf.Client(context.Background(), token)
	userInfoResp, err := client.Get(oauth.GoogleUserInfoAPI)

	defer userInfoResp.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if userInfoResp.StatusCode != http.StatusOK {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userInfo, err := ioutil.ReadAll(userInfoResp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var user model.User
	json.Unmarshal(userInfo, &user)

	session.Options = &sessions.Options{
		Path:   "/",
		MaxAge: 86400,
	}
	session.Values["user"] = user

	session.Save(r, w)

	http.Redirect(w, r, "/profile", http.StatusFound)
	return
}
