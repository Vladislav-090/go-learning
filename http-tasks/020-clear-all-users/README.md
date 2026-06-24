# Task 020 - Clear All Users

## Goal

Create an HTTP endpoint that deletes all users from memory.

## Concepts

- HTTP DELETE
- Handler
- Slice
- In-memory storage
- Clearing slice
- http.MethodDelete
- ResponseWriter

## What I Learned

- DELETE requests can be used to remove data.
- A slice can be cleared by assigning an empty slice.
- `users = []User{}` replaces the current users list with an empty list.
- Handlers are not called directly.
- The client sends an HTTP request, and Go calls the matching handler.
- Data in memory exists only while the server is running.

## Features

- Add users with POST `/addUser`
- Get users with GET `/getUsers`
- Clear all users with DELETE `/clearUsers`
- Reject non-DELETE requests

## Verification

Run:

go run main.go

Add users:

Invoke-RestMethod -Uri "http://localhost:8080/addUser" `
  -Method POST `
  -ContentType "application/json" `
  -Body '{"name":"Max","age":10}'

Invoke-RestMethod -Uri "http://localhost:8080/addUser" `
  -Method POST `
  -ContentType "application/json" `
  -Body '{"name":"Viola","age":20}'

Open:

http://localhost:8080/getUsers

Expected:

[
  {
    "name": "Max",
    "age": 10
  },
  {
    "name": "Viola",
    "age": 20
  }
]

Clear users:

Invoke-RestMethod -Uri "http://localhost:8080/clearUsers" `
  -Method DELETE

Expected:

All Users deleted!

Open again:

http://localhost:8080/getUsers

Expected:

[]