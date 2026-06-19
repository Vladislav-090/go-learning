# Task 014 - JSON With Omitempty

## Goal

Learn how to hide empty fields and sensitive fields from JSON output.

## Concepts

- Structs
- JSON
- JSON tags
- omitempty
- json:"-"
- encoding/json
- json.MarshalIndent
- []byte
- error handling

## What I Learned

- `omitempty` removes a field from JSON if it has an empty value.
- An empty string `""` is considered empty.
- `json:"-"` completely excludes a field from JSON.
- Sensitive fields like passwords should not be exposed in JSON.
- `json.MarshalIndent()` can be used to print formatted JSON.
- JSON tags can control both field names and visibility.

## Features

- Create user with email and password
- Hide password from JSON
- Hide empty email from JSON
- Print formatted JSON
- Handle marshal errors

## Run

go run main.go