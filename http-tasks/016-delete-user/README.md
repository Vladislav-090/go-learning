# Task 016 - Delete User From Memory

## Goal

Create an HTTP endpoint that deletes a user from memory by name.

## Concepts

- HTTP DELETE
- Query parameters
- r.URL.Query().Get()
- Slice
- Range with index
- append
- Delete from slice
- In-memory storage
- return

## What I Learned

- DELETE requests can be used to remove data.
- Query parameters can be used to select which user to delete.
- To delete an element from a slice, the index is required.
- `for i, user := range users` gives both index and value.
- `users = append(users[:i], users[i+1:]...)` removes an element from a slice.
- `return` stops handler execution after deleting a user.
- `User not found` should be returned only after checking the whole slice.

## Features

- Add user with POST `/addUser`
- Find user with GET `/getUser?name=...`
- Delete user with DELETE `/deleteUser?name=...`
- Validate empty name
- Handle user not found
- Store users in memory

## Verification

1. Run the application:

go run main.go

2. Add user:

Invoke-RestMethod -Uri "http://localhost:8080/addUser" `
  -Method POST `
  -ContentType "application/json" `
  -Body '{"name":"Max","age":10}'

Expected:

User added succsessfully!

3. Find user:

http://localhost:8080/getUser?name=Max

Expected:

{
  "name": "Max",
  "age": 10
}

4. Delete user:

Invoke-RestMethod -Uri "http://localhost:8080/deleteUser?name=Max" `
  -Method DELETE

Expected:

Student deleted!

5. Try to find deleted user:

http://localhost:8080/getUser?name=Max

Expected:

user not found

## Run

go run main.go