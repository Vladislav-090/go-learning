package main

import "fmt"

type Validator struct {
	Name         string
	IsOnline     bool
	SignedBlocks int
}

func (v *Validator) SignBlock(){
	if !v.IsOnline {
		fmt.Println("Validator is offline, cannot sign block")
		return
	}

	v.SignedBlocks ++
	fmt.Println("Block signed")
}

func (v *Validator) Start() {
	if v.IsOnline {
		fmt.Println("Validator is already Online!")
		return
	}
	
	v.IsOnline = true
	fmt.Println("Validator started")
}

func (v *Validator) Stop() {
	if !v.IsOnline {
		fmt.Println("Validator is already offline")
		return
	}

	v.IsOnline = false
	fmt.Println("Validator Stopped")
}

func (v *Validator) PrintInfo(){
	fmt.Println("Name of validator", v.Name)
	fmt.Println("Status:", v.IsOnline)
	fmt.Println("quantity blocks signed:", v.SignedBlocks)

}

func main() {
	validator := Validator{
		Name: "Vladislav",
		IsOnline: true,
		SignedBlocks: 40,
	}

validator.PrintInfo()

validator.SignBlock()
validator.SignBlock()

validator.Stop()
validator.SignBlock()

validator.Start()
validator.SignBlock()

validator.PrintInfo()

}