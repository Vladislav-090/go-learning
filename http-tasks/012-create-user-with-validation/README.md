# Task 012 - Create User With Validation

## Goal

Create a POST endpoint that validates user data before creating a user.

## Concepts

- HTTP POST
- Request Body
- JSON
- json.NewDecoder
- Validation
- if
- return
- r.Method
- http.MethodPost

## What I Learned

- Only POST requests should be accepted.
- JSON data is received through r.Body.
- json.NewDecoder(r.Body).Decode(&user) fills a struct from JSON.
- User data should be validated before processing.
- return stops handler execution.
- Error responses should be sent immediately.

## Validation Rules

- Name must not be empty
- Age must be greater than zero

## Features

- Reject GET requests
- Read JSON from request body
- Validate name
- Validate age
- Return validation errors
- Return success message

## Verification

POST valid user:

{"name":"Vladislav","age":29}

Expected:

User created! Name: Vladislav Age: 29

POST empty name:

{"name":"","age":29}

Expected:

Name is empty

POST invalid age:

{"name":"Vladislav","age":0}

Expected:

Age must be positive!

POST invalid json:

{"name":"Vladislav","age":}

Expected:

Invalid json

## Run

go run main.go