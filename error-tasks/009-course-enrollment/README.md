# Task 009 - Course Enrollment With Error

## Goal

Create a course enrollment system with error handling.

## Concepts

- Structs
- Methods
- Pointer Receiver
- Slices
- append
- range
- len
- errors.New
- error return value
- nil
- if err != nil
- Duplicate Check
- Limit Validation

## What I Learned

- A slice can store course students.
- Student names should be validated before adding or removing.
- Duplicate students can be prevented.
- Course capacity can be limited with `MaxStudents`.
- `len()` can be used to check current number of students.
- Invalid enrollment operations should return errors.
- Removing an item from a slice uses `append(slice[:i], slice[i+1:]...)`.
- Printing the returned error helps understand why an operation failed.

## Features

- Add student
- Remove student
- Prevent empty student name
- Prevent duplicate student enrollment
- Prevent adding students when course is full
- Count students
- Print course information

## Run

```bash
go run main.go