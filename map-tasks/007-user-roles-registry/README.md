# Task 007 - User Roles Registry With Map And Slice

## Goal

Create a user roles registry that stores multiple roles for each user.

## Concepts

* Interfaces
* Structs
* Methods
* Maps
* Slices
* map[string][]string
* append
* range
* len
* Delete item from slice
* Duplicate Check

## What I Learned

* A map can store a slice as a value.
* One user can have multiple roles.
* A user can be created automatically if the key does not exist.
* A role can be added to a user.
* Duplicate roles can be prevented.
* A role can be deleted from a slice inside a map.
* `range` can be used to iterate through users and roles.
* `len` returns the number of users.

## Features

* Add role
* Delete role
* Get roles of user
* View all users
* Count users
* Prevent duplicate roles

## Run

```bash
go run main.go
```
