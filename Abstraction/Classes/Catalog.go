package Classes

import (
	"fmt"
	"../../Interface"
)

// A struct named Catalog is declared to represent catalog of various products.
// The type of Details field uses a slice of the Information interface
type Catalog struct {
	Name    string
	Details []Interface.Information
}

func (c Catalog) CatalogDetails() {


	fmt.Printf("\nStore : %s \n", c.Name)
	fmt.Println("===================Products==================")
	for _, key := range c.Details {
		key.General()
		key.Attributes()
		key.Inventory()
		fmt.Println("=============================================")
	}
}
