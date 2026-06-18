# Task 006 - JSON Array To Products Slice

## Goal

Convert a JSON array into a slice of product structs.

## Concepts

- Structs
- Methods
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
- `json.Unmarshal()` can fill `[]Product`.
- `json.Unmarshal()` requires a pointer to the target slice.
- `[]byte(jsonString)` converts a JSON string into bytes.
- `range` can be used to iterate over decoded products.
- Error should be checked before using decoded data.

## Features

- Create JSON array string
- Convert JSON array to `[]Product`
- Print each product
- Handle unmarshal error

## Run

go run main.go