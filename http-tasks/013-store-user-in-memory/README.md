# Task 013 - Store User In Memory

## Goal

Create a POST endpoint that receives a user from JSON and stores it in memory.

## Concepts

- HTTP POST
- Request Body
- JSON
- json.NewDecoder
- Validation
- Slice
- append
- Global variables
- In-memory storage

## What I Learned

- A global slice can be used as temporary application storage.
- Data stored in a global slice remains available while the server is running.
- `append()` adds a new element to a slice.
- `json.NewDecoder(r.Body).Decode(&user)` reads JSON into a struct.
- Validation should be performed before storing data.
- `r.Method` can be used to restrict HTTP methods.

## Features

- Accept POST requests only
- Decode JSON into User struct
- Validate user name
- Validate user age
- Store user in memory
- Return success response

## Storage

Users are stored in:

var users []User

Data exists only while the application is running.

## Verification

Valid request:

Invoke-RestMethod -Uri "http://localhost:8080/user" `
  -Method POST `
  -ContentType "application/json" `
  -Body '{"name":"Vladislav","age":29}'

Expected:

User added successfully

Empty name:

{"name":"","age":29}

Expected:

Empty name

Invalid age:

{"name":"Vladislav","age":0}

Expected:

age must be positive!

Invalid method:

GET /user

Expected:

method not allowed

## Run

go run main.go