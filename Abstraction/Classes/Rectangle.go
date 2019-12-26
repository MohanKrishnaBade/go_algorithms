package Classes

// Define a new data type "Rectangle"
type Rectangle struct {
	Length, Width float32
}

func (r Rectangle) Area() float32 {
	return r.Length * r.Width
}
