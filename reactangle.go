package golang_united_school_homework

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

func NewRectangle(height, width float64) *Rectangle {
	return &Rectangle{Height: height, Weight: width}
}

func (r Rectangle) CalcPerimeter() float64 {
	return (r.Height + r.Weight) * 2
}

func (r Rectangle) CalcArea() float64 {
	return r.Height * r.Weight
}
