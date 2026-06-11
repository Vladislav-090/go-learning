# Task 012 - User Manager Interface

## Goal

Create a UserManager interface and use it to manage users through different services.

## Concepts

- Interfaces
- Multiple Methods
- Pointer Receiver
- State Management
- Return Values
- Common Behavior

## What I Learned

- An interface can describe several actions.
- Different services can implement the same interface.
- Methods like `CreateUser` and `DeleteUser` change the struct state.
- Methods like `GetUserCount` and `GetServiceName` return data.
- One function can work with different services through the same interface.

## Run

```bash
go run main.go