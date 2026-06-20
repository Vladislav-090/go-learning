# Task 004 - Query Parameters

## Goal

Read a query parameter from an HTTP request and return a personalized response.

## Concepts

- HTTP server
- Handler
- http.ResponseWriter
- *http.Request
- Query parameters
- r.URL.Query()
- Get()
- fmt.Fprintln
- net/http

## What I Learned

- Query parameters are passed after `?` in the URL.
- `/user?name=Vladislav` contains query parameter `name`.
- `r.URL.Query().Get("name")` reads the `name` parameter from the request.
- `r` is used to read data from the incoming request.
- `w` is used to write the response back to the browser.
- `fmt.Fprintln(w, ...)` writes data to the HTTP response.
- `fmt.Println(...)` writes data only to the terminal.
- `return` stops handler execution after sending an error response.

## Features

- Register `/user` route
- Read `name` query parameter
- Return greeting with name
- Return error message when name is empty

## Verification

1. Run the application:

go run main.go

2. Open:

http://localhost:8080/user?name=Vladislav

Expected response:

Hello Vladislav

3. Open:

http://localhost:8080/user

Expected response:

name is empty!

## Run

go run main.go