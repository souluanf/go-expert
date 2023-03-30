package tax

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax(1000)
	assert.Nil(t, err, "error should be nil")
	assert.Equal(t, 10.0, tax, "tax should be 10.0")

	tax, err = CalculateTax(0)
	assert.Error(t, err, "amount should be greater than 0")
	assert.Equal(t, 0.0, tax, "tax should be 0.0")
}
