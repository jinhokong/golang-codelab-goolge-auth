package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golangkorea/codelab/handler"
)

const bind = ":1333"

func main() {

	http.HandleFunc("/", handler.RenderIndex)
	http.HandleFunc("/auth", handler.RenderLogin)
	http.HandleFunc("/auth/callback", handler.GoogleOAuthCallBack)
	http.HandleFunc("/profile", handler.RenderProfile)
	fmt.Printf("Server is listening on http://localhost%s\n", bind)

	http.ListenAndServe(bind, nil)

	log.Fatal(http.ListenAndServe(bind, nil))
}
