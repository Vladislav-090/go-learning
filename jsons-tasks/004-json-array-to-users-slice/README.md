# Task 004 - JSON Array To Users Slice

## Goal

Convert a JSON array into a slice of Go structs.

## Concepts

- Structs
- Slices
- JSON
- JSON tags
- encoding/json
- json.Unmarshal
- []byte
- Pointer to slice
- error handling
- range loop

## What I Learned

- A JSON array can be converted into a slice of structs.
- `json.Unmarshal()` can fill a slice of structs.
- `json.Unmarshal()` accepts `[]byte` and a pointer to the target variable.
- `&users` is required because Unmarshal needs to modify the slice.
- `range` can be used to iterate over decoded users.
- Error should be checked before using decoded data.

## Features

- Create JSON array string
- Convert JSON array to `[]User`
- Print each user
- Handle unmarshal error

## Run

go run main.go