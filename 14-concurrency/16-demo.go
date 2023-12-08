// streaming using channels

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go genNos(ch)
	for {
		if val, isOpen := <-ch; isOpen {
			fmt.Println(val)
			continue
		}
		break
	}
}

func genNos(ch chan int) {
	for i := 1; i <= 5; i++ {
		time.Sleep(1 * time.Second)
		ch <- i * 10
	}
	close(ch)

}
