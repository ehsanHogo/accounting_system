package randgenerator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomCode(t *testing.T) {

	t.Run("test length of code is between 1 and 64", func(t *testing.T) {
		c := GenerateRandomCode()
		println(c)
		length := len(c)
		assert.GreaterOrEqual(t, length, 1, "Length should be at least 1")
		assert.LessOrEqual(t, length, 64, "Length should be at most 64")
	})

	t.Run("test rendomness of function ", func(t *testing.T) {
		c1 := GenerateRandomCode()
		c2 := GenerateRandomCode()
		assert.NotEqual(t, c1, c2, "the return value is not random")
	})

}

func TestGenerateRandomTitle(t *testing.T) {

	t.Run("test length of title is between 1 and 64", func(t *testing.T) {
		c := GenerateRandomTitle()
		println(c)
		length := len(c)
		assert.GreaterOrEqual(t, length, 1, "Length should be at least 1")
		assert.LessOrEqual(t, length, 64, "Length should be at most 64")
	})

	t.Run("test rendomness of function ", func(t *testing.T) {
		c1 := GenerateRandomTitle()
		c2 := GenerateRandomTitle()
		println(c1)
		assert.NotEqual(t, c1, c2, "the return value is not random")
	})

}
