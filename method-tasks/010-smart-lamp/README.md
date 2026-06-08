# Task 010 - Smart Lamp

## Goal

Create a Lamp struct with methods for turning the lamp on, turning it off, increasing brightness, and displaying lamp information.

## Concepts

- Structs
- Methods
- Pointer Receiver
- Boolean Fields
- State Management
- Value Limitation
- Early Return

## What I Learned

- Methods can change the state of a struct.
- Boolean fields can control whether an action is allowed.
- Pointer receivers are used when a method modifies the original struct.
- A method can stop execution early using `return`.
- Values can be limited, for example brightness should not be higher than 100.

## Run

```bash
go run main.go