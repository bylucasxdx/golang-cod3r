package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	grade := 6.9
	finalGrade := int(grade)
	fmt.Println(finalGrade)

	// Tabela asc
	fmt.Println("Test" + string(rune(97)))

	// Convert Int to String
	fmt.Println("Teste" + strconv.Itoa(123))

	num, _ := strconv.Atoi("123")
	fmt.Println(num)

	boolean, _ := strconv.ParseBool("true")
	if boolean {
		fmt.Println("True")
	}
}
