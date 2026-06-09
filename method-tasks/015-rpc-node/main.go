package main

import "fmt"

type RPCNode struct {
	Name      string
	IsRunning bool
	Request   int
}

func (r *RPCNode) Start() {
	if r.IsRunning {
		fmt.Println("Node is already running!")
		return
	}
	r.IsRunning = true
	fmt.Println("Node has been started")
}

func (r *RPCNode) Stop() {
	if !r.IsRunning {
		fmt.Printf("Node is already stopped")
		return
	}
	r.IsRunning = false 
	fmt.Println("Node has been stopped")
}

func (r *RPCNode) HandleRequest() {
	if !r.IsRunning {
		fmt.Println("You need to started Node first!")
		return
	}
	r.Request ++
	fmt.Println("request was complited")
}

func (r *RPCNode) PrintInfo() {

fmt.Println("RpcNode Name is:", r.Name)
fmt.Println("RpcNode Status:", r.IsRunning)
fmt.Println("Quantity of requests", r.Request)
}

func main() {
	rpcnode := RPCNode{
		Name: "ArbitrumNode",
		IsRunning: true,
		Request: 0,
	}

	rpcnode.PrintInfo()

	rpcnode.HandleRequest()
	rpcnode.Start()
	rpcnode.Stop()
	rpcnode.HandleRequest()
	rpcnode.Start()
	rpcnode.HandleRequest()
	rpcnode.HandleRequest()
	rpcnode.PrintInfo()
}