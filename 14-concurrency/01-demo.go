package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() // scheduling the execution of f1 through the scheduler
	f2()
	// blocking the execution of the main function so that the scheduler can look for other goroutines scheduled and execute them
	time.Sleep(2 * time.Second)
	// fmt.Scanln()
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(4 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
