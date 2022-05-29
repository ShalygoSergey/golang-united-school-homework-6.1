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
	b := box{
		shapesCapacity: shapesCapacity,
	}
	//b.shapes = make([]Shape, shapesCapacity)
	return &b
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if cap(b.shapes) >= b.shapesCapacity {
		return errors.New("AddShape error")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if len(b.shapes) <= i {
		return nil, errors.New("GetByIndex error 1")
	}
	shape := b.shapes[i]
	if shape == nil {
		return nil, errors.New("GetByIndex error 2")
	}
	return shape, nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if len(b.shapes) <= i {
		return nil, errors.New("ExtractByIndex error 1")
	}
	shape := b.shapes[i]
	if shape == nil {
		return nil, errors.New("ExtractByIndex error 2")
	}
	b.RemoveElement(i)
	return shape, nil
}

func (b *box) RemoveElement(i int) {
	copy(b.shapes[i:], b.shapes[i+1:])
	b.shapes[len(b.shapes)-1] = nil
	b.shapes = b.shapes[:len(b.shapes)-1]
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if len(b.shapes) <= i {
		return nil, errors.New("ReplaceByIndex error 1")
	}
	replacedShape := b.shapes[i]
	if replacedShape == nil {
		return nil, errors.New("ReplaceByIndex error 2")
	}
	b.shapes[i] = shape
	return replacedShape, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var perimeter float64 = 0
	for _, v := range b.shapes {
		if v != nil {
			perimeter = perimeter + v.CalcPerimeter()
		}
	}
	return perimeter
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var area float64 = 0
	for _, v := range b.shapes {
		if v != nil {
			area = area + v.CalcArea()
		}
	}
	return area
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	exist := false
	for i := 1; i < len(b.shapes); i++ {
		v := b.shapes[i]
		if v != nil {
			switch v.(type) {
			case *Circle:
				b.RemoveElement(i)
				i = i - 1
				exist = true
			}
		}
	}
	if !exist {
		return errors.New("RemoveAllCircles")
	}
	return nil
}
