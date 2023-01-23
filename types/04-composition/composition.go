package main

import "fmt"

type Car struct {
	name     string
	velocity int
}

type Ferrari struct {
	Car
	boost bool
}

func main() {
	f := Ferrari{}
	f.name = "F40"
	f.velocity = 0
	f.boost = true

	fmt.Println(f)
}
