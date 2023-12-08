package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

// fmt.Stringer interface implementation
func (product Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f\n", product.Id, product.Name, product.Cost)
}

func main() {

	// the below is a standard approach
	var pen Product = Product{
		Id:   100,
		Name: "Pen",
		Cost: 5,
	}

	fmt.Println(pen)

}

func applyDiscount(product *Product, discountPercentage float32) {
	product.Cost = product.Cost * ((100 - discountPercentage) / 100)
}
