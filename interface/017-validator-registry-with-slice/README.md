# Task 017 - Validator Registry With Slice

## Goal

Create a validator registry that stores validator names in a slice and works through an interface.

## Concepts

- Interfaces
- Structs
- Methods
- Slices
- append
- range
- len
- Deleting from slice

## What I Learned

- A slice can store a list of validators.
- `append` adds a new validator to the slice.
- `range` can be used to loop through validators.
- `len` returns the current number of validators.
- A validator can be removed from a slice by combining elements before and after it.
- An interface can describe behavior for managing a registry.

## Run

```bash
go run main.go