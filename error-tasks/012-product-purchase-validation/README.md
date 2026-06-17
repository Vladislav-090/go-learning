# Task 012 - Product Purchase Validation With Error

## Goal

Create a product purchase system with error handling.

## Concepts

- Structs
- Methods
- Pointer Receiver
- errors.New
- error return value
- nil
- if err != nil
- Stock validation
- Quantity validation

## What I Learned

- Product stock should be validated before purchase.
- Quantity should be positive for both buying and restocking.
- A product cannot be bought if stock is empty.
- A product cannot be bought if requested quantity is greater than stock.
- Failed operations should not change product stock.
- `nil` means the operation completed successfully.
- `err != nil` means the operation failed.

## Features

- Buy product
- Restock product
- Prevent buying zero or negative quantity
- Prevent buying more than available stock
- Prevent invalid restock quantity
- Print product information

## Run

```bash
go run main.go