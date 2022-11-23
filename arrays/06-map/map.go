package main

import "fmt"

func main() {
	// var approveds map[int]string
	approveds := make(map[int]string)

	approveds[12345] = "Lucas"
	approveds[32145] = "Test"
	fmt.Println(approveds)

	for id, name := range approveds {
		fmt.Printf("%s (Id: %d)\n", name, id)
	}

	fmt.Println(approveds[12345])
	delete(approveds, 12345)
	fmt.Println(approveds)
}
