# Task 006 - Library Borrow Book With Error

## Goal

Create a simple library book system with error handling.

## Concepts

- Structs
- Methods
- Boolean validation
- errors.New
- error return value
- nil
- if err != nil

## What I Learned

- Business rules can be validated with errors.
- A book cannot be borrowed twice.
- A book cannot be returned if it is already available.
- Boolean fields can represent object state.
- Errors prevent invalid operations.

## Features

- Borrow a book
- Return a book
- Prevent duplicate borrowing
- Prevent duplicate returning
- Print book information

## Run

```bash
go run main.go