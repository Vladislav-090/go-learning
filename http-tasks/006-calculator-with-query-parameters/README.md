# Task 006 - Calculator With Query Parameters

## Goal

Create an HTTP calculator that reads numbers from query parameters and returns their sum.

## Concepts

- HTTP server
- Handler
- http.ResponseWriter
- *http.Request
- Query parameters
- r.URL.Query()
- Get()
- strconv.Atoi
- String to int conversion
- error handling
- fmt.Fprintln

## What I Learned

- Query parameters from URL are received as strings.
- `r.URL.Query().Get("a")` reads the `a` query parameter.
- `r.URL.Query().Get("b")` reads the `b` query parameter.
- `strconv.Atoi()` converts a string into an integer.
- Strings must be converted to numbers before mathematical operations.
- `"10" + "20"` would produce `"1020"`, but `10 + 20` produces `30`.
- `return` stops handler execution after sending an error response.

## Features

- Register `/calc` route
- Read `a` query parameter
- Read `b` query parameter
- Validate empty parameters
- Convert string values to integers
- Calculate sum
- Return result to browser

## Verification

1. Run the application:

go run main.go

2. Open:

http://localhost:8080/calc?a=10&b=20

Expected response:

Summ of numbers is: 30

3. Open:

http://localhost:8080/calc?a=10

Expected response:

b is empty

4. Open:

http://localhost:8080/calc?a=hello&b=20

Expected response:

invalid number

## Run

go run main.go