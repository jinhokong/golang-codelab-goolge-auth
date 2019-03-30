package main

import (
	"fmt"
	"log"
	"net/http"
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
	w.Write([]byte("Hello World"))
	return
}
