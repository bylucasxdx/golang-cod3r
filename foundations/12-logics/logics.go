package main

import "fmt"

func isBuyAvailable(job1, job2 bool) (bool, bool, bool) {
	buyTv50 := job1 && job2
	buyTv32 := job1 != job2
	buyIcecream := job1 || job2

	return buyTv50, buyTv32, buyIcecream
}

func main() {
	tv50, tv32, icecream := isBuyAvailable(true, true)
	fmt.Printf("TV 50: %t, TV 32: %t, Icecream: %t", tv50, tv32, icecream)
}
