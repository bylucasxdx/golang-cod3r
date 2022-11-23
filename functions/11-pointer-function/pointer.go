package main

import "fmt"

func firstIncrement(number int) {
	number++
}

func secondIncrement(number *int) {
	*number++
}

func main() {
	number := 1

	firstIncrement(number)
	fmt.Println(number)

	secondIncrement(&number)
	fmt.Println(number)
}
