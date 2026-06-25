# Task 025 - Helper Functions For JSON Responses

## Goal

Create helper functions for JSON success and error responses.

## Concepts

- Helper functions
- JSON response
- ErrorResponse
- SuccessResponse
- json.NewEncoder
- http.ResponseWriter
- Status codes
- DRY principle

## What I Learned

- Repeated response code can be moved into helper functions.
- `writeJSON()` can set Content-Type, status code, and encode data.
- `writeError()` can create an ErrorResponse and send it through writeJSON.
- `writeSuccess()` can create a SuccessResponse and send it through writeJSON.
- Handlers become cleaner when response logic is moved into helper functions.
- User can be passed by value when the function only reads it and does not modify it.
- Pointers are needed when a function must modify the original value.

## Features

- JSON success response
- JSON error response
- Helper for writing JSON
- Helper for writing errors
- Helper for writing success
- Duplicate user validation
- Proper HTTP status codes

## Verification

Run:

go run main.go

Add user:

Invoke-RestMethod -Uri "http://localhost:8080/addUser" `
  -Method POST `
  -ContentType "application/json" `
  -Body '{"name":"Max","age":10}'

Expected:

{
  "message": "User added successfully",
  "user": {
    "name": "Max",
    "age": 10
  }
}

Duplicate user:

Invoke-RestMethod -Uri "http://localhost:8080/addUser" `
  -Method POST `
  -ContentType "application/json" `
  -Body '{"name":"Max","age":10}'

Expected:

{
  "error": "User already exist"
}

## Run

go run main.go