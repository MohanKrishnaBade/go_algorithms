package Interface

/*
The interface type Information is a contract for creating
various Product types in a catalog. The Information interface
 provides three behaviors in its contract: General,Attributes and Inventory.
*/
type Information interface {
	General()
	Attributes()
	Inventory()
}
