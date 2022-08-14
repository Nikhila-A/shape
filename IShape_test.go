package shape

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShapes(t *testing.T) {
	t.Run("check if the rectangle implements from IShape", func(t *testing.T) {
		assert.Implements(t, new(IShape), NewRectangle(10.0, 2.0))
	})

	//t.Run("check if the square implements from IShape", func(t *testing.T) {
	//	assert.Implements(t, new(IShape), NewSquare(10.0))
	//})
}
