# Task 017 - Update User Age

## Goal

Create an HTTP endpoint that updates the age of an existing user.

## Concepts

- HTTP PUT
- Query Parameters
- strconv.Atoi
- Slice
- Update slice element
- Validation
- return

## What I Learned

- PUT requests can be used to update existing data.
- Query parameters can be read using:
  r.URL.Query().Get()
- strconv.Atoi converts a string to an integer.
- Slice elements can be modified by index.
- Range can provide both index and value:
  for i, user := range users
- Data validation should be performed before updating data.

## Features

- Add user with POST /addUser
- Find user with GET /getUser
- Delete user with DELETE /deleteUser
- Update user age with PUT /updateUser
- Validate name
- Validate age
- Handle user not found

## Verification

Add user:

Invoke-RestMethod -Uri "http://localhost:8080/addUser" `
  -Method POST `
  -ContentType "application/json" `
  -Body '{"name":"Max","age":10}'

Expected:

User added successfully!

Update age:

Invoke-RestMethod -Uri "http://localhost:8080/updateUser?name=Max&age=15" `
  -Method PUT

Expected:

Age has been updated!

Check user:

http://localhost:8080/getUser?name=Max

Expected:

{
  "name": "Max",
  "age": 15
}

Empty name:

http://localhost:8080/updateUser?age=15

Expected:

name is empty

Invalid age:

http://localhost:8080/updateUser?name=Max&age=0

Expected:

age must be positive!

User not found:

http://localhost:8080/updateUser?name=John&age=15

Expected:

user not found

## Run

go run main.go