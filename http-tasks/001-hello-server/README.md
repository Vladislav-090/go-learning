# Task 001 - Hello Server

## Goal

Create a simple HTTP server in Go.

## Concepts

- HTTP server
- net/http
- Handler
- http.ResponseWriter
- *http.Request
- http.HandleFunc
- http.ListenAndServe
- localhost
- port 8080

## What I Learned

- A Go program can work as an HTTP server.
- A handler is a function that processes an HTTP request.
- `http.ResponseWriter` is used to write a response to the client.
- `*http.Request` contains information about the incoming request.
- `http.HandleFunc()` connects a URL path with a handler.
- `http.ListenAndServe()` starts the server on a selected port.
- `fmt.Fprintln(w, ...)` writes data to the HTTP response.
- `fmt.Println(...)` writes data to the terminal.

## Features

- Start HTTP server
- Register `/hello` route
- Return `Hello World!` response
- Open endpoint in browser

## Verification

1. Run the application:

go run main.go

2. Open browser:

http://localhost:8080/hello

3. Expected response:

Hello World!

## Run

go run main.go