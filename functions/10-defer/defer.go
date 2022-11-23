package main

import "fmt"

func testDefer() {
	defer fmt.Println("Last execution line")
	defer fmt.Println("Other line")
	fmt.Println("Test")
}

func main() {
	testDefer()
}
