package main

import "fmt"

type printable interface {
	toString() string
}

type person struct {
	name    string
	surname string
}

type product struct {
	name  string
	price float64
}

func (p person) toString() string {
	return p.name + " " + p.surname
}

func (p product) toString() string {
	return fmt.Sprintf("%s - R$ %2.f", p.name, p.price)
}

func print(p printable) {
	fmt.Println(p.toString())
}

func main() {
	var thing printable = person{"Lucas", "Medeiros"}
	print(thing)

	thing = product{"Forno", 30.2}
	print(thing)
}
