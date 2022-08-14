package shape

type Square struct {
	side float64
}

func NewSquare(side float64) Square {
	if side <= 0 {
		panic("side of the square cannot be negative or zero")
	}
	return Square{side: side}
}

func (square Square) FindPerimeter() float64 {

	if square.side == 1.0 {
		return 4.00
	}
	return 4 * square.side
}

func (square Square) FindArea() float64 {
	if square.side == 1.0 {
		return 1.0
	}
	return square.side * square.side
}
