package main

import "fmt"

type DeliveryService interface {
	GetServiceName() string
	CreateOrder(name string)
	DeleteOrder(name string)
	GetDeliveryCount()int
}

type Yandex struct {
	Name string
	DeliveryCount int
}

type Glovo struct {
	Name string
	DeliveryCount int
}

type Wolt struct {
	Name string
	DeliveryCount int
}

func (g *Glovo) GetServiceName() string {
	return g.Name
}

func (g *Glovo) CreateOrder(name string) {
	g.DeliveryCount++
	fmt.Println("New order created!", name)
}

func (g *Glovo) DeleteOrder(name string) {
	if g.DeliveryCount == 0 {
		fmt.Println("Dont have any orders")
		return
	}
	g.DeliveryCount --
	fmt.Println("Order has been deleted!", name)
}

func (g *Glovo) GetDeliveryCount() int {
	return g.DeliveryCount
}

////////////////////////////

func (y *Yandex) GetServiceName() string {
	return y.Name
}

func (y *Yandex) CreateOrder(name string) {
	y.DeliveryCount++
	fmt.Println("New order created!", name)
}

func (y *Yandex) DeleteOrder(name string) {
	if y.DeliveryCount == 0 {
		fmt.Println("Dont have any orders")
		return
	}
	y.DeliveryCount --
	fmt.Println("Order has been deleted!", name)
}

func (y *Yandex) GetDeliveryCount() int {
	return y.DeliveryCount
}

///////////////////////////////////////

func (w *Wolt) GetServiceName() string {
	return w.Name
}

func (w *Wolt) CreateOrder(name string) {
	w.DeliveryCount++
	fmt.Println("New order created!", name)
}

func (w *Wolt) DeleteOrder(name string) {
	if w.DeliveryCount == 0 {
		fmt.Println("Dont have any orders")
		return
	}
	w.DeliveryCount --
	fmt.Println("Order has been deleted!", name)
}

func (w *Wolt) GetDeliveryCount() int {
	return w.DeliveryCount
}

func PrintInfo(d DeliveryService, name string) {
	fmt.Println(d.GetServiceName())

	d.CreateOrder(name)
	fmt.Println("Delivery Count is:", d.GetDeliveryCount())

	d.DeleteOrder(name)
	fmt.Println("Delivery Count is:", d.GetDeliveryCount())

}

func main() {
	glovo := Glovo{
		Name: "Glovo",
		DeliveryCount: 10,
	}

	wolt := Wolt{
		Name: "Wolt",
		DeliveryCount: 10,
	}

	yandex := Yandex{
		Name: "Yandex",
		DeliveryCount: 10,
	}
	PrintInfo(&glovo,"MacDonalds")
	PrintInfo(&wolt, "KFC")
	PrintInfo(&yandex, "Supermarket")

}