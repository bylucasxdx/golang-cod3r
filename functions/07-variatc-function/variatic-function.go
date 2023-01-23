package main

import "fmt"

func printApproveds(approveds ...string) {
	fmt.Println("List")

	for i, approved := range approveds {
		fmt.Printf("%d) %s\n", i+1, approved)
	}
}

func main() {
	name := "Maria"
	approveds := []string{"Lucas", "Clarissa", name}
	printApproveds(approveds...)
}
