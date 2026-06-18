# Task 015 - Blockchain Transaction Validation With Error

## Goal

Create a blockchain transaction validation system with error handling.

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
- Transaction validation

## What I Learned

- Transaction data should be validated before execution.
- `Validate` checks whether a transaction can be executed.
- `Execute` performs the transaction and changes balance.
- A pointer can be used to modify an external balance.
- `nil` pointer should be validated before dereferencing.
- Transaction amount and fee should be validated.
- Sender and receiver addresses should be different.
- Failed transactions should not change balance or success status.

## Features

- Validate sender address
- Validate receiver address
- Prevent sending to the same address
- Validate amount
- Validate fee
- Validate balance
- Execute transaction
- Update transaction status
- Print transaction information

## Run

```bash
go run main.go