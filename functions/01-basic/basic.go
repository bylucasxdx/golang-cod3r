package main

import "fmt"

func f1() {
	fmt.Println("Hello world!")
}

func f2(param1 string, param2 float64) {
	fmt.Println(param1, param2)
}

func f3() string {
	return "string function 3"
}

func f4(p1, p2 string) string {
	return fmt.Sprint(p1, p2)
}

func f5() (string, string) {
	return "First string", "Second string"
}

func main() {
	f1()
	f2("Lucas", 123)

	r3 := f3()
	fmt.Println(r3)

	r4 := f4("A", "B")
	fmt.Println(r4)

	r51, f52 := f5()
	fmt.Println(r51, f52)
}
