# Task 011 - Allow Only POST

## Goal

Allow only POST requests for creating a user from JSON.

## Concepts

- HTTP server
- Handler
- POST request
- HTTP method validation
- r.Method
- http.MethodPost
- Request body
- JSON
- json.NewDecoder
- Decode
- return

## What I Learned

- `r.Method` contains the HTTP method of the request.
- `http.MethodPost` is a Go constant for `"POST"`.
- Handler can reject requests with the wrong HTTP method.
- `return` stops handler execution after sending an error response.
- `r.Body` contains JSON data sent by the client.
- `json.NewDecoder(r.Body).Decode(&user)` reads JSON from request body into a struct.

## Features

- Register `/user` route
- Allow only POST requests
- Reject GET requests
- Decode JSON body into User struct
- Return created user name

## Verification

1. Run the application:

go run main.go

2. Open in browser:

http://localhost:8080/user

Expected response:

method not allowed

3. Test POST with PowerShell:

Invoke-RestMethod -Uri "http://localhost:8080/user" `
  -Method POST `
  -ContentType "application/json" `
  -Body '{"name":"Vladislav","age":29}'

Expected response:

User created Vladislav

## Run

go run main.go