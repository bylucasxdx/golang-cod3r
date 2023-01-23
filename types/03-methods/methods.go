package main

import (
	"fmt"
	"strings"
)

type Person struct {
	name    string
	surname string
}

func (p Person) getFullName() string {
	return p.name + " " + p.surname
}

func (p *Person) setFullName(fullname string) {
	parts := strings.Split(fullname, " ")
	p.name = parts[0]
	p.surname = parts[1]
}

func main() {
	person := Person{
		name:    "Lucas",
		surname: "Medeiros",
	}

	fmt.Println(person)
	fmt.Println(person.getFullName())

	person.setFullName("Teste Sobrenome")

	fmt.Println(person.getFullName())
}
