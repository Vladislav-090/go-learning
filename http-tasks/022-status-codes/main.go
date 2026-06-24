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

func addUserHandler(w http.ResponseWriter, r *http.Request) {

	var user User

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid json", http.StatusBadRequest)
		return
	}

	if user.Name == "" {
		http.Error(w, "Empty name", http.StatusBadRequest)
		return
	} else if user.Age <= 0 {
		http.Error(w, "age must be positive!", http.StatusBadRequest)
		return
	}

	for _, userExist := range users {
		if userExist.Name == user.Name {
			http.Error(w, "name already exist", http.StatusConflict)
			return
		}
	}

	users = append(users, user)
	fmt.Fprintln(w, "user succsessfully added!")

}

func main() {
	http.HandleFunc("/addUser", addUserHandler)

	http.ListenAndServe(":8080", nil)
}
