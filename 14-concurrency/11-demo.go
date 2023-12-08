/* channel behaviors */
package main

import "fmt"

func main() {
	ch := make(chan int)
	// modify the below in such a way that the receive operation is performed in a goroutine
	go func() {
		ch <- 100
	}()
	data := <-ch
	fmt.Println(data)
}
