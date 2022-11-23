package main

import "fmt"

func main() {
	integerNumber1 := 1

	// Go não tem aritmética de ponteiro
	var pointer *int = nil

	// Pega o endereço da variável
	pointer = &integerNumber1
	*pointer++
	integerNumber1++

	fmt.Println(pointer, *pointer, integerNumber1, &integerNumber1)
}
