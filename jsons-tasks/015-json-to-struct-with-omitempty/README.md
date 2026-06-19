# Task 015 - JSON To Struct With Omitempty

## Goal

Learn how missing JSON fields are handled during Unmarshal.

## Concepts

- Structs
- JSON
- JSON tags
- omitempty
- json:"-"
- encoding/json
- json.Unmarshal
- []byte
- Zero values
- error handling

## What I Learned

- Missing JSON fields do not cause errors.
- Unmarshal fills only fields that exist in JSON.
- Missing fields receive Go zero values.
- String zero value is "".
- Integer zero value is 0.
- `json:"-"` ignores the field during JSON processing.
- Unmarshal can work with partial JSON data.

## Features

- Decode multiple JSON objects
- Handle missing fields
- Print decoded structs
- Demonstrate zero values
- Handle unmarshal errors

## Run

go run main.go