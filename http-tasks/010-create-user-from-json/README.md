# Task 010 - Create User From JSON

## Goal

Create an HTTP endpoint that reads JSON from request body and converts it into a Go struct.

## Concepts

- HTTP server
- POST request
- Request body
- JSON
- Struct
- JSON tags
- json.NewDecoder
- Decode
- *http.Request
- http.ResponseWriter

## What I Learned

- `r.Body` contains data sent by the client.
- JSON from request body can be decoded into a Go struct.
- `json.NewDecoder(r.Body).Decode(&user)` reads JSON and fills the struct.
- `&user` is required because Decode needs to modify the struct.
- POST requests are used to send data to the server.
- PowerShell can be used to test POST requests.

## Verification

Run:

go run main.go

Test with PowerShell:

Invoke-RestMethod -Uri "http://localhost:8080/user" `
  -Method POST `
  -ContentType "application/json" `
  -Body '{"name":"Vladislav","age":29}'

Expected response:

User created! Vladislav

## Run

go run main.go