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

func usersHandler(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{Name: "Vladislav", Age: 29},
		{Name: "Viola", Age: 30},
	}
	jsonData, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		fmt.Fprintln(w, "Marshal error")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(jsonData))
}

func main() {
	http.HandleFunc("/users", usersHandler)
	http.ListenAndServe(":8080", nil)
}
