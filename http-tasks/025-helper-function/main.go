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

type SuccsessResponse struct {
	Message string `json:"message"`
	User    User   `json:"user"`
}

var users []User

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func WriteError(w http.ResponseWriter, status int, message string) {
	errorResponse := ErrorResponse{
		Error: message,
	}
	WriteJSON(w, status, errorResponse)
}

func WriteSuccess(w http.ResponseWriter, status int, message string, user User) {

	successResponse := SuccsessResponse{
		Message: "User added successfully!",
		User:    user,
	}
	WriteJSON(w, status, successResponse)
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if r.Method != http.MethodPost {
		WriteError(w, http.StatusMethodNotAllowed, "Method not allowed!")
		return
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		WriteError(w, http.StatusBadRequest, "Invalid JSON!")
		return
	}

	if user.Name == "" {
		WriteError(w, http.StatusBadRequest, "Name is empty!")
		return
	}

	if user.Age <= 0 {
		WriteError(w, http.StatusBadRequest, "Age must be positive!")
		return
	}

	for _, existUser := range users {
		if existUser.Name == user.Name {
			WriteError(w, http.StatusConflict, "User already exist!")
			return
		}
	}
	users = append(users, user)
	WriteSuccess(w, http.StatusCreated, "User added succsessfully", user)
}

func main() {
	http.HandleFunc("/addUser", addUserHandler)

	http.ListenAndServe(":8080", nil)
}
