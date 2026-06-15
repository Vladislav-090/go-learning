# Task 010 - Validator Delegators Registry With Map And Slice

## Goal

Create a validator delegators registry that stores multiple delegators for each validator.

## Concepts

- Interfaces
- Structs
- Methods
- Maps
- Slices
- map[string][]string
- append
- range
- len
- Delete item from slice
- Duplicate Check

## What I Learned

- A map can store a slice as a value.
- One validator can have multiple delegators.
- A validator can be created automatically if the key does not exist.
- A delegator can be added to a validator.
- Duplicate delegators can be prevented.
- A delegator can be deleted from a slice inside a map.
- `range` can be used to iterate through validators and delegators.
- `len` returns the number of validators.

## Run

```bash
go run main.go