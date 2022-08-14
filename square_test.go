package shape

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSquare(t *testing.T) {
	t.Run("check initialization of square", func(t *testing.T) {
		assert.IsType(t, Square{}, NewSquare(2.5))
	})

	t.Run("check initialization of square with negative side", func(t *testing.T) {
		assert.Panics(t, func() {
			NewSquare(-5.0)
		})
	})

	t.Run("check initialization of square with zero side", func(t *testing.T) {
		assert.Panics(t, func() {
			NewSquare(0.0)
		})
	})

	t.Run("check initialization of square with positive side", func(t *testing.T) {
		assert.NotPanics(t, func() {
			NewSquare(5.0)
		})
	})
}

func TestSquareFindPerimeter(t *testing.T) {
	t.Run("check perimeter with positive length", func(t *testing.T) {
		assert.Equal(t, 6.4, NewSquare(1.6).FindPerimeter())
	})
}

func TestSquareFindArea(t *testing.T) {
	t.Run("check area with positive side", func(t *testing.T) {
		assert.Equal(t, 6.25, NewSquare(2.5).FindArea())
	})
}
