# Task 022 - Status Codes

## Goal

Improve HTTP error responses by using proper HTTP status codes.

## Concepts

- HTTP status codes
- http.Error
- StatusBadRequest
- StatusMethodNotAllowed
- StatusConflict
- JSON Decode
- Validation
- Duplicate check

## What I Learned

- `http.Error()` sends both an error message and an HTTP status code.
- `http.StatusBadRequest` is used for invalid client data.
- `http.StatusMethodNotAllowed` is used for wrong HTTP methods.
- `http.StatusConflict` is used when a resource already exists.
- After `http.Error()`, `return` is required to stop handler execution.
- Without `return`, the handler may continue and still add invalid data.

## Features

- Reject non-POST requests with 405
- Reject invalid JSON with 400
- Reject empty name with 400
- Reject invalid age with 400
- Reject duplicate users with 409
- Add valid users successfully

## Verification

Run:

go run main.go

Valid user:

Invoke-RestMethod -Uri "http://localhost:8080/addUser" `
  -Method POST `
  -ContentType "application/json" `
  -Body '{"name":"Max","age":10}'

Expected:

user succsessfully added!

Duplicate user:

Invoke-RestMethod -Uri "http://localhost:8080/addUser" `
  -Method POST `
  -ContentType "application/json" `
  -Body '{"name":"Max","age":10}'

Expected:

409 Conflict

Empty name:

Invoke-RestMethod -Uri "http://localhost:8080/addUser" `
  -Method POST `
  -ContentType "application/json" `
  -Body '{"name":"","age":10}'

Expected:

400 Bad Request

Wrong method:

Open:

http://localhost:8080/addUser

Expected:

405 Method Not Allowed

## Run

go run main.go