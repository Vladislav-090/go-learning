# Task 003 - Users Slice To JSON

## Goal

Create a slice of users and convert it to a JSON array.

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

- A slice of structs can be converted to JSON.
- `json.Marshal()` can work with both a single struct and a slice of structs.
- One struct becomes a JSON object.
- A slice of structs becomes a JSON array.
- `json.Marshal()` returns `[]byte` and `error`.
- `[]byte` can be converted to string with `string(data)`.
- Error should be checked before printing JSON data.

## Features

- Create a slice of users
- Convert users slice to JSON
- Print JSON array
- Handle marshal error

## Run

go run main.go