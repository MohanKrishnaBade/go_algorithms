package Abstraction

import "fmt"

/*
A struct Product is declared with fields for holding its state and methods
implemented based on the behaviors defined in the Information interface. */
type Product struct {
	Name, Description string
	Weight, Price     float64
	Stock             int
}

func (prd Product) General() {
	fmt.Printf("\n%s", prd.Name)
	fmt.Printf("\n%s\n", prd.Description)
	fmt.Println(prd.Price)
}

func (prd Product) Attributes() {
	fmt.Println(prd.Weight)
}

func (prd Product) Inventory() {
	fmt.Println(prd.Stock)
}
