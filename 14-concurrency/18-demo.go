package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := genFib()
	for no := range ch {
		fmt.Println(no)
	}
	fmt.Println("Done")
}

// update this function to generate fibonocci series ( rand.Intn(20) count )
func genFib() <-chan int {
	ch := make(chan int)
	go func() {
		for x, y, count := 0, 1, rand.Intn(20); count >= 0; count-- {
			ch <- x
			time.Sleep(500 * time.Millisecond)
			x, y = y, x+y
		}
		close(ch)
	}()
	return ch
}
