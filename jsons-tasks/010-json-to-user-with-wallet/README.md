# Task 010 - JSON To User With Wallet

## Goal

Convert nested JSON into a Go struct with a nested wallet struct.

## Concepts

- Structs
- Nested Structs
- JSON
- JSON tags
- encoding/json
- json.Unmarshal
- []byte
- Pointer to struct
- error handling

## What I Learned

- Nested JSON objects can be converted into nested Go structs.
- `json.Unmarshal()` fills nested struct fields automatically.
- JSON object should be decoded into a struct.
- JSON array should be decoded into a slice.
- `&user` is required because Unmarshal needs to modify the struct.
- JSON tags connect JSON field names with Go struct fields.

## Features

- Create nested JSON string
- Convert JSON to User struct
- Fill nested Wallet struct
- Print user information
- Handle unmarshal error

## Run

go run main.go