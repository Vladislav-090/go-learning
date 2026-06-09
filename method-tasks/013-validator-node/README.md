# Task 013 - Validator Node

## Goal

Create a Validator struct with methods for starting, stopping, signing blocks, and displaying validator information.

## Concepts

- Structs
- Methods
- Pointer Receiver
- State Management
- Early Return
- Blockchain Basics

## What I Learned

- Methods can represent actions of a blockchain validator.
- A validator can sign blocks only when it is online.
- Early return helps prevent invalid operations.
- Pointer receivers are used when methods modify struct fields.

## Run

```bash
go run main.go