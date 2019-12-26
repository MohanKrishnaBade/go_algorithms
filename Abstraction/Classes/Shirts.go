package Classes

import (
	"fmt"
	"../../../go_algorithms/Abstraction"
)

//A struct Shirts is declared in which the type Product is embedded.
type Shirts struct {
	Abstraction.Product
	Size, Pattern, Sleeve, Fabric string
}

//Overrides for the Attributes Method for the Shirts struct
func (sh Shirts) Attributes() {
	fmt.Println("\nSpecification:")
	fmt.Println(sh.Size)
	fmt.Println(sh.Pattern)
	fmt.Println(sh.Sleeve)
	fmt.Println(sh.Fabric)
}
