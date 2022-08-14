package shape

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRectangle(t *testing.T) {
	t.Run("check initialization of rectangle", func(t *testing.T) {
		assert.IsType(t, Rectangle{}, NewRectangle(1.0, 2.0))
	})

	t.Run("check initialization of rectangle with zero length", func(t *testing.T) {
		assert.Panics(t, func() {
			NewRectangle(0.0, 2.0)
		})
	})

	t.Run("check initialization of rectangle with zero width", func(t *testing.T) {
		assert.Panics(t, func() {
			NewRectangle(1.0, 0.0)
		})
	})

	t.Run("check initialization of rectangle with negative length", func(t *testing.T) {
		assert.Panics(t, func() {
			NewRectangle(-1.0, 2.0)
		})
	})

	t.Run("check initialization of rectangle with negative width", func(t *testing.T) {
		assert.Panics(t, func() {
			NewRectangle(1.0, -2.0)
		})
	})

	t.Run("check initialization of rectangle with positive sides", func(t *testing.T) {
		assert.NotPanics(t, func() {
			NewRectangle(1.0, 2.0)
		})
	})
}

func TestFindPerimeter(t *testing.T) {
	t.Run("check perimeter when length is equal to 1", func(t *testing.T) {
		assert.Equal(t, 10.0, NewRectangle(1, 4.0).FindPerimeter())
	})

	t.Run("check perimeter when width is equal to 1", func(t *testing.T) {
		assert.Equal(t, 3.0, NewRectangle(0.5, 1).FindPerimeter())
	})

	t.Run("check perimeter when length is equal to width", func(t *testing.T) {
		assert.Equal(t, 18.0, NewRectangle(7.2, 1.8).FindPerimeter())
	})

	t.Run("check perimeter when length and width are not equal", func(t *testing.T) {
		assert.Equal(t, 5.6, NewRectangle(2.0, 0.8).FindPerimeter())
	})

	t.Run("check perimeter for swapped rectangles", func(t *testing.T) {
		assert.Equal(t, NewRectangle(3.4, 4.69).FindPerimeter(), NewRectangle(4.69, 3.4).FindPerimeter())
	})

	t.Run("check if the perimeter of the rectangle is positive", func(t *testing.T) {
		assert.Less(t, 0.0, NewRectangle(6.0, 7.2).FindPerimeter())
	})
}

func TestFindArea(t *testing.T) {
	t.Run("check area when length is equal to 1", func(t *testing.T) {
		assert.Equal(t, 5.6, NewRectangle(1, 5.6).FindArea())
	})

	t.Run("check area when width is equal to 1", func(t *testing.T) {
		assert.Equal(t, 6.4, NewRectangle(6.4, 1).FindArea())
	})

	t.Run("check area when length is equal to width", func(t *testing.T) {
		assert.Equal(t, 6.25, NewRectangle(2.5, 2.5).FindArea())

	})

	t.Run("check if the area of the rectangle is positive", func(t *testing.T) {
		assert.Less(t, 0.0, NewRectangle(5.0, 4.0).FindArea())
	})

	t.Run("check area when length and width are not equal", func(t *testing.T) {
		assert.Equal(t, 12.8, NewRectangle(2.0, 6.4).FindArea())
	})

	t.Run("check area for swapped rectangles", func(t *testing.T) {
		assert.Equal(t, NewRectangle(3.4, 4.6).FindArea(), NewRectangle(4.6, 3.4).FindArea())
	})
}
