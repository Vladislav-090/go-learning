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

type ErrorResponse struct {
	Error string `json:"error"`
}

var users []User

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	if r.Method != http.MethodPost {
		errorResponse := ErrorResponse{
			Error: "Method not allowed!",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		errorResponse := ErrorResponse{
			Error: "Invalid json!",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	if user.Name == "" {
		errorResponse := ErrorResponse{
			Error: "Name is empty!",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	if user.Age <= 0 {
		errorResponse := ErrorResponse{
			Error: "Age must be positive!",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	for _, userExist := range users {
		if userExist.Name == user.Name {
			errorResponse := ErrorResponse{
				Error: "User already exist",
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(errorResponse)
			return
		}
	}
	users = append(users, user)
	fmt.Fprintln(w, "User successfully added!")
}

func main() {
	http.HandleFunc("/addUser", addUserHandler)

	http.ListenAndServe(":8080", nil)
}
