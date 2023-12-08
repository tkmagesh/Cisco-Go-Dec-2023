/* deadlock */

package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(10)
	go fn(wg)
	wg.Wait()
}

func fn(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("fn invoked")
}
