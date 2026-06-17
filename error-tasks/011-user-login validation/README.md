# Task 011 - User Login Validation With Error

## Goal

Create a user login validation system with error handling.

## Concepts

- Structs
- Methods
- Pointer Receiver
- Boolean validation
- errors.New
- error return value
- nil
- if err != nil
- Input validation
- Login validation

## What I Learned

- Login data should be validated before successful authentication.
- Empty username and password values should return errors.
- Incorrect username and password values should return errors.
- Inactive users should not be allowed to log in.
- Validation order matters.
- `nil` means the login operation completed successfully.
- `err != nil` means the login operation failed.

## Features

- Validate empty username
- Validate empty password
- Validate incorrect username
- Validate incorrect password
- Validate inactive user
- Successful login check
- Print user information

## Run

```bash
go run main.go