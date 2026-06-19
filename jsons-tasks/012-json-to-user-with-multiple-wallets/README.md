# Task 012 - JSON To User With Multiple Wallets

## Goal

Convert nested JSON with a wallets array into a Go struct.

## Concepts

- Structs
- Slices
- Nested Structs
- JSON
- JSON tags
- encoding/json
- json.Unmarshal
- []byte
- Pointer to struct
- range loop

## What I Learned

- JSON object can contain an array of nested objects.
- A JSON array inside an object can be decoded into a slice field.
- `Wallets []Wallet` matches JSON field `wallets`.
- `json.Unmarshal()` fills nested slices automatically.
- `range` can be used to iterate over decoded wallets.
- Error should be checked before using decoded data.

## Run

go run main.go