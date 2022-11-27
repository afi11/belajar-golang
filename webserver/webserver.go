package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "apa kabar sayang!")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "halo!")
	})

	http.HandleFunc("/index", index)

	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}