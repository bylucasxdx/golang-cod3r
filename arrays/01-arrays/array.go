package main

import "fmt"

func main() {
	// array is always the same type and static (length fixed)
	var grades [3]float64
	fmt.Println(grades)

	grades[0], grades[1], grades[2] = 1.2, 1.3, 1.5
	fmt.Println(grades)

	total := 0.0
	for i := 0; i < len(grades); i++ {
		total += grades[i]
	}

	median := total / float64(len(grades))
	fmt.Printf("Median: %2.f\n", median)
}
