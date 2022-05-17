package structs_methods_and_interfaces

import (
	testingTools "github.com/starsky-mik/kgl-go-learing/internal/pkg/testing_tools"
	structs "github.com/starsky-mik/kgl-go-learing/lesson/structs-methods-and-interfaces"
	"testing"
)

type AreaTestCase struct {
	shape structs.Shape
	value float64
}

func TestArea(t *testing.T) {
	cases := []AreaTestCase{
		{structs.Rectangle{}, 0},
		{structs.Rectangle{Width: 2}, 0},
		{structs.Rectangle{Height: 5}, 0},
		{structs.Rectangle{Width: -1, Height: 5}, 0},
		{structs.Rectangle{Width: 2, Height: -5}, 0},
		{structs.Rectangle{Width: 2, Height: 5}, 10},
		{structs.Circle{Radius: -1}, 0},
		{structs.Circle{}, 0},
		{structs.Circle{Radius: 10}, 314.1592653589793},
		{structs.IsoscelesTriangle{Base: -1, Height: 5}, 0},
		{structs.IsoscelesTriangle{Base: 10, Height: -1}, 0},
		{structs.IsoscelesTriangle{Height: 5}, 0},
		{structs.IsoscelesTriangle{Base: 10}, 0},
		{structs.IsoscelesTriangle{Base: 10, Height: 5}, 25},
	}

	for _, testCase := range cases {
		testingTools.AssertEquals(t, testCase.shape.Area(), testCase.value)
	}
}

func BenchmarkArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		structs.Rectangle{Width: 2.0, Height: 5.0}.Area()
		structs.Circle{Radius: 10}.Area()
		structs.IsoscelesTriangle{Base: 10, Height: 5}.Area()
	}
}
