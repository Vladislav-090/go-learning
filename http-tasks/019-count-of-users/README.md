# Task 019 - Count Users

## Goal

Create an HTTP endpoint that returns the number of users stored in memory.

## Concepts

- HTTP GET
- Handler
- In-memory storage
- Slice
- len()
- http.MethodGet
- ResponseWriter

## What I Learned

- `len(users)` returns the number of users in the slice.
- A handler can read data from global in-memory storage.
- GET requests are used to retrieve data.
- `fmt.Fprintln(w, ...)` sends the response to the browser.
- Data exists only while the server is running.

## Features

- Add users with POST `/addUser`
- Count users with GET `/countUsers`
- Return total number of users
- Reject non-GET requests

## Verification

Run the application:

go run main.go

Add users:

Invoke-RestMethod -Uri "http://localhost:8080/addUser" `
  -Method POST `
  -ContentType "application/json" `
  -Body '{"name":"Vladislav","age":29}'

Invoke-RestMethod -Uri "http://localhost:8080/addUser" `
  -Method POST `
  -ContentType "application/json" `
  -Body '{"name":"Viola","age":20}'

Invoke-RestMethod -Uri "http://localhost:8080/addUser" `
  -Method POST `
  -ContentType "application/json" `
  -Body '{"name":"Max","age":10}'

Open:

http://localhost:8080/countUsers

Expected:

Users count: 3

## Run

go run main.go