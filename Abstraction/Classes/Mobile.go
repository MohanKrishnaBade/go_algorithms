package Classes

import (
	"../../../go_algorithms/Abstraction"
	"fmt"
)

//A struct Mobile is declared in which the type Product is embedded.
type Mobile struct {
	Abstraction.Product
	DisplayFeatures   []string
	ProcessorFeatures []string
}

//Overrides for the Attributes Method for the Mobile struct
func (mob Mobile) Attributes() {
	mob.Product.Attributes()
	fmt.Println("\nDisplay Features:")
	for _, key := range mob.DisplayFeatures {
		fmt.Println(key)
	}
	fmt.Println("\nProcessor Features:")
	for _, key := range mob.ProcessorFeatures {
		fmt.Println(key)
	}
}
