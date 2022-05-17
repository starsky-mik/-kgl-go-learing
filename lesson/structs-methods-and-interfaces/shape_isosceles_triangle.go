package structs_methods_and_interfaces

import "math"

type IsoscelesTriangle struct {
	Base   float64
	Height float64
}

func (t IsoscelesTriangle) valid() bool {
	return t.Base > 0 && t.Height > 0
}

func (t IsoscelesTriangle) Perimeter() float64 {
	if !t.valid() {
		return 0
	}

	return (2 * math.Sqrt(t.Base*t.Base-t.Height*t.Height)) + (2 * t.Base)
}

func (t IsoscelesTriangle) Area() float64 {
	if !t.valid() {
		return 0
	}

	return 0.5 * t.Base * t.Height
}
