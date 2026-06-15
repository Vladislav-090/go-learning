# Task 021 - Library Registry With Slice

## Goal

Create a library registry that stores books using a slice.

## Concepts

- Interfaces
- Structs
- Methods
- Slices
- append
- range
- len
- Update item
- Delete item
- Duplicate Check

## What I Learned

- A slice can be used to store a collection of books.
- New books can be added using `append`.
- Duplicate books can be prevented.
- Books can be updated by index.
- Books can be deleted using:

```go
append(slice[:i], slice[i+1:]...)