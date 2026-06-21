# Task 008 - JSON Users Response

## Goal

Create an HTTP endpoint that returns a JSON array of users.

## Concepts

- HTTP server
- Handler
- JSON
- Struct
- Slice
- json.MarshalIndent
- Response Headers
- Content-Type
- application/json

## What I Learned

- HTTP handlers can return JSON responses.
- `json.MarshalIndent()` converts Go structures into formatted JSON.
- A slice of structs can be converted into a JSON array.
- `w.Header().Set()` allows setting response headers.
- `Content-Type: application/json` tells the client that the response contains JSON.
- `fmt.Fprintln(w, ...)` sends data back to the browser.

## Features

- Register `/users` route
- Create users slice
- Convert users slice to JSON
- Return JSON response
- Set JSON content type

## Verification

1. Run the application:

go run main.go

2. Open:

http://localhost:8080/users

Expected response:

[
  {
    "name": "Vladislav",
    "age": 29
  },
  {
    "name": "Viola",
    "age": 30
  }
]

## Run

go run main.go