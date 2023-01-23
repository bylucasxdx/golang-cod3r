package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	// send data to channel
	ch <- 1

	// receive data from channel
	<-ch

	ch <- 2
	fmt.Println(<-ch)
}
