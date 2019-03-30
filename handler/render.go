package handler

import (
	"encoding/gob"
	"html/template"
	"net/http"

	"github.com/golangkorea/codelab/model"
	"github.com/golangkorea/codelab/oauth"
	"github.com/gorilla/sessions"
)

var tmpl *template.Template
var store *sessions.CookieStore

func init() {
	store = sessions.NewCookieStore([]byte("very-secret-key"))
	tmpl = template.Must(template.ParseGlob("templates/*.html"))

	gob.Register(model.User{})
}

func RenderIndex(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func RenderLogin(w http.ResponseWriter, r *http.Request) {
	state := "randomTest"

	session, _ := store.Get(r, "auth")
	session.Options = &sessions.Options{
		Path:   "/auth",
		MaxAge: 300,
	}
	session.Values["state"] = state
	session.Save(r, w)

	url := oauth.GoogleAuthorizationURL(state)
	tmpl.ExecuteTemplate(w, "login.html", url)
}

func RenderProfile(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "auth")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, ok := session.Values["user"]; !ok {
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}

	user := session.Values["user"].(model.User)

	tmpl.ExecuteTemplate(w, "profile.html", user)

}
