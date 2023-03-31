package mock

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateTaxAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil)
	repository.On("SaveTax", 0.0).Return(errors.New("amount should be greater than 0"))

	err := CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err, "error should be nil")

	err = CalculateTaxAndSave(0.0, repository)
	assert.Error(t, err, "amount should be greater than 0")

	repository.AssertExpectations(t)

}
