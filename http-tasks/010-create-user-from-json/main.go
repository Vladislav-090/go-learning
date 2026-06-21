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
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprintln(w, "Invalid Json")
		return
	}
	fmt.Fprintln(w, "User created!", user.Name)
}

func main() {
	http.HandleFunc("/user", userHandler)
	http.ListenAndServe(":8080", nil)
}
