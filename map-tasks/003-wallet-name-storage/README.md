# Task 003 - Wallet Owner Storage With Map

## Goal

Create a wallet owner storage that stores wallet addresses and owner names using a map.

## Concepts

- Interfaces
- Structs
- Methods
- Maps
- map[string]string
- Add value by key
- Delete value by key
- Get value by key
- Duplicate Check
- range
- len

## What I Learned

- A map can store wallet address as a key and owner name as a value.
- A value can be added by key.
- Existing keys can be checked before adding new data.
- `delete` removes a wallet by address.
- `range` can be used to print all key-value pairs.
- `len` returns the number of wallets in the map.

## Run

```bash
go run main.go