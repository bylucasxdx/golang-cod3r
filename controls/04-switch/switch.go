package main

import (
	"fmt"
	"time"
)

func gradeConcept(number float64) string {
	var grade = int(number)

	switch grade {
	case 10:
		fallthrough
	case 9:
		return "A"
	default:
		return "Invalid"
	}
}

func greeting() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 18:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good night")
	}
}

func types(inteface interface{}) string {
	switch inteface.(type) {
	case int:
		return "integer"
	case float32, float64:
		return "float"
	case string:
		return "string"
	case func():
		return "function"
	default:
		return "invalid"
	}
}

func main() {
	fmt.Println(gradeConcept(10.9))
	greeting()

	fmt.Println(types(1))
	fmt.Println(types(2.0))
	fmt.Println(types("Test"))
	fmt.Println(types(func() {}))
	fmt.Println(types(time.Now()))
}
