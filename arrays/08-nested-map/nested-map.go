package main

import "fmt"

func main() {
	employeesByChar := map[string]map[string]float64{
		"L": {
			"Lucas": 1234.0,
			"Lara":  12333.0,
		},
		"M": {
			"Maria": 11111.22,
		},
	}

	fmt.Println(employeesByChar)
	fmt.Println(employeesByChar["L"])

	for char, employee := range employeesByChar {
		fmt.Println(char, employee)
		for name, salary := range employee {
			fmt.Println(name, salary)
		}
	}
}
