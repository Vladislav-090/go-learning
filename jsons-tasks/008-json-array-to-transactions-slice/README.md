# Task 008 - JSON Array To Transactions Slice

## Goal

Convert a JSON array into a slice of transaction structs.

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
- `json.Unmarshal()` fills the target variable with decoded data.
- `json.Unmarshal()` requires a pointer to the target slice.
- `&transactions` gives `Unmarshal` access to modify the slice.
- `[]byte(jsonString)` converts a JSON string into bytes.
- `range` can be used to iterate over decoded transactions.
- Error should be checked before using decoded data.

## Features

- Create JSON array string
- Convert JSON array to `[]Transaction`
- Print each transaction
- Handle unmarshal error

## Run

go run main.go