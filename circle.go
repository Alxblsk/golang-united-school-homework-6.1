package golang_united_school_homework

import (
	"math"
)

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func NewCircle(radius float64) *Circle {
	return &Circle{Radius: radius}
}

func (c Circle) CalcPerimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) CalcArea() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}
