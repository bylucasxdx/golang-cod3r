package main

import "fmt"

type sport interface {
	turnOnBoost()
}

type ferrari struct {
	model          string
	boostOn        bool
	actualVelocity int
}

func (f *ferrari) turnOnBoost() {
	f.boostOn = true
}

func main() {
	car := ferrari{"F40", false, 0}
	car.turnOnBoost()

	var car2 sport = &ferrari{"F40", false, 0}
	car2.turnOnBoost()

	fmt.Println(car, car2)
}
