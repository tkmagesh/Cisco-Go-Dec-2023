package main

import (
	"fmt"
	"log"
)

func main() {
	/*
		add(100, 200)
		subtract(100, 200)
	*/

	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	logOperation(100, 200, add)
	logOperation(100, 200, subtract)
	logOperation(100, 200, func(x, y int) {
		fmt.Println("Multiply Result :", x*y)
	})
}

func logOperation(x, y int, oper func(int, int)) {
	log.Println("Invocation started")
	oper(x, y)
	log.Println("Invocation completed")
}

func logAdd(x, y int) {
	log.Println("Invocation started")
	add(x, y)
	log.Println("Invocation completed")
}

func logSubtract(x, y int) {
	log.Println("Invocation started")
	subtract(x, y)
	log.Println("Invocation completed")
}

func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}
