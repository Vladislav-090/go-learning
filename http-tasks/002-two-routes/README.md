# Task 002 - Two Routes

## Goal

Create an HTTP server with two different routes.

## Concepts

- HTTP server
- net/http
- Handler
- http.ResponseWriter
- *http.Request
- http.HandleFunc
- http.ListenAndServe
- Routes
- localhost
- port 8080

## What I Learned

- One HTTP server can have multiple routes.
- Each route can have its own handler function.
- `http.HandleFunc()` connects a path with a handler.
- `/hello` can call `helloHandler`.
- `/users` can call `usersHandler`.
- `fmt.Fprintln(w, ...)` writes a response to the browser.
- `*http.Request` contains data about the incoming request.

## Features

- Start HTTP server
- Register `/hello` route
- Register `/users` route
- Return different responses for different routes

## Verification

1. Run the application:

go run main.go

2. Open:

http://localhost:8080/hello

Expected response:

Hello World!

3. Open:

http://localhost:8080/users

Expected response:

Users page

## Run

go run main.go