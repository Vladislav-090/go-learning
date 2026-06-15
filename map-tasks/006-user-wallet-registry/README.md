# Task 006 - User Wallet Registry With Map And Slice

## Goal

Create a user wallet registry that stores multiple wallets for each user.

## Concepts

- Interfaces
- Structs
- Methods
- Maps
- Slices
- map[string][]string
- append
- range
- len
- Delete item from slice
- Duplicate Check

## What I Learned

- A map can store a slice as a value.
- `map[string][]string` can represent one user with many wallets.
- A user can be created automatically if the key does not exist.
- A wallet can be added to a user's wallet list.
- Duplicate wallets can be prevented.
- A wallet can be deleted from a slice inside a map.
- `len` can count users in a map.

## Run

```bash
go run main.go