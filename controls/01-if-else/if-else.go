package main

import "fmt"

func printResult(grade float64) {
	if grade >= 7 {
		fmt.Println("Approved:", grade)
	} else if grade >= 5 {
		fmt.Println("Recovery:", grade)
	} else {
		fmt.Println("Reproved:", grade)
	}
}

func main() {
	printResult(6.5)
	printResult(4.9)
	printResult(7.2)
}
