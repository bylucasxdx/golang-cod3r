package main

import (
	"fmt"
	"time"
)

func routine(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	fmt.Println("Execution...")
	ch <- 6
}

func main() {
	ch := make(chan int, 3)
	go routine(ch)

	time.Sleep(time.Second * 5)
	fmt.Println(<-ch)
}
