package structs_methods_and_interfaces

import (
	testingTools "github.com/starsky-mik/kgl-go-learing/internal/pkg/testing_tools"
	structs "github.com/starsky-mik/kgl-go-learing/lesson/structs-methods-and-interfaces"
	"testing"
)

type PerimeterTestCase struct {
	shape structs.Shape
	value float64
}

func TestPerimeter(t *testing.T) {
	cases := []PerimeterTestCase{
		{structs.Rectangle{}, 0},
		{structs.Rectangle{Width: 2}, 0},
		{structs.Rectangle{Height: 5}, 0},
		{structs.Rectangle{Width: -1, Height: 5}, 0},
		{structs.Rectangle{Width: 2, Height: -5}, 0},
		{structs.Rectangle{Width: 2, Height: 5}, 14},
		{structs.Circle{Radius: -1}, 0},
		{structs.Circle{}, 0},
		{structs.Circle{Radius: 10}, 62.83185307179586},
		{structs.IsoscelesTriangle{Base: -1, Height: 5}, 0},
		{structs.IsoscelesTriangle{Base: 10, Height: -1}, 0},
		{structs.IsoscelesTriangle{Height: 5}, 0},
		{structs.IsoscelesTriangle{Base: 10}, 0},
		{structs.IsoscelesTriangle{Base: 10, Height: 5}, 37.32050807568878},
	}

	for _, testCase := range cases {
		testingTools.AssertEquals(t, testCase.shape.Perimeter(), testCase.value)
	}
}

func BenchmarkPerimeter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		structs.Rectangle{Width: 2.0, Height: 5.0}.Perimeter()
		structs.Circle{Radius: 10}.Perimeter()
		structs.IsoscelesTriangle{Base: 10, Height: 5}.Perimeter()
	}
}
