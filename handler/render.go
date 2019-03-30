package handler

import (
	"html/template"
	"net/http"

	"github.com/golangkorea/codelab/oauth"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func RenderIndex(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func RenderLogin(w http.ResponseWriter, r *http.Request) {
	state := "randomTest"
	url := oauth.GoogleAuthorizationURL(state)
	tmpl.ExecuteTemplate(w, "login.html", url)
}

func RenderProfile(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "profile.html", nil)
}
