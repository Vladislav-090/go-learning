# Task 002 - Bank Account Transfer With Error

## Goal

Create bank accounts with deposit, withdraw, and transfer operations using error handling.

## Concepts

- Structs
- Methods
- Pointer Receiver
- errors.New
- error return value
- nil
- if err != nil
- Business logic validation

## What I Learned

- Methods can return errors when an operation is invalid.
- `nil` means the operation completed successfully.
- `err != nil` means an error happened.
- Validation should happen before changing balance.
- One method can reuse another method.
- `Transfer` can use `Withdraw` and `Deposit` internally.
- Failed operations should not change account balances.

## Run

```bash
go run main.go