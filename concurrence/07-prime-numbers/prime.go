package main

import (
	"fmt"
	"time"
)

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func primeNumbers(num int, ch chan int) {
	initialNumber := 2

	for i := 0; i < num; i++ {
		for prime := initialNumber; ; prime++ {
			if isPrime(prime) {
				ch <- prime
				initialNumber = prime + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}

	close(ch)
}

func main() {
	ch := make(chan int, 30)

	go primeNumbers(100, ch)
	for primeNumber := range ch {
		fmt.Printf("%d ", primeNumber)
	}
}
