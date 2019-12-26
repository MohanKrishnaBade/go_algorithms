package Classes

// Define a new data type "Circle"
type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return 3.14 * (c.Radius * c.Radius)
}
