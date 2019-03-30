package handler

import (
	"html/template"
	"net/http"
)

// RenderIndex Just Rendering HTML template
func RenderIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, nil)
}

// RenderLogin Just Rendering HTML template
func RenderLogin(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, nil)
}

// RenderProfile Just Rendering HTML template
func RenderProfile(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/profile.html")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, nil)
}
