# Task 024 - JSON Success Response

## Goal

Return successful API responses in JSON format.

## Concepts

- JSON success response
- SuccessResponse struct
- ErrorResponse struct
- json.NewEncoder
- HTTP status codes
- Content-Type
- POST request
- Validation

## What I Learned

- Successful responses can also be returned as JSON.
- `SuccessResponse` can contain a message and the created user.
- `json.NewEncoder(w).Encode(...)` writes JSON directly to the response.
- `http.StatusCreated` means that a resource was created successfully.
- APIs should return consistent JSON responses for both success and error cases.

## Features

- Add user with POST `/addUser`
- Return JSON error responses
- Return JSON success response
- Include created user in success response
- Use proper HTTP status codes

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
  "message": "User added successfully!",
  "user": {
    "name": "Max",
    "age": 10
  }
}

Add duplicate user:

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