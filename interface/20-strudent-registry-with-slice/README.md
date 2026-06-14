# Task 020 - Student Registry With Slice

## Goal

Create a student registry that stores students in a slice and works through an interface.

## Concepts

- Interfaces
- Structs
- Methods
- Slices
- append
- range
- len
- Duplicate Check
- Update Item In Slice
- Delete Item From Slice

## What I Learned

- A slice can store a list of students.
- `append` adds a new student to the slice.
- `range` can be used to search through students.
- Duplicate checks prevent adding the same student twice.
- Duplicate checks can also prevent updating a student to an existing name.
- A student can be updated by changing the value at a specific index.
- A student can be deleted from a slice by combining the elements before and after it.
- An interface can describe registry behavior.

## Run

```bash
go run main.go