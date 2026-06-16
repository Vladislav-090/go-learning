# Task 001 - Wallet Withdraw With Error

## Goal

Create a wallet that supports deposit and withdraw operations using error handling.

## Concepts

- Structs
- Methods
- Pointer Receiver
- errors.New
- error return value
- nil
- if err != nil
- Balance validation

## What I Learned

- A method can return an `error`.
- `errors.New()` creates a new error.
- `nil` means there is no error.
- `err != nil` means an error happened.
- Validation should happen before changing the balance.
- Business logic can return errors instead of printing messages directly.

## Features

- Deposit money
- Withdraw money
- Prevent negative or zero deposit
- Prevent negative or zero withdraw
- Prevent withdraw when balance is not enough
- Get final balance

## Run

```bash
go run main.go