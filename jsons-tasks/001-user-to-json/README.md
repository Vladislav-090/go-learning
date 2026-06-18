# Task 001 - User To JSON

## Goal

Create a user struct and convert it to JSON.

## Concepts

- Structs
- Methods
- Pointer Receiver
- JSON
- JSON tags
- encoding/json
- json.Marshal
- []byte
- error handling
- string conversion

## What I Learned

- `json.Marshal()` converts a Go struct into JSON.
- `json.Marshal()` returns `[]byte` and `error`.
- `[]byte` can be converted to string with `string(data)`.
- JSON tags control field names in the final JSON output.
- Struct fields must start with a capital letter to be exported.
- Error should be checked after calling `json.Marshal()`.

## Features

- Create user struct
- Convert struct to JSON
- Print JSON as string
- Handle marshal error

## Run

go run main.go