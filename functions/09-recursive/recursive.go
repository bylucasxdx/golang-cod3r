package main

import (
	"fmt"
)

func factorial(number int) (int, error) {
	switch {
	case number < 0:
		return -1, fmt.Errorf("invalid number %d", number)
	case number == 0:
		return 1, nil
	default:
		lastFactorial, _ := factorial(number - 1)
		return number * lastFactorial, nil
	}
}

func positiveFactorial(number uint) uint {
	switch {
	case number == 0:
		return 1
	default:
		lastFactorial := positiveFactorial(number - 1)
		return number * lastFactorial
	}
}

func main() {
	_, err := factorial(-1)
	fmt.Println(err)

	result, _ := factorial(5)
	fmt.Println(result)

	positiveResult := positiveFactorial(5)
	fmt.Println(positiveResult)
}
