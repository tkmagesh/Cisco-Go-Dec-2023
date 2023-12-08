package main

import "fmt"

type Organization struct {
	Name string
	City string
}

type Employee struct {
	Id   int
	Name string
	Org  *Organization
}

func main() {
	cisco := Organization{Name: "CISCO", City: "Bengaluru"}

	e1 := Employee{Id: 100, Name: "Magesh", Org: &cisco}
	e2 := Employee{Id: 101, Name: "Suresh", Org: &cisco}

	fmt.Println(e1.Org)
	fmt.Println(e2.Org)

	// Change the city of the Organization
	fmt.Println("After changing the city of cisco")
	cisco.City = "NCR"
	fmt.Println(e1.Org)
	fmt.Println(e2.Org)
}

func fn(e Employee) {

}
