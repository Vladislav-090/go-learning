package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	if r.Method != http.MethodPost {
		fmt.Fprintln(w, "method not allowed")
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprintln(w, "Invalid json")
		return
	} else if user.Name == "" {
		fmt.Fprintln(w, "Name is empty")
		return
	} else if user.Age <= 0 {
		fmt.Fprintln(w, "Age must be positive!")
		return
	}

	fmt.Fprintln(w, "User created! Name:", user.Name, "Age:", user.Age)
}

func main() {
	http.HandleFunc("/user", userHandler)
	http.ListenAndServe(":8080", nil)
}
