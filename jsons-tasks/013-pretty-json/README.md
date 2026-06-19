# Task 013 - Pretty JSON

## Goal

Convert a struct with nested wallet data into formatted JSON.

## Concepts

- Structs
- Slices
- Nested Structs
- JSON
- JSON tags
- encoding/json
- json.MarshalIndent
- []byte
- error handling

## What I Learned

- `json.MarshalIndent()` creates formatted JSON with indentation.
- Pretty JSON is easier to read than compact JSON.
- Nested structs and slices are formatted automatically.
- `json.MarshalIndent()` returns `[]byte` and `error`.
- Formatted JSON is useful for debugging and configuration files.

## Features

- Create user with multiple wallets
- Convert struct to formatted JSON
- Print formatted JSON
- Handle marshal errors

## Run

go run main.go