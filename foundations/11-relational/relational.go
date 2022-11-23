package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Strings:", "Banana" == "Banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println(">=", 3 >= 2)
	fmt.Println("<=", 3 <= 2)

	date1 := time.Unix(0, 0)
	date2 := time.Unix(0, 0)
	fmt.Println("Dates:", date1 == date2)
	fmt.Println("Dates:", date1.Equal(date2))

	type Person struct {
		Nome string
	}

	person1 := Person{"Lucas"}
	person2 := Person{"Lucas"}
	fmt.Println("Struct:", person1 == person2)
}
