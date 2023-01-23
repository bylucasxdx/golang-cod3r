package main

import (
	"fmt"
	"time"
)

func speak(person, text string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteration %d)\n", person, text, i)
	}
}

func main() {
	// Serial
	speak("Lucas", "Hi", 3)
	speak("João", "Hello", 3)

	// Concurrent
	go speak("Lucas", "Hi", 10)
	go speak("João", "Hello", 10)

	go speak("Lucas", "Hi", 10)
	speak("João", "Hello", 5)
}
