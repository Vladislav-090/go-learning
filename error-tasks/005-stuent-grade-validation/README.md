# Task 005 - Student Grade Validation With Error

## Goal

Create a student grade validation system with error handling.

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
- Grade validation

## What I Learned

- Methods can return errors when an operation is invalid.
- `nil` means there is no error.
- `err != nil` means an error happened.
- A method can return both `bool` and `error`.
- Validation should happen before changing grade.
- Future values should be checked before updating data.
- Grade cannot be less than 0 or greater than 100.

## Features

- Update grade
- Increase grade
- Prevent negative grade
- Prevent grade greater than 100
- Check if student is excellent
- Print student information

## Run

```bash
go run main.go