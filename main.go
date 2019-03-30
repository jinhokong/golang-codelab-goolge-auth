package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const bind = ":1333"

func main() {

	http.HandleFunc("/", RenderIndex)
	http.HandleFunc("/auth", RenderLogin)
	http.HandleFunc("/profile", RenderProfile)
	fmt.Printf("Server is listening on %s\n", bind)

	http.ListenAndServe(bind, nil)

	log.Fatal(http.ListenAndServe(bind, nil))
}

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
