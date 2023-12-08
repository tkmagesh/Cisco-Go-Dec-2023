// streaming using channels

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go genNos(ch)
	for i := 1; i <= 5; i++ {
		fmt.Println(<-ch)
	}
}

func genNos(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i * 10
		time.Sleep(2 * time.Second)
	}

}
