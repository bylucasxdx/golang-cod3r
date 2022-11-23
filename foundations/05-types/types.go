package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println("Type of variable 31000 is", reflect.TypeOf(31000))

	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println("Max int value is", i1)
	fmt.Println("Type of i1 is", reflect.TypeOf(i1))

}
