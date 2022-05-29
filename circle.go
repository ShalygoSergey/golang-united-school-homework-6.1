package golang_united_school_homework

import "math"

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (o Circle) CalcPerimeter() float64 {
	return 2 * math.Pi * o.Radius
}

func (o Circle) CalcArea() float64 {
	return math.Pi * o.Radius * o.Radius
}
