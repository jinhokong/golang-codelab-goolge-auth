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

// RenderIndex Just Rendering HTML template
func RenderIndex(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

// RenderLogin Just Rendering HTML template
func RenderLogin(w http.ResponseWriter, r *http.Request) {
	state := "randomTest"
	url := oauth.GoogleAuthorizationURL(state)
	tmpl.ExecuteTemplate(w, "login.html", url)
}

// RenderProfile Just Rendering HTML template
func RenderProfile(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "profile.html", nil)
}
