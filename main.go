package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Printf("hello %s , %s", vars["title"], vars["page"])
	})
	http.ListenAndServe(":80", r)
}
