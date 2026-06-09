package main

import "fmt"

type Service interface {
	Start()
	Stop()
}

type Bot struct {
	Name string
}

type Node struct {
	Name string
}

func (b *Bot) Start(){
	fmt.Println("Bot has been started")
}

func (b *Bot) Stop(){
	fmt.Println("Bot has been stopped")
}

func (n *Node) Start(){
	fmt.Println("Node has been started")
}

func (n *Node) Stop(){
	fmt.Println("Node has been stopped")
}

func ManageService(s Service) {
	s.Start()
	s.Stop()
}

func main() {
	bot := Bot{Name: "Alert"}
	node := Node{Name: "Arbitrum"}

	ManageService(&bot)
	ManageService(&node)
}