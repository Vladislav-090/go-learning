package main

import "fmt"

type ValidatorRegistryService interface {
	AddValidator(name string)
	DeleteValidator(name string)
	ViewAllValidators()
	QuantityValidators() int
}

type ValidatorRegistry struct {
	Validators []string
}

func (v *ValidatorRegistry) AddValidator(name string) {
	v.Validators = append(v.Validators, name)
	fmt.Println("New validator added!", name)
}

func (v *ValidatorRegistry) DeleteValidator(name string) {
	for i, validator := range v.Validators{
		if validator == name {
			v.Validators = append(v.Validators[:i], v.Validators[i+1:]...)
			fmt.Println("The validator has been deleted!", name)
			return
		}
	}
	fmt.Println("Validator not found!", name)
}

func (v *ValidatorRegistry) ViewAllValidators() {
	for _, validator := range v.Validators {
		fmt.Println("Validator name is:", validator)
	}
}

func (v *ValidatorRegistry) QuantityValidators() int {
	return	len(v.Validators)
}

func PrintInfo(v ValidatorRegistryService, name string) {
	fmt.Println("Now quantity of Validators is:", v.QuantityValidators())
	v.AddValidator(name)
	v.ViewAllValidators()
	fmt.Println("Now quantity of Validators is:", v.QuantityValidators())
	v.DeleteValidator(name)
	v.ViewAllValidators()
} 

func main() {
	validatorregisrty := ValidatorRegistry{
		Validators: []string {"Vladislav", "Viola", "Alexandr", "Daniel"},
	}
	PrintInfo(&validatorregisrty,"Samuel")
}