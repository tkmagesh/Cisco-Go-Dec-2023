package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

type Dummy struct {
	Id int
}

type PerishableProduct struct {
	Product // inheritance (composition) coz an attr name is not explicitly mentioned
	// Dummy
	Expiry string
	Id     string // overriding the Product.Id
}

func NewPerishableProduct(id string, name string, cost float32, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			Name: name,
			Cost: cost,
		},
		Expiry: expiry,
		Id:     id,
	}
}

func main() {
	// milk := PerishableProduct{Product{100, "Milk", 50}, "2 Days"}
	milk := PerishableProduct{
		Product: Product{
			Id:   100,
			Name: "Milk",
			Cost: 50,
		},
		Expiry: "2 Days",
		Id:     "milk-101",
	}
	fmt.Printf("%#v\n", milk)
	// fmt.Println(milk.Product.Id, milk.Product.Name)
	fmt.Println(milk.Id, milk.Name)

	grapes := NewPerishableProduct("G-102", "Grapes", 100, "2 Days")
	fmt.Println(grapes)
}
