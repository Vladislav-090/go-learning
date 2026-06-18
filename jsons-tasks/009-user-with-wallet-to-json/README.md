# Task 009 - User With Wallet To JSON

## Goal

Create a user with a nested wallet structure and convert it to JSON.

## Concepts

- Structs
- Nested Structs
- JSON
- JSON tags
- encoding/json
- json.Marshal
- []byte
- error handling
- string conversion

## What I Learned

- A struct can contain another struct.
- Nested structs become nested JSON objects.
- `json.Marshal()` recursively converts nested structs.
- JSON tags control field names in nested objects.
- `json.Marshal()` returns `[]byte` and `error`.
- Errors should be checked before using JSON data.

## Features

- Create wallet struct
- Create user struct with wallet
- Convert nested struct to JSON
- Print JSON
- Handle marshal error

## Run

go run main.go