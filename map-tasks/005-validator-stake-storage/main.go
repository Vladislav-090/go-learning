package main

import "fmt"

type ValidatorStakeService interface {
	AddValidator(name string, stake int)
	DeleteValidator(name string)
	GetStakeOfValidator(name string) int
	UpdateStakeOfValidator(name string, newStake int)
	ViewAllValidators()
	ValidatorsCount() int
}

type ValidatorStorage struct {
	Validators map[string]int
}

func (v *ValidatorStorage) AddValidator(name string, stake int) {
	if stake < 0 {
		fmt.Println("Validator cannot be added with negative stake!")
		return
	}

	if _, exists := v.Validators[name]; exists {
		fmt.Println("Validator is already exist!", name)
		return
	}

	v.Validators[name] = stake
	fmt.Println("Validator added:", name, stake)
}

func (v *ValidatorStorage) DeleteValidator(name string) {
	if _, exists := v.Validators[name]; exists {
		delete(v.Validators, name)
		fmt.Println("Validator has been deleted!", name)
		return
	}
	fmt.Println("Validator not found!")
}

func (v *ValidatorStorage) GetStakeOfValidator(name string) int {
	return v.Validators[name]
}

func (v *ValidatorStorage) UpdateStakeOfValidator(name string, newStake int) {
	if newStake < 0 {
		fmt.Println("Stake cannot be negative!")
		return
	}

	if _, exists := v.Validators[name]; exists {
		v.Validators[name] = newStake
		fmt.Printf("Stake of Validator %s has been updated to  %d\n", name, newStake)
		return
	}
	fmt.Println("Validator not found!", name)
}

func (v *ValidatorStorage) ViewAllValidators() {
	for name, stake := range v.Validators {
		fmt.Println("Validator info :", name, stake)
	}
}

func (v *ValidatorStorage) ValidatorsCount() int {
	return len(v.Validators)
}

func PrintInfo(v ValidatorStakeService, name string, stake int, newStake int) {
	v.ViewAllValidators()
	fmt.Println("Validators count:", v.ValidatorsCount())

	v.AddValidator(name, stake)
	fmt.Println("Stake:", v.GetStakeOfValidator(name))

	v.ViewAllValidators()
	fmt.Println("Validators count:", v.ValidatorsCount())

	v.UpdateStakeOfValidator(name, newStake)
	fmt.Println("Updated stake:", v.GetStakeOfValidator(name))

	v.ViewAllValidators()

	v.DeleteValidator(name)

	v.ViewAllValidators()
	fmt.Println("Validators count:", v.ValidatorsCount())
}

func main() {
	validatorStorage := ValidatorStorage{
		Validators: map[string]int{
			"Vladislav": 1000,
			"Viola":     4000,
			"Afina":     5000,
		},
	}
	PrintInfo(&validatorStorage, "Bruce", 4000, -100)
}
