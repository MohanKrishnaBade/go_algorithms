package Classes

type Triangle struct {
	Base, Height float32
}

func (t Triangle) Area() float32 {
	return 0.5 * t.Base * t.Height
}
