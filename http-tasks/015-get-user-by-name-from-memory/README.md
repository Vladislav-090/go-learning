# Task 015 - Get User By Name

## Goal

Find a user in memory by name and return the user as JSON.

## Concepts

- HTTP GET
- Query Parameters
- Slice
- Range
- Search
- JSON
- MarshalIndent
- Content-Type

## What I Learned

- Query parameters can be read using:
  r.URL.Query().Get()

- A user can be searched in a slice using range.

- MarshalIndent converts a struct into JSON.

- return stops handler execution immediately after sending a response.

- Content-Type application/json tells the client that JSON is returned.

## Features

- Accept GET requests only
- Read name from query parameters
- Search user in memory
- Return user as JSON
- Return error if user not found

## Verification

Add user:

Invoke-RestMethod -Uri "http://localhost:8080/addUser" `
  -Method POST `
  -ContentType "application/json" `
  -Body '{"name":"Max","age":10}'

Find user:

http://localhost:8080/getUser?name=Max

Expected:

{
  "name": "Max",
  "age": 10
}

User not found:

http://localhost:8080/getUser?name=John

Expected:

user not found

Empty name:

http://localhost:8080/getUser

Expected:

name is empty

## Run

go run main.go