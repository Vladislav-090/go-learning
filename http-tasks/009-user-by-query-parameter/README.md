# Task 009 - Get User By Name

## Goal

Create an HTTP endpoint that returns a specific user in JSON format using a query parameter.

## Concepts

- HTTP server
- Query parameters
- Slice
- Range
- JSON
- json.MarshalIndent
- Content-Type
- Search in slice
- return

## What I Learned

- Query parameters can be used to filter data.
- `r.URL.Query().Get("name")` reads the user name from the request.
- A slice can be searched using a `for range` loop.
- `json.MarshalIndent()` converts a struct into formatted JSON.
- `return` immediately stops handler execution.
- `User not found` should be returned only after checking all users.
- `Content-Type: application/json` tells the client that the response contains JSON.

## Features

- Register `/user` route
- Read `name` query parameter
- Search user by name
- Return user as JSON
- Return "Name is empty!" if parameter is missing
- Return "User not found!" if user does not exist

## Verification

1. Run the application:

go run main.go

2. Open:

http://localhost:8080/user?name=Vladislav

Expected response:

{
  "name": "Vladislav",
  "age": 29
}

3. Open:

http://localhost:8080/user?name=Viola

Expected response:

{
  "name": "Viola",
  "age": 20
}

4. Open:

http://localhost:8080/user?name=John

Expected response:

User not found!

5. Open:

http://localhost:8080/user

Expected response:

Name is empty!

## Run

go run main.go