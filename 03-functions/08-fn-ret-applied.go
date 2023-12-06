package main

import (
	"fmt"
	"log"
	"time"
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

	// composing log capability
	/*
		logAdd := getLogOperation(add)
		logAdd(100, 200)

		logSubtract := getLogOperation(subtract)
		logSubtract(100, 200)

		getLogOperation(func(x, y int) {
			fmt.Println("Multiply result :", x*y)
		})(100, 200)
	*/

	// composing profile capability
	/*
		profiledAdd := getProfiledOperation(add)
		profiledAdd(100, 200)

		profiledSubtract := getProfiledOperation(subtract)
		profiledSubtract(100, 200)
	*/

	// composing log & profile capability
	logAdd := getLogOperation(add)
	profiledLogAdd := getProfiledOperation(logAdd)
	profiledLogAdd(100, 200)

	logSubtract := getLogOperation(subtract)
	profiledLogSubtract := getProfiledOperation(logSubtract)
	profiledLogSubtract(100, 200)

	getProfiledOperation(getLogOperation(func(x, y int) {
		fmt.Println("Multiply Result :", x*y)
	}))(100, 200)

}
func getProfiledOperation(oper func(int, int)) func(int, int) {
	return func(x, y int) {
		start := time.Now()
		oper(x, y)
		elapsed := time.Since(start)
		fmt.Println("Elapsed :", elapsed)
	}
}

func getLogOperation(oper func(int, int)) func(int, int) {
	return func(x, y int) {
		log.Println("Invocation started")
		oper(x, y)
		log.Println("Invocation completed")
	}
}

/*
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
*/

func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}
