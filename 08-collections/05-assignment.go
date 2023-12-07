/* Write a function that returns all the prime numbers beween the given start and end
print the generated prime numbers in the main function
*/

package main

import "fmt"

func main() {
	primes := genPrimes(2, 100)
	fmt.Println(primes)
}

func genPrimes(start, end int) []int {
	primes := make([]int, 0)
LOOP:
	for no := start; no <= end; no++ {
		for i := 2; i <= (no / 2); i++ {
			if no%i == 0 {
				continue LOOP
			}
		}
		primes = append(primes, no)
	}
	return primes
}
