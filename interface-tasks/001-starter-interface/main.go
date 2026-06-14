package main

import "fmt"

type Starter interface {
	Start()
}

type TelegramBot struct {
	Name string
}
type RPCNode struct {
	Name string
}

func (t *TelegramBot) Start() {
	fmt.Println("Bot has been started!")
}

func (r *RPCNode) Start() {
	fmt.Println("Node has been started!")
}

func RunService(s Starter) {
	s.Start()
}

func main() {
	bot := TelegramBot{Name: "AlertBot"}
	node := RPCNode{Name: "ArbitrumNode"}

	RunService(&bot)
	RunService(&node)
}