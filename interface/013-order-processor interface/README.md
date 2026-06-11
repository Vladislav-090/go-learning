# Task 013 - Order Processor Interface

## Goal

Create an OrderProcessor interface and use it to process orders through different services.

## Concepts

- Interfaces
- Multiple Methods
- Pointer Receiver
- State Management
- Return Values
- Common Behavior

## What I Learned

- Different services can implement the same interface.
- One function can work with online shop, taxi, and food delivery services.
- Methods can change internal state such as orders count.
- Methods can return service information and current state.
- Interfaces help write reusable functions for different implementations.

## Run

```bash
go run main.go