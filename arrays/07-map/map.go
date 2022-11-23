package main

import "fmt"

func main() {
	employees := map[string]float64{
		"Lucas": 1122.2,
		"Maria": 1234.1,
	}

	employees["Clarissa"] = 1345
	fmt.Println(employees)
	delete(employees, "Teste")

	for name, salary := range employees {
		fmt.Println(name, salary)
	}

}
