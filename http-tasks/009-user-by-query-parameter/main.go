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
	users := []User{
		{Name: "Vladislav", Age: 29},
		{Name: "Viola", Age: 20},
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprintln(w, " Name is empty!")
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
	fmt.Fprintln(w, "User not found!")
}

func main() {
	http.HandleFunc("/user", userHandler)
	http.ListenAndServe(":8080", nil)
}
