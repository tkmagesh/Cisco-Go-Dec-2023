package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("	[main] - deferred")
		// recover()
	}()
	fmt.Println("main started")
	f1()
	fmt.Println("main completed")
}

func f1() {
	/*
		defer func() {
			fmt.Println("	[f1] - deferred - 1")
		}()
		defer func() {
			fmt.Println("	[f1] - deferred - 2")
		}()
		defer func() {
			fmt.Println("	[f1] - deferred - 3")
		}()
	*/

	defer fmt.Println("	[f1] - deferred - 1")
	defer fmt.Println("	[f1] - deferred - 2")
	defer fmt.Println("	[f1] - deferred - 3")
	fmt.Println("f1 started")
	f2()
	fmt.Println("f1 completed")
}

func f2() {
	defer func() {
		fmt.Println("	[f2] - deferred")
	}()
	fmt.Println("f2 started")
	panic(errors.New("dummy error"))
	fmt.Println("f2 completed")
}
