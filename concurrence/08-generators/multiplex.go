package main

import "fmt"

func send(origin <-chan string, destination chan string) {
	for {
		destination <- <-origin
	}
}

func join(e1, e2 <-chan string) <-chan string {
	ch := make(chan string)
	go send(e1, ch)
	go send(e2, ch)
	return ch
}

func main() {
	ch := join(
		title("https://coinmarketcap.com/currencies/wemix/", "https://www.google.com"),
		title("https://www.google.com", "https://www.cod3r.com.br"),
	)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
