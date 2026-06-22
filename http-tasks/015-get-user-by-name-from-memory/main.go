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

func getUserHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		fmt.Fprintln(w, "method not allowed!")
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprintln(w, "name is empty!")
		return
	}

	for _, user := range users {
		if user.Name == name {
			jsonData, err := json.MarshalIndent(user, "", "  ")
			if err != nil {
				fmt.Fprintln(w, "Marshal error")
				return
			}
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, string(jsonData))
			return
		}
	}
	fmt.Fprintln(w, "user not found")
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	if r.Method != http.MethodPost {
		fmt.Fprintln(w, "method not allowed")
		return
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprintln(w, "invalid json!")
		return
	}

	if user.Name == "" {
		fmt.Fprintln(w, "Empty name")
		return
	} else if user.Age <= 0 {
		fmt.Fprintln(w, "Age must be positive!")
		return
	}

	users = append(users, user)
	fmt.Fprintln(w, "User added succsessfully!")
}

func main() {
	http.HandleFunc("/addUser", addUserHandler)
	http.HandleFunc("/getUser", getUserHandler)

	http.ListenAndServe(":8080", nil)
}
