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
		fmt.Fprintln(w, "method not allowed!")
		return
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprintln(w, "invalid json")
		return
	}

	if user.Name == "" {
		fmt.Fprintln(w, "Name is empty!")
		return
	} else if user.Age <= 0 {
		fmt.Fprintln(w, "age ,ust be positive!")
		return
	}

	users = append(users, user)
	fmt.Fprintln(w, "user added succsessfully")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Fprintln(w, "method not allowed!")
		return
	}
	w.Header().Set("Content-Type", "application/json")

	jsonData, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		fmt.Fprintln(w, "Marshal error")
		return
	}

	fmt.Fprintln(w, string(jsonData))
}

func main() {
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/users", usersHandler)

	http.ListenAndServe(":8080", nil)
}
