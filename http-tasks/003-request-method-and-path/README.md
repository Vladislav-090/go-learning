# Task 003 - Request Method And Path

## Goal

Read information from an HTTP request and return it to the client.

## Concepts

- HTTP server
- Handler
- http.ResponseWriter
- *http.Request
- Request Method
- Request URL
- Route
- net/http

## What I Learned

- `r.Method` contains the HTTP method.
- `r.URL.Path` contains the request path.
- A handler can read data from an incoming request.
- `fmt.Fprintln(w, ...)` sends data back to the client.
- The request object contains information about the client request.

## Features

- Register `/info` route
- Read request method
- Read request path
- Return request information to the browser

## Verification

1. Run the application:

go run main.go

2. Open:

http://localhost:8080/info

Expected response:

Method: GET
Path: /info

## Run

go run main.go