package main

import (
	"fmt"
	"time"
)

func routine(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // block operation
	fmt.Println("Waiting read channel")
}

func main() {
	c := make(chan int) // without buffer

	go routine(c)
	fmt.Println(<-c)
	fmt.Println("Reading...")

	fmt.Println(<-c)
	fmt.Println("End")
}
