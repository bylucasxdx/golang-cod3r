package main

import "fmt"

type sport interface {
	boostOn()
}

type luxury interface {
	parking()
}

type luxurySport interface {
	sport
	luxury
}

type bmw struct{}

func (b *bmw) boostOn() {
	fmt.Println("Boost on...")
}

func (b *bmw) parking() {
	fmt.Println("Parking...")
}

func main() {
	var b luxurySport = &bmw{}
	b.boostOn()
	b.parking()

}
