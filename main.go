package main

import (
	"fmt"
	"log"
	"net/http"
)

const bind = ":1333"

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	fmt.Printf("Server is listening on %s\n", bind)

	http.ListenAndServe(bind, nil)

	log.Fatal(http.ListenAndServe(bind, nil))

}
