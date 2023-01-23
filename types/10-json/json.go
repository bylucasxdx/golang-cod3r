package main

import (
	"encoding/json"
	"fmt"
)

type product struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

func main() {
	p1 := product{1, "Notebook", 1000.99, []string{"Promotion", "Eletronic"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	var p2 product
	p2JsonString := `{"id":1,"name":"Notebook","price":1000.99,"tags":["Promotion","Eletronic"]}`
	json.Unmarshal([]byte(p2JsonString), &p2)
	fmt.Println(p2.Tags)
}
