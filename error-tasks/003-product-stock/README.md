# Task 003 - Product Stock With Error

## Goal

Create a product stock management system with error handling.

## Concepts

- Structs
- Methods
- Pointer Receiver
- errors.New
- error return value
- nil
- if err != nil
- Stock validation

## What I Learned

- Methods can return errors when an operation is invalid.
- `errors.New()` creates a new error.
- `nil` means there is no error.
- `err != nil` means an error happened.
- Validation should happen before changing stock.
- Stock cannot become negative.
- Stock can be updated with a new value.
- Failed operations should not change product stock.

## Features

- Add stock
- Remove stock
- Update stock
- View product information
- Prevent negative stock
- Prevent removing more stock than available

## Run

```bash
go run main.go