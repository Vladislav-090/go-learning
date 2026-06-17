# Task 013 - Order Payment Validation With Error

## Goal

Create an order payment system with error handling.

## Concepts

- Structs
- Methods
- Pointer Receiver
- Pointers
- errors.New
- error return value
- nil
- if err != nil
- Business validation

## What I Learned

- A pointer can be used to modify an external value.
- `*balance` accesses the value stored at the address.
- `&balance` gets the address of a variable.
- `nil` means a pointer does not reference any value.
- An order cannot be paid twice.
- An order amount must be positive.
- Payment requires sufficient balance.
- An unpaid order cannot be canceled.

## Features

- Pay order
- Cancel order
- Validate order amount
- Validate balance availability
- Validate nil pointer
- Prevent double payment
- Print order information

## Run

```bash
go run main.go