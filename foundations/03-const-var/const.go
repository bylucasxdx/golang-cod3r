package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2

	area := PI * math.Pow(raio, 2)
	fmt.Println("Área completa é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(c, d)

	var e, f bool = true, false

	fmt.Println(e, f)

	number, boolType, text := 2, false, "string"

	fmt.Println(number, boolType, text)
}
