package main

import "fmt"

type Product struct {
	name     string
	price    float64
	discount float64
}

func (product Product) priceWithDiscount() float64 {
	return product.price - (product.price * product.discount)
}

func main() {
	product := Product{
		name:     "Box",
		price:    10.0,
		discount: 0.05,
	}
	fmt.Println(product, product.priceWithDiscount())

	product2 := Product{"Notebook", 1000.0, 0.10}
	fmt.Println(product2, product2.priceWithDiscount())
}
