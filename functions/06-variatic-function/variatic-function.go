package main

import "fmt"

func median(numbers ...float64) float64 {
	total := 0.0

	for _, number := range numbers {
		total += number
	}

	return total / float64(len(numbers))
}

func main() {
	result := median(2, 3, 4, 5, 6)
	fmt.Println(result)
}
