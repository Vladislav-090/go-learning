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

var users []User

func userHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	if r.Method != http.MethodPost {
		fmt.Fprintln(w, "method not allowed")
		return
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprintln(w, "invalid json")
		return
	}
	if user.Name == "" {
		fmt.Fprintln(w, "Empty name")
		return
	} else if user.Age <= 0 {
		fmt.Fprintln(w, "age must be positive!")
		return
	}

	users = append(users, user)
	fmt.Fprintln(w, "User added successfully!")
}

func main() {
	http.HandleFunc("/user", userHandler)
	http.ListenAndServe(":8080", nil)
}
