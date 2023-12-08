/* resource utilization */

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var count int
	flag.IntVar(&count, "count", 0, "# of goroutines")
	flag.Parse()
	wg := &sync.WaitGroup{}
	fmt.Printf("Starting %d goroutines.. Hit ENTER to start\n", count)
	fmt.Scanln()
	for i := 1; i <= count; i++ {
		wg.Add(1)    // increment the wg counter by 1
		go fn(wg, i) // scheduling the execution of f1 through the scheduler
	}
	wg.Wait() // block until the counter becomes 0 (default = 0)
	fmt.Println("Hit ENTER to shutdown...")
	fmt.Scanln()
}

func fn(wg *sync.WaitGroup, id int) {
	defer wg.Done() // decrement the wg counter by 1
	fmt.Printf("fn[%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	fmt.Printf("fn[%d] completed\n", id)
}
