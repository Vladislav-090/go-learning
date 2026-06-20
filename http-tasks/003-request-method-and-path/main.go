package main

import (
	"fmt"
	"net/http"
)

func infoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Method:", r.Method)
	fmt.Fprintln(w, "Path:", r.URL.Path)
}

func main() {
	http.HandleFunc("/info", infoHandler)
	http.ListenAndServe(":8080", nil)
}
