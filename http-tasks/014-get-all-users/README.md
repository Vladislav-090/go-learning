# Task 014 - Get All Users From Memory

## Goal

Create an HTTP endpoint that returns all users stored in memory.

## Concepts

- HTTP GET
- HTTP POST
- Handler
- Request Body
- JSON
- json.NewDecoder
- json.MarshalIndent
- Slice
- append
- In-memory storage
- Content-Type

## What I Learned

- One handler can add data to a global slice.
- Another handler can read data from the same global slice.
- `POST /user` can be used to create and store a user.
- `GET /users` can be used to return all stored users.
- `json.MarshalIndent()` converts a slice of structs into a JSON array.
- `Content-Type: application/json` tells the client that the response is JSON.
- Data stored in memory exists only while the server is running.

## Features

- Add user with POST `/user`
- Validate user name
- Validate user age
- Store users in memory
- Get all users with GET `/users`
- Return users as JSON array

## Verification

1. Run the application:

go run main.go

2. Add first user:

Invoke-RestMethod -Uri "http://localhost:8080/user" `
  -Method POST `
  -ContentType "application/json" `
  -Body '{"name":"Vladislav","age":29}'

Expected:

user added succsessfully

3. Add second user:

Invoke-RestMethod -Uri "http://localhost:8080/user" `
  -Method POST `
  -ContentType "application/json" `
  -Body '{"name":"Viola","age":20}'

Expected:

user added succsessfully

4. Open:

http://localhost:8080/users

Expected:

[
  {
    "name": "Vladislav",
    "age": 29
  },
  {
    "name": "Viola",
    "age": 20
  }
]

## Run

go run main.go