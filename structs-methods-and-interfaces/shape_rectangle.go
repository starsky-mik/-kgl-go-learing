package structs_methods_and_interfaces

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) valid() bool {
	return r.Width > 0 && r.Height > 0
}

func (r Rectangle) Perimeter() float64 {
	if !r.valid() {
		return 0
	}

	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	if !r.valid() {
		return 0
	}

	return r.Width * r.Height
}
