package main

import (
	"fmt"

	"github.com/fatih/color"
	// "github.com/tkmagesh/cisco-go-dec-2023/09-modularity/calculator"
	calc "github.com/tkmagesh/cisco-go-dec-2023/09-modularity/calculator"
)

func main() {
	color.Red("modularity app executed")

	/*
		fmt.Println(calculator.Add(100, 200))
		fmt.Println(calculator.Subtract(100, 200))
		fmt.Println(calculator.Multiply(100, 200))
		fmt.Println("operation count :", calculator.OpCount())
	*/

	fmt.Println(calc.Add(100, 200))
	fmt.Println(calc.Subtract(100, 200))
	fmt.Println(calc.Multiply(100, 200))
	fmt.Println("operation count :", calc.OpCount())

}
