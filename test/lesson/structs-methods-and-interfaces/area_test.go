package structs_methods_and_interfaces

import (
	. "github.com/starsky-mik/kgd-go-learing/internal/pkg/testing_tools"
	. "github.com/starsky-mik/kgd-go-learing/lesson/structs-methods-and-interfaces"
	"testing"
)

type AreaTestCase struct {
	shape Shape
	value float64
}

func TestArea(t *testing.T) {
	cases := []AreaTestCase{
		{Rectangle{}, 0},
		{Rectangle{Width: 2}, 0},
		{Rectangle{Height: 5}, 0},
		{Rectangle{Width: -1, Height: 5}, 0},
		{Rectangle{Width: 2, Height: -5}, 0},
		{Rectangle{Width: 2, Height: 5}, 10},
		{Circle{Radius: -1}, 0},
		{Circle{}, 0},
		{Circle{Radius: 10}, 314.1592653589793},
		{IsoscelesTriangle{Base: -1, Height: 5}, 0},
		{IsoscelesTriangle{Base: 10, Height: -1}, 0},
		{IsoscelesTriangle{Height: 5}, 0},
		{IsoscelesTriangle{Base: 10}, 0},
		{IsoscelesTriangle{Base: 10, Height: 5}, 25},
	}

	for _, testCase := range cases {
		AssertEquals(t, testCase.shape.Area(), testCase.value)
	}
}

func BenchmarkArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Rectangle{Width: 2.0, Height: 5.0}.Area()
		Circle{Radius: 10}.Area()
		IsoscelesTriangle{Base: 10, Height: 5}.Area()
	}
}
