package structs_methods_and_interfaces

import "math"

type Circle struct {
	Radius float64
}

func (c Circle) valid() bool {
	return c.Radius > 0
}

func (c Circle) Perimeter() float64 {
	if !c.valid() {
		return 0
	}

	return 2 * math.Pi * c.Radius
}

func (c Circle) Area() float64 {
	if !c.valid() {
		return 0
	}

	return math.Pi * c.Radius * c.Radius
}
