package structs_methods_and_interfaces

import (
	"testing"
)

func assertEquals(t testing.TB, actual, expected any) {
	t.Helper()

	if actual != expected {
		t.Errorf("expected '%v' but got '%v'", expected, actual)
	}
}

type TestCase struct {
	shape Shape
	value float64
}

func TestPerimeter(t *testing.T) {
	cases := []TestCase{
		{Rectangle{0, 0}, 0},
		{Rectangle{2, 0}, 0},
		{Rectangle{0, 5}, 0},
		{Rectangle{-1, 5}, 0},
		{Rectangle{2, -5}, 0},
		{Rectangle{2, 5}, 14},
		{Circle{-1}, 0},
		{Circle{0}, 0},
		{Circle{10}, 62.83185307179586},
		{IsoscelesTriangle{-1, 5}, 0},
		{IsoscelesTriangle{10, -1}, 0},
		{IsoscelesTriangle{0, 5}, 0},
		{IsoscelesTriangle{10, 0}, 0},
		{IsoscelesTriangle{10, 5}, 37.32050807568878},
	}

	for _, testCase := range cases {
		assertEquals(t, testCase.shape.Perimeter(), testCase.value)
	}
}

func BenchmarkPerimeter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Rectangle{2.0, 5.0}.Perimeter()
		Circle{10}.Perimeter()
	}
}

func TestArea(t *testing.T) {
	cases := []TestCase{
		{Rectangle{0, 0}, 0},
		{Rectangle{2, 0}, 0},
		{Rectangle{0, 5}, 0},
		{Rectangle{-1, 5}, 0},
		{Rectangle{2, -5}, 0},
		{Rectangle{2, 5}, 10},
		{Circle{-1}, 0},
		{Circle{0}, 0},
		{Circle{10}, 314.1592653589793},
		{IsoscelesTriangle{-1, 5}, 0},
		{IsoscelesTriangle{10, -1}, 0},
		{IsoscelesTriangle{0, 5}, 0},
		{IsoscelesTriangle{10, 0}, 0},
		{IsoscelesTriangle{10, 5}, 25},
	}

	for _, testCase := range cases {
		assertEquals(t, testCase.shape.Area(), testCase.value)
	}
}

func BenchmarkArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Rectangle{2.0, 5.0}.Area()
		Circle{10}.Area()
	}
}
