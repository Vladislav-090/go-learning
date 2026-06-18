# Task 007 - Transactions Slice To JSON

## Goal

Create a slice of transactions and convert it to a JSON array.

## Concepts

- Structs
- Slices
- JSON
- JSON tags
- encoding/json
- json.Marshal
- []byte
- error handling
- string conversion

## What I Learned

- A slice of transactions can be converted into a JSON array.
- `json.Marshal()` works with slices of structs.
- JSON field names are controlled by struct tags.
- `json.Marshal()` returns `[]byte` and `error`.
- JSON data can be printed using `string(data)`.
- Errors should be checked before using the result.

## Features

- Create transaction slice
- Convert transactions to JSON
- Print JSON array
- Handle marshal errors

## Run

go run main.go