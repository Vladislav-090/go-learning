# Task 008 - Car Rental With Error

## Goal

Create a car rental system with error handling.

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
- Fuel validation

## What I Learned

- A boolean field can represent rental state.
- A car cannot be rented twice.
- A car cannot be rented if fuel is too low.
- Fuel amount should be validated before updating.
- A returned car can update its current fuel level.
- Invalid business operations should return errors.
- `nil` means the operation completed successfully.
- `err != nil` means an error happened.

## Features

- Rent a car
- Return a car
- Add fuel
- Prevent renting a car twice
- Prevent renting a car with low fuel
- Prevent invalid fuel values
- Print car information

## Run

```bash
go run main.go