# Task 011 - User With Multiple Wallets To JSON

## Goal

Create a user with multiple wallets and convert it to JSON.

## Concepts

- Structs
- Slices
- Nested Structs
- JSON
- JSON tags
- encoding/json
- json.Marshal
- []byte
- error handling

## What I Learned

- A struct can contain a slice of another struct.
- `Wallets []Wallet` becomes a JSON array of wallet objects.
- Nested slices are converted automatically by `json.Marshal()`.
- JSON tags control field names in nested structures.
- `json.Marshal()` returns `[]byte` and `error`.

## Features

- Create wallet struct
- Create user struct with multiple wallets
- Convert nested slice structure to JSON
- Print JSON
- Handle marshal error

## Run

go run main.go