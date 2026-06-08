# Task 009 - Phone Battery

## Goal

Create a Phone struct with methods for powering on, using, charging, and printing phone information.

## Concepts

- Structs
- Methods
- Pointer Receiver
- Boolean Fields
- State Management
- Value Limitation
- Early Return

## What I Learned

- A struct can store both data and state.
- Methods can change several fields of the same struct.
- A boolean field can control whether an action is allowed.
- `return` can stop method execution when an action is not allowed.
- Values can be limited, for example battery level should not be higher than 100.

## Run

```bash
go run main.go