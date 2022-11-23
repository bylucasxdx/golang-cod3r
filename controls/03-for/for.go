package main

import "fmt"

func main() {
	value := 1
	for value <= 10 {
		fmt.Printf("%d ", value)
		value++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}
