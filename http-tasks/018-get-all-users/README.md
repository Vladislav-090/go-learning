# Task 018 - Get All Users

## Goal

Create an HTTP endpoint that returns all users stored in memory.

## Concepts

- HTTP GET
- JSON
- json.MarshalIndent
- Slice
- Content-Type
- ResponseWriter

## What I Learned

- A slice can be converted directly into JSON.
- json.MarshalIndent can convert a slice of structs into a JSON array.
- A loop is not required when returning the entire slice.
- Content-Type should be set to application/json.
- GET requests are used to retrieve data.

## Features

- Return all users from memory
- Accept GET requests only
- Return JSON array
- Handle marshal errors

## Verification

Add users:

Invoke-RestMethod -Uri "http://localhost:8080/addUser" `
  -Method POST `
  -ContentType "application/json" `
  -Body '{"name":"Vladislav","age":29}'

Invoke-RestMethod -Uri "http://localhost:8080/addUser" `
  -Method POST `
  -ContentType "application/json" `
  -Body '{"name":"Viola","age":20}'

Invoke-RestMethod -Uri "http://localhost:8080/addUser" `
  -Method POST `
  -ContentType "application/json" `
  -Body '{"name":"Max","age":10}'

Open:

http://localhost:8080/users

Expected:

[
  {
    "name": "Vladislav",
    "age": 29
  },
  {
    "name": "Viola",
    "age": 20
  },
  {
    "name": "Max",
    "age": 10
  }
]

## Run

go run main.go