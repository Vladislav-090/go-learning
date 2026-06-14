package main

import "fmt"

type DeliveryService interface {
	GetServiceName() string
	CreateOrder(name string)
	DeleteOrder(name string)
	OrdersCount() int
	PrintOrders()
}

type Glovo struct {
	Name   string
	Orders []string
}

func (g *Glovo) CreateOrder(name string) {
	g.Orders = append(g.Orders, name)
	fmt.Println("Got a new order!", name)
}

func (g *Glovo) DeleteOrder(name string) {
	for index, order := range g.Orders{
		if order == name {
		g.Orders = append(g.Orders[:index], g.Orders[index+1:]... )
		fmt.Println("Order has been deleted!",name)
		return
		}
	}
	fmt.Println("Order not found!", name)
}

func (g *Glovo) OrdersCount()int {
	return len(g.Orders)
}

func (g *Glovo) GetServiceName()string {
	return g.Name
}

func (g *Glovo) PrintOrders() {
	for _, order := range g.Orders {
		fmt.Println("Current Order:", order)
	}
}


func PrintInfo(d DeliveryService, name string) {
	fmt.Println("Service name is :", d.GetServiceName())

	d.PrintOrders()
	d.CreateOrder(name)
	d.PrintOrders()
	fmt.Println("Orders count", d.OrdersCount())

	d.DeleteOrder(name)
	d.PrintOrders()
	fmt.Println("Orders count", d.OrdersCount())
}

func main() {
	glovo := Glovo{
		Name: "Glovo",
		Orders: []string{"KFC", "McDonalds", "SuperMarket"},
	}

	PrintInfo(&glovo, "NiceDog")
}
