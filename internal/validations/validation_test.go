package validations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckEmpty(t *testing.T) {
	t.Run("return error when find empty field", func(t *testing.T) {
		s := ""
		err := CheckEmpty(s)
		assert.Error(t, err, "Expected error indicate empty field ")
	})

	t.Run("when field is non empty , successfully return nil", func(t *testing.T) {
		s := "test"
		err := CheckEmpty(s)
		assert.NoError(t, err, "Expected no error for non empty field")
	})
}

func TestCheckMaxLength(t *testing.T) {
	t.Run("return error when length is greater than 64", func(t *testing.T) {
		s := "a9b8c7d6e5f4g3h2i1j0k9l8m7n6o5p4q3r2s1t0u9v8w7x6y5z4a3b2c1d012345"
		err := CheckMaxLength(s, 64)
		assert.Error(t, err, "Expected error indicate length greater than max length")
	})

	t.Run("when field length is less than 64 , successfully return nil", func(t *testing.T) {
		s := "test"
		maxl := 64
		err := CheckMaxLength(s, maxl)
		assert.NoError(t, err, "Expected no error for length less than max length %d", maxl)
	})

	t.Run("when field length is equal 64 , successfully return nil", func(t *testing.T) {
		s := "a9b8c7d6e5f4g3h2i1j0k9l8m7n6o5p4q3r2s1t0u9v8w7x6y5z4a3b2c1d01234"
		maxl := 64
		err := CheckMaxLength(s, 64)
		assert.NoError(t, err, "Expected no error for length less than max length %d", maxl)
	})
}
