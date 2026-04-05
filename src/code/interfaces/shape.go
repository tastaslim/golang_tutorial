package interfaces

import "math"

type Shape interface {
	Area() float64
}

type Square struct {
	Length, Breadth float64
}

type Circle struct {
	Radius float64
}

func (s Square) Area() float64 {
	return s.Breadth * s.Length
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func CalculateArea(s Shape) float64 {
	return s.Area()
}
