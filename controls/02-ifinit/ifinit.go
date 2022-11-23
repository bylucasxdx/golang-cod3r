package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumber() int {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	return random.Intn(10)
}

func main() {
	if i := randomNumber(); i > 5 {
		fmt.Println("Winner")
	} else {
		fmt.Println("Loser")
	}
}
