package main

import "fmt"

type item struct {
	productID int
	quantity  int
	price     float64
}

type order struct {
	userID int
	items  []item
}

func (o order) calculateTotalValue() float64 {
	total := 0.00

	for _, item := range o.items {
		total += item.price * float64(item.quantity)
	}

	return total
}

func main() {
	order := order{
		userID: 1,
		items: []item{
			{productID: 1, quantity: 1, price: 2.0},
			{productID: 2, quantity: 3, price: 3.0},
		},
	}

	fmt.Println("Total", order.calculateTotalValue())
}
