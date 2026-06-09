# Task 015 - RPC Node

## Goal

Create an RPCNode struct with methods for starting, stopping, handling requests, and displaying node information.

## Concepts

- Structs
- Methods
- Pointer Receiver
- Boolean Fields
- State Management
- Early Return
- Blockchain Infrastructure Basics

## What I Learned

- An RPC node can process requests only when it is running.
- Methods can control the state of an object.
- Pointer receivers are used when methods modify struct fields.
- Early return helps prevent invalid operations.
- Request counters can be updated through methods.

## Run

```bash
go run main.go