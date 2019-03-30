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
