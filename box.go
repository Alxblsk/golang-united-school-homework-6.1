package golang_united_school_homework

import (
	"errors"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if b.shapesCapacity <= len(b.shapes) {
		return errors.New("Capacity is out")
	}

	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i >= 0 && len(b.shapes) > i {
		return b.shapes[i], nil
	}

	return nil, errors.New("Shape is not available by this index")
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	_, err := b.GetByIndex(i)

	if err == nil {
		shapeToExtract := b.shapes[i]
		b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
		return shapeToExtract, nil
	}

	return nil, errors.New("Shape can't be extracted by this index")
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	_, err := b.GetByIndex(i)

	if err == nil {
		extracted := b.shapes[i]
		b.shapes[i] = shape
		return extracted, nil
	}

	return nil, errors.New("Shape can't be replaced by this index")

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var counter float64 = 0

	for _, v := range b.shapes {
		counter = counter + v.CalcPerimeter()
	}

	return counter
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var counter float64 = 0

	for _, v := range b.shapes {
		counter = counter + v.CalcArea()
	}

	return counter
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	circlesFound := 0
	cleanedShapes := []Shape{}

	for _, shape := range b.shapes {
		if _, ok := shape.(*Circle); ok {
			circlesFound += 1
		} else {
			cleanedShapes = append(cleanedShapes, shape)
		}
	}

	b.shapes = cleanedShapes

	if circlesFound == 0 {
		return errors.New("There are no Circles in the list")
	}

	return nil
}
