# Task 005 - Products Slice To JSON

## Goal

Create a slice of products and convert it to a JSON array.

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

- A slice of structs can be converted into a JSON array.
- `json.Marshal()` works with slices.
- JSON field names are controlled by struct tags.
- `json.Marshal()` returns `[]byte` and `error`.
- Errors should be checked before using the result.
- JSON data can be printed by converting `[]byte` to string.

## Features

- Create product slice
- Convert slice to JSON
- Print JSON array
- Handle marshal errors

## Run

go run main.go