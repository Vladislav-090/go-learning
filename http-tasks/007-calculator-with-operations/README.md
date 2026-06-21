# Task 007 - Calculator With Operations

## Goal

Create an HTTP calculator that supports different operations using query parameters.

## Concepts

- HTTP server
- Handler
- http.ResponseWriter
- *http.Request
- Query parameters
- r.URL.Query()
- Get()
- strconv.Atoi
- switch
- error handling
- division by zero validation

## What I Learned

- Query parameters are received as strings.
- `strconv.Atoi()` converts string values to integers.
- `switch` can be used to choose logic based on the `op` query parameter.
- Query parameter `op` can define the operation type.
- Division by zero should be checked before performing division.
- `return` stops handler execution after sending an error response.
- `fmt.Fprintln(w, ...)` sends the result back to the browser.

## Features

- Register `/calc` route
- Read `a`, `b`, and `op` query parameters
- Validate empty parameters
- Convert query values to integers
- Support addition
- Support subtraction
- Support multiplication
- Support division
- Handle unknown operation
- Handle division by zero

## Verification

1. Run the application:

go run main.go

2. Addition:

http://localhost:8080/calc?a=10&b=20&op=sum

Expected response:

Result: 30

3. Subtraction:

http://localhost:8080/calc?a=50&b=20&op=sub

Expected response:

Result: 30

4. Multiplication:

http://localhost:8080/calc?a=5&b=6&op=mul

Expected response:

Result: 30

5. Division:

http://localhost:8080/calc?a=60&b=2&op=div

Expected response:

Result: 30

6. Division by zero:

http://localhost:8080/calc?a=10&b=0&op=div

Expected response:

division by zero

7. Unknown operation:

http://localhost:8080/calc?a=10&b=20&op=test

Expected response:

Uknown operation

## Run

go run main.go