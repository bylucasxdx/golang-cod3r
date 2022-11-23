package main

import (
	"fmt"
	"math"
)

func main() {
	number1 := 10
	number2 := 2

	fmt.Println("Sum => ", number1+number2)
	fmt.Println("Subtract => ", number1-number2)
	fmt.Println("Division => ", number1/number2)
	fmt.Println("Multiplication => ", number1*number2)
	fmt.Println("Module => ", number1%number2)

	// Bitwise
	fmt.Println("E => ", number1&number2)
	fmt.Println("Or => ", number1|number2)
	fmt.Println("Xor => ", number1^number2)

	// Math Module
	numberFloat1 := 2.0
	numberFloat2 := 3.0

	fmt.Println("Greater than => ", math.Max(numberFloat1, numberFloat2))
	fmt.Println("Less than => ", math.Min(numberFloat1, numberFloat2))

}
