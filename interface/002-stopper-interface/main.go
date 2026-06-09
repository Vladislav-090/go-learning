package main 

import "fmt"

type Stopper interface {
	Stop()
}

type TelegramBot struct {
	Name string
}

type RPCNode struct {
	Name string
}

func (t *TelegramBot) Stop() {
	fmt.Println("Bot has been stopped")
}

func (r *RPCNode) Stop() {
	fmt.Println("Node has been stopped")
}

func StopService(s Stopper) {
	s.Stop()
}

func main() {
	b := TelegramBot{Name: "AlertBot"}
	n := RPCNode{Name: "Arbitrum"}

	StopService(&b)
	StopService(&n)
}