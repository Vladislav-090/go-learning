# Task 014 - User Registration Validation With Error

## Goal

Create a user registration validation system with error handling.

## Concepts

- Structs
- Methods
- Pointer Receiver
- errors.New
- error return value
- nil
- if err != nil
- String validation
- len()
- strings.Contains()

## What I Learned

- User input should be validated before registration.
- Empty username, email, and password values should return errors.
- Username length can be checked with `len()`.
- Password length can be checked with `len()`.
- Email format can be partially validated with `strings.Contains()`.
- `!strings.Contains(email, "@")` checks that email does not contain `@`.
- Validation order matters because the function returns the first error it finds.

## Features

- Validate empty username
- Validate short username
- Validate empty email
- Validate email contains `@`
- Validate empty password
- Validate short password
- Successful registration check
- Print user registration data

## Run

```bash
go run main.go