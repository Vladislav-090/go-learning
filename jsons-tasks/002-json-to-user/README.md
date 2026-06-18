# Task 002 - JSON To User

## Goal

Convert JSON data into a Go struct.

## Concepts

- Structs
- Methods
- JSON
- JSON tags
- encoding/json
- json.Unmarshal
- []byte
- error handling

## What I Learned

- `json.Unmarshal()` converts JSON into a Go struct.
- `json.Unmarshal()` accepts `[]byte` and a pointer to a struct.
- A pointer is required because `Unmarshal` fills the struct fields.
- JSON field names are matched using struct tags.
- Error should be checked after calling `json.Unmarshal()`.

## Features

- Create JSON string
- Convert JSON to struct
- Print struct fields
- Handle unmarshal errors

## Run

go run main.go