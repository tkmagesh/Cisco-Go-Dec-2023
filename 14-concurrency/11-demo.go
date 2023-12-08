/* channel behaviors */
package main

import "fmt"

func main() {

	// modify the below in such a way that the receive operation is performed in a goroutine
	/*
		ch := make(chan int)
		go func() {
			ch <- 100
		}()
		data := <-ch
		fmt.Println(data)
	*/
	/*
		ch := make(chan int)
		go func() {
			data := <-ch
			fmt.Println(data)
		}()
		ch <- 100
	*/

	ch := make(chan int, 1)
	ch <- 10
	data := <-ch
	fmt.Println(data)
}
