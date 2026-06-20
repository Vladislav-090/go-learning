package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, " hello world")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, " users:")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/users", usersHandler)
	http.ListenAndServe(":8080", nil)
}
