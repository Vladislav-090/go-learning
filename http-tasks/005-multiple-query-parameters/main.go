package main

import (
	"fmt"
	"net/http"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")

	if name == "" {
		fmt.Fprintln(w, "Empty name")
		return
	}
	if age == "" {
		fmt.Fprintln(w, "Empty age")
		return
	}
	fmt.Fprintln(w, "Hello", name, "Age", age)
}

func main() {
	http.HandleFunc("/greet", greetHandler)
	http.ListenAndServe(":8080", nil)
}
