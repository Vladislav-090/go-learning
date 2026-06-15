# Task 009 - Product Tags Registry With Map And Slice

## Goal

Create a product tags registry that stores multiple tags for each product.

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
- One product can have multiple tags.
- A product can be created automatically if the key does not exist.
- A tag can be added to a product.
- Duplicate tags can be prevented.
- A tag can be deleted from a slice inside a map.
- `range` can be used to iterate through products and tags.
- `len` returns the number of products.

## Run

```bash
go run main.go