# Task 007 - Hotel Room Booking With Error

## Goal

Create a hotel room booking system with error handling.

## Concepts

- Structs
- Methods
- Pointer Receiver
- Boolean validation
- errors.New
- error return value
- nil
- if err != nil
- State validation

## What I Learned

- A boolean field can represent booking state.
- A room cannot be booked twice.
- A booking cannot be canceled if the room is not booked.
- Invalid business operations should return errors.
- `nil` means the operation completed successfully.
- `err != nil` means an error happened.
- Error handling helps protect object state from invalid changes.

## Features

- Book a room
- Cancel booking
- Prevent duplicate booking
- Prevent canceling an unbooked room
- Print room information

## Run

```bash
go run main.go