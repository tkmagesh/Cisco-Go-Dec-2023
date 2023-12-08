package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

func (product Product) Print() {
	fmt.Printf("Id = %d, Name = %q, Cost = %0.2f\n", product.Id, product.Name, product.Cost)
}

func (product *Product) applyDiscount(discountPercentage float32) {
	product.Cost = product.Cost * ((100 - discountPercentage) / 100)
}

type PerishableProduct struct {
	Product // inheritance (composition) coz an attr name is not explicitly mentioned
	Expiry  string
}

func NewPerishableProduct(id int, name string, cost float32, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			Id:   id,
			Name: name,
			Cost: cost,
		},
		Expiry: expiry,
	}
}

// Overriding the "Print()" method of the Product (in PerishableProduct)
func (pp PerishableProduct) Print() {
	fmt.Printf("Id = %d, Name = %q, Cost = %0.2f, Expiry = %q\n", pp.Id, pp.Name, pp.Cost, pp.Expiry)
}

func main() {
	grapes := NewPerishableProduct(102, "Grapes", 100, "2 Days")
	grapes.Print() // invoking the Print() method of the "Product" through PerishableProduct
	grapes.applyDiscount(10)
	grapes.Print()
}
