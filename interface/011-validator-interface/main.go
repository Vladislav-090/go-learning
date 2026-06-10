package main

import "fmt"

type ValidatorInterface interface {
	GetName() string
	Start()
	Stop()
	GetStatus() string
	SignBlock()
	GetSignedBlocks() int
}

type EthereumValidator struct {
	Name         string
	IsOnline     bool
	SignedBlocks int
}
type CosmosValidator struct {
	Name         string
	IsOnline     bool
	SignedBlocks int
}

func (e *EthereumValidator) GetName()string {
	return e.Name
}

func (e *EthereumValidator) Start() {
	if e.IsOnline {
		fmt.Println("Validator is already started")
		return
	}
	e.IsOnline = true
	fmt.Println("Validator has been started")
}

func (e *EthereumValidator) Stop() {
	if !e.IsOnline {
		fmt.Println("Validator is already stopped")
		return
	}
	e.IsOnline = false
	fmt.Println("Validator has been stopped")
}

func (e *EthereumValidator) GetStatus() string {
	if e.IsOnline { 
	return "Validator is online"
	} 
		return "Validator is offline"
} 

func (e *EthereumValidator) SignBlock() {
	if !e.IsOnline {
		fmt.Println("Validator must be online first")
		return
	} 
	e.SignedBlocks ++
		fmt.Println("block has been signed")
}

func (e *EthereumValidator) GetSignedBlocks() int {
	return  e.SignedBlocks
} 


////////////////////
func (c *CosmosValidator) GetName()string {
	return c.Name
}

func (c *CosmosValidator) Start() {
	if c.IsOnline {
		fmt.Println("Validator is already started")
		return
	}
	c.IsOnline = true
	fmt.Println("Validator has been started")
}

func (c *CosmosValidator) Stop() {
	if !c.IsOnline {
		fmt.Println("Validator is already stopped")
		return
	}
	c.IsOnline = false
	fmt.Println("Validator has been stopped")
}

func (c *CosmosValidator) GetStatus() string {
	if c.IsOnline { 
	return "Validator is online"
	} 
		return "Validator is offline"
} 

func (c *CosmosValidator) SignBlock() {
	if !c.IsOnline {
		fmt.Println("Validator must be online first")
		return
	} 
	c.SignedBlocks ++
		fmt.Println("block has been signed")
}

func (c *CosmosValidator) GetSignedBlocks() int {
	return  c.SignedBlocks
} 


func ManageValidator(v ValidatorInterface) {
	fmt.Println(v.GetName())
	fmt.Println(v.GetStatus())
	v.Start()
	v.SignBlock()
	fmt.Println("Now signed blocks:", v.GetSignedBlocks())
	v.Stop()
}


func main() {
	ethereum := EthereumValidator{
		Name: "Ethereum Validator",
		IsOnline: true,
		SignedBlocks: 3,
	}

	cosmos := CosmosValidator{
		Name: "Cosmos Validator",
		IsOnline: false,
		SignedBlocks: 1,
	}
	ManageValidator(&ethereum)
	ManageValidator(&cosmos)
}