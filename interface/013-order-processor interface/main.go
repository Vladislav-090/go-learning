package main

import "fmt"

type OrderProcessor interface {
	CreateOrder(name string)
	DeleteOrder(name string)
	GetOrderCount() int
	GetServiceName() string
}

type OnlineShop struct {
	Name        string
	OrdersCount int
}

type Taxi struct {
	Name        string
	OrdersCount int
}

type FoodDelivery struct {
	Name        string
	OrdersCount int
}

func (o *OnlineShop) CreateOrder(name string) {
	o.OrdersCount++
	fmt.Println("Order has been created!", name)
}

func (o *OnlineShop) DeleteOrder(name string) {
	if o.OrdersCount == 0 {
		fmt.Println("dont have any order to delete...")
		return
	}
	o.OrdersCount --
	fmt.Println("Order has been deleted!")
}

func (o *OnlineShop) GetOrderCount() int {
	return o.OrdersCount
}

func (o *OnlineShop) GetServiceName() string {
	return o.Name
}

///////////////////////////////////////////
func (t *Taxi) CreateOrder(name string) {
	t.OrdersCount++
	fmt.Println("Order has been created!", name)
}

func (t *Taxi) DeleteOrder(name string) {
	if t.OrdersCount == 0 {
		fmt.Println("dont have any order to delete...")
		return
	}
	t.OrdersCount --
	fmt.Println("Order has been deleted!")
}

func (t *Taxi) GetOrderCount() int {
	return t.OrdersCount
}

func (t *Taxi) GetServiceName() string {
	return t.Name
}

///////////////////////////////////////

func (f *FoodDelivery) CreateOrder(name string) {
	f.OrdersCount++
	fmt.Println("Order has been created!", name)
}

func (f *FoodDelivery) DeleteOrder(name string) {
	if f.OrdersCount == 0 {
		fmt.Println("dont have any order to delete...")
		return
	}
	f.OrdersCount --
	fmt.Println("Order has been deleted!")
}

func (f *FoodDelivery) GetOrderCount() int {
	return f.OrdersCount
}

func (f *FoodDelivery) GetServiceName() string {
	return f.Name
}

func PrintInfo(or OrderProcessor, name string) {
	fmt.Println(or.GetServiceName())

	or.CreateOrder(name)
	fmt.Println("Orders count is :", or.GetOrderCount())

	or.DeleteOrder(name)
	fmt.Println("Orders count is :", or.GetOrderCount())
}


func main() {
	onlineshop := OnlineShop{
		Name: "OnlineShop",
		OrdersCount: 10,
	}

	taxi := Taxi{
		Name: "9 region",
		OrdersCount: 45,
	}

	foodDelivery:= FoodDelivery{
		Name: "KFC",
		OrdersCount: 8,
	}

	PrintInfo(&onlineshop, "T-Shirt")
	PrintInfo(&taxi, "Astana-Karagandy /324 km/")
	PrintInfo(&foodDelivery, "1 combo, 6 wings, 2 coke")
}