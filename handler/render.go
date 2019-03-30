package handler

import (
	"html/template"
	"net/http"
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
	tmpl.ExecuteTemplate(w, "login.html", nil)
}

// RenderProfile Just Rendering HTML template
func RenderProfile(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "profile.html", nil)
}
