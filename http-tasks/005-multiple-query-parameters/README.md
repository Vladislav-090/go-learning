# Task 005 - Multiple Query Parameters

## Goal

Read multiple query parameters from an HTTP request.

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

- A URL can contain multiple query parameters.
- Query parameters are separated with `&`.
- `/greet?name=Vladislav&age=29` contains two parameters: `name` and `age`.
- `r.URL.Query().Get("name")` reads the `name` parameter.
- `r.URL.Query().Get("age")` reads the `age` parameter.
- `r` is used to read data from the incoming request.
- `w` is used to write the response back to the browser.
- `return` stops handler execution after sending an error response.

## Features

- Register `/greet` route
- Read `name` query parameter
- Read `age` query parameter
- Return greeting with name and age
- Return error message when name is empty
- Return error message when age is empty

## Verification

1. Run the application:

go run main.go

2. Open:

http://localhost:8080/greet?name=Vladislav&age=29

Expected response:

Hello Vladislav, age 29

3. Open:

http://localhost:8080/greet?name=Vladislav

Expected response:

Empty age

4. Open:

http://localhost:8080/greet?age=29

Expected response:

Empty name

## Run

go run main.go