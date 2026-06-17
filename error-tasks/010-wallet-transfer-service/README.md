# Task 010 - Wallet Transfer Service With Error

## Goal

Create a wallet transfer service with error handling.

## Concepts

- Structs
- Methods
- Pointer Receiver
- errors.New
- error return value
- nil
- if err != nil
- Method reuse
- Balance validation
- Nil pointer validation

## What I Learned

- A wallet can deposit, withdraw, and transfer money.
- `Transfer` can reuse `Withdraw` and `Deposit`.
- `nil` can be used to check if the receiver wallet exists.
- A method receiver can represent the sender wallet.
- Invalid transfer operations should return errors.
- Failed transfers should not change balances.
- Business logic should validate data before updating state.

## Features

- Deposit money
- Withdraw money
- Transfer money between wallets
- Prevent transfer to nil wallet
- Prevent zero or negative amount
- Prevent transfer when balance is not enough
- Print wallet information

## Run

```bash
go run main.go