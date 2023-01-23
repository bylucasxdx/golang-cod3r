package main

import (
	"fmt"
	"time"
)

// Channel is used to make communication between goroutines

func multiply(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	ch := make(chan int)

	go multiply(2, ch)
	first, second := <-ch, <-ch
	fmt.Println(first, second)

	fmt.Println(<-ch)
}
