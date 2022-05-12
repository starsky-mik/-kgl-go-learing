package structs_methods_and_interfaces

type Shape interface {
	Perimeter() float64
	Area() float64
	valid() bool
}
