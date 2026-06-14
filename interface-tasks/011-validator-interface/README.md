# Task 011 - Validator Interface

## Goal

Create a Validator interface and manage different blockchain validators through a common interface.

## Concepts

- Interfaces
- Multiple Methods
- State Management
- Return Values
- Polymorphism
- Blockchain Validators

## Interface

```go
type ValidatorInterface interface {
	GetName() string
	Start()
	Stop()
	GetStatus() string
	SignBlock()
	GetSignedBlocks() int
}