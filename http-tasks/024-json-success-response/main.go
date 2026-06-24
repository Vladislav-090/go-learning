package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"success"`
	User    User   `json:"user"`
}

var users []User

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	if r.Method != http.MethodPost {
		errorResponse := ErrorResponse{
			Error: "method not allowed!",
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

	for _, existUser := range users {
		if existUser.Name == user.Name {
			errorResponse := ErrorResponse{
				Error: "user Already exist!",
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(errorResponse)
			return
		}
	}
	users = append(users, user)
	successResponse := SuccessResponse{
		Message: "User added successfully!",
		User:    user,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(successResponse)
}

func main() {
	http.HandleFunc("/addUser", addUserHandler)

	http.ListenAndServe(":8080", nil)
}
