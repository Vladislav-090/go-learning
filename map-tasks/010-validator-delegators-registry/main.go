package main

import (
	"fmt"
)

type ValidatorDelegatorsRegistry interface {
	AddDelegator(validator string, delegator string)
	DeleteDelegator(validator string, delegator string)
	ViewAllValidators()
	GetDelegators(validator string) []string
	ValidatorsCount() int
}

type ValidatorDelegatorsStruct struct {
	Validators map[string][]string
}

func (v *ValidatorDelegatorsStruct) AddDelegator(validator string, delegator string) {
	if _, exist := v.Validators[validator]; !exist {
		v.Validators[validator] = []string{}
		fmt.Println("New Validator with empty delegators slice created", validator)
	}
	for _, currentDelegator := range v.Validators[validator] {
		if currentDelegator == delegator {
			fmt.Println("Delegator already exist!", delegator)
			return
		}
	}
	v.Validators[validator] = append(v.Validators[validator], delegator)
	fmt.Println("Delegator added!", validator, delegator)
}

func (v *ValidatorDelegatorsStruct) DeleteDelegator(validator string, delegator string) {
	delegators, exist := v.Validators[validator]
	if !exist {
		fmt.Println("Validator not found!", validator)
		return
	}

	for i, currentDelegator := range delegators {
		if currentDelegator == delegator {
			delegators = append(delegators[:i], delegators[i+1:]...)
			v.Validators[validator] = delegators
			fmt.Println("Delegator deleted:", delegator)
			return
		}
	}
	fmt.Println("Delegator not found", delegator)
}

func (v *ValidatorDelegatorsStruct) ViewAllValidators() {
	for validator, delegators := range v.Validators {
		fmt.Printf("Validator %s has %v delegators \n", validator, delegators)
	}
}

func (v *ValidatorDelegatorsStruct) GetDelegators(validator string) []string {
	return v.Validators[validator]
}

func (v *ValidatorDelegatorsStruct) ValidatorsCount() int {
	return len(v.Validators)
}
func PrintInfo(v ValidatorDelegatorsRegistry, validator string, delegator string) {
	v.ViewAllValidators()
	fmt.Println("Validators count:", v.ValidatorsCount())

	v.AddDelegator(validator, delegator)
	fmt.Println("Delegators of validator:", validator, v.GetDelegators(validator))

	v.ViewAllValidators()

	v.DeleteDelegator(validator, delegator)
	fmt.Println("Delegators of validator:", validator, v.GetDelegators(validator))

	v.ViewAllValidators()
	fmt.Println("Validators count:", v.ValidatorsCount())
}

func main() {
	registry := ValidatorDelegatorsStruct{
		Validators: map[string][]string{
			"ValidatorOne": {"Vladislav", "Alex"},
			"ValidatorTwo": {"Viola", "Max"},
		},
	}

	PrintInfo(&registry, "ValidatorOne", "Bruce")
}
