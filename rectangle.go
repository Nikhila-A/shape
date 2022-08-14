package shape

import "math"

type Rectangle struct {
	length float64
	width  float64
}

func NewRectangle(length float64, width float64) Rectangle {
	if length <= 0 || width <= 0 {
		panic("length or width should be positive")
	}
	return Rectangle{length: length, width: width}
}

func (rect Rectangle) FindPerimeter() float64 {
	if rect.length == 1 {
		return 2 + 2*rect.width
	} else if rect.width == 1 {
		return 2 + 2*rect.length
	} else if rect.length == rect.width {
		return rect.length * 4
	}
	return 2 * (rect.length + rect.width)
}

func (rect Rectangle) FindArea() float64 {
	if rect.length == 1 {
		return rect.width
	} else if rect.width == 1 {
		return rect.length
	} else if rect.length == rect.width {
		return math.Pow(rect.length, 2)
	}
	return rect.length * rect.width
}
