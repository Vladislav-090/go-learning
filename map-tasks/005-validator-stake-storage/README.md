# Task 005 - Validator Stake Storage With Map

## Goal

Create a validator stake storage that stores validator stakes using a map.

## Concepts

- Interfaces
- Structs
- Methods
- Maps
- map[string]int
- Add value by key
- Delete value by key
- Get value by key
- Update value by key
- Duplicate Check
- Validation
- range
- len

## What I Learned

- A map can store validator name as a key and stake as a value.
- Existing validators can be checked before adding new data.
- Stake values can be validated before adding or updating.
- A value can be updated directly by key.
- `delete` removes a validator by key.
- `range` can be used to print all validators and stakes.
- `len` returns the number of validators in the map.

## Run

```bash
go run main.go