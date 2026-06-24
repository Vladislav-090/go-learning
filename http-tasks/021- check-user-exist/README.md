# Task 021 - Prevent Duplicate Users

## Goal

Prevent adding a user if a user with the same name already exists in memory.

## Concepts

- HTTP POST
- JSON
- Request Body
- json.NewDecoder
- Slice
- Range
- Duplicate check
- Validation
- return

## What I Learned

- Before adding a new user, existing users should be checked.
- `for _, existingUser := range users` can be used to search through stored users.
- `existingUser.Name == user.Name` checks whether a user already exists.
- Duplicate validation should happen before `append`.
- `return` stops the handler after sending an error response.

## Features

- Add user with POST `/addUser`
- Validate JSON
- Validate name
- Validate age
- Prevent duplicate users
- Return error when user already exists

## Verification

Run:

go run main.go

Add user first time:

Invoke-RestMethod -Uri "http://localhost:8080/addUser" `
  -Method POST `
  -ContentType "application/json" `
  -Body '{"name":"Max","age":10}'

Expected:

User added successfully!

Add same user again:

Invoke-RestMethod -Uri "http://localhost:8080/addUser" `
  -Method POST `
  -ContentType "application/json" `
  -Body '{"name":"Max","age":10}'

Expected:

user already exist!

## Run

go run main.go