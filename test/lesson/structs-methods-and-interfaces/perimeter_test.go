package structs_methods_and_interfaces

import (
	. "github.com/starsky-mik/kgd-go-learing/internal/pkg/testing_tools"
	. "github.com/starsky-mik/kgd-go-learing/lesson/structs-methods-and-interfaces"
	"testing"
)

type PerimeterTestCase struct {
	shape Shape
	value float64
}

func TestPerimeter(t *testing.T) {
	cases := []PerimeterTestCase{
		{Rectangle{}, 0},
		{Rectangle{Width: 2}, 0},
		{Rectangle{Height: 5}, 0},
		{Rectangle{Width: -1, Height: 5}, 0},
		{Rectangle{Width: 2, Height: -5}, 0},
		{Rectangle{Width: 2, Height: 5}, 14},
		{Circle{Radius: -1}, 0},
		{Circle{}, 0},
		{Circle{Radius: 10}, 62.83185307179586},
		{IsoscelesTriangle{Base: -1, Height: 5}, 0},
		{IsoscelesTriangle{Base: 10, Height: -1}, 0},
		{IsoscelesTriangle{Height: 5}, 0},
		{IsoscelesTriangle{Base: 10}, 0},
		{IsoscelesTriangle{Base: 10, Height: 5}, 37.32050807568878},
	}

	for _, testCase := range cases {
		AssertEquals(t, testCase.shape.Perimeter(), testCase.value)
	}
}

func BenchmarkPerimeter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Rectangle{Width: 2.0, Height: 5.0}.Perimeter()
		Circle{Radius: 10}.Perimeter()
		IsoscelesTriangle{Base: 10, Height: 5}.Perimeter()
	}
}
