# Task 023 - JSON Error Responses

## Goal

Return error responses in JSON format instead of plain text.

## Concepts

- HTTP errors
- JSON response
- ErrorResponse struct
- json.NewEncoder
- Encode
- Content-Type
- WriteHeader
- Status codes

## What I Learned

- API errors can be returned as JSON objects.
- `ErrorResponse` can be used as a standard error response structure.
- `json.NewEncoder(w).Encode(...)` writes JSON directly to the HTTP response.
- `w.WriteHeader(...)` sets the HTTP status code.
- `Content-Type: application/json` tells the client that the response body is JSON.
- `return` is required after writing an error response.

## Features

- Return JSON error for wrong method
- Return JSON error for invalid JSON
- Return JSON error for empty name
- Return JSON error for invalid age
- Return JSON error for duplicate user
- Use proper HTTP status codes

## Verification

Run:

go run main.go

Empty name:

Invoke-RestMethod -Uri "http://localhost:8080/addUser" `
  -Method POST `
  -ContentType "application/json" `
  -Body '{"name":"","age":10}'

Expected:

{
  "error": "Name is empty!"
}

Invalid age:

Invoke-RestMethod -Uri "http://localhost:8080/addUser" `
  -Method POST `
  -ContentType "application/json" `
  -Body '{"name":"Max","age":0}'

Expected:

{
  "error": "Age must be positive!"
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