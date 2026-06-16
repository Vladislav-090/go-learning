# Task 004 - User Age Validation With Error

## Goal

Create a user age validation system with error handling.

## Concepts

- Structs
- Methods
- Pointer Receiver
- errors.New
- error return value
- nil
- Multiple return values
- bool and error
- if err != nil
- Age validation

## What I Learned

- A method can return more than one value.
- `IsAdult()` returns both `bool` and `error`.
- `nil` means there is no error.
- `err != nil` means an error happened.
- Validation should happen before updating age.
- Invalid age values should return an error.
- Being under 18 is not an error, it is a valid result.

## Features

- Update user age
- Validate negative age
- Validate too large age
- Check if user is adult
- Print user information

## Run

```bash
go run main.go