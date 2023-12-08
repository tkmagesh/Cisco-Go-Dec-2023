package main

import "fmt"

type MyStr string

func (s MyStr) Length() int {
	return len(s)
}

func main() {
	str := MyStr("Quis voluptate proident amet non ipsum ex commodo quis adipisicing tempor Lorem. Dolore Lorem magna officia ullamco tempor ullamco incididunt occaecat laborum qui. Pariatur qui pariatur quis veniam veniam excepteur quis officia. Anim elit duis commodo eiusmod aliquip magna eu velit commodo officia elit. Consequat proident enim proident nostrud aliqua ut labore excepteur sunt. Consequat est aliqua enim commodo incididunt non incididunt incididunt officia minim ipsum. Consequat pariatur labore incididunt sint ipsum nulla do anim ut amet aliquip sint.")
	fmt.Println(str.Length())
}
