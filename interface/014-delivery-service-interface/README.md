# Task 014 - Delivery Service Interface

## Goal

Create a DeliveryService interface and use it to manage orders through different delivery services.

## Concepts

- Interfaces
- Multiple Methods
- Pointer Receiver
- State Management
- Return Values
- Common Behavior

## What I Learned

- Different delivery services can implement the same interface.
- One function can work with Glovo, Wolt, and Yandex through a common interface.
- Methods can change internal state such as delivery count.
- Early return prevents invalid state changes.
- Interfaces help reuse logic across different implementations.

## Run

```bash
go run main.go