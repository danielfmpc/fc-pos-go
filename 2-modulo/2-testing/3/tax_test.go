package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TaxRepositoryMock struct {
	mock.Mock
}

func (m *TaxRepositoryMock) SaveTax(amount float64) error {
	args := m.Called(amount)

	return args.Error(0)
}

func TestCalculateTax(t *testing.T) {
	result, err := CalculateTax(1000.5)

	assert.Nil(t, err)
	assert.Equal(t, 10.0, result)

	result, err = CalculateTax(0)
	assert.Error(t, err, "amount must be greater than 0")
	assert.Equal(t, 0.0, result)
	assert.Contains(t, err.Error(), "be greater than 0")
}

func TestCalculateTaxAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}

	repository.On("SaveTax", 10.0).Return(nil)
	repository.On("SaveTax", 0.0).Return(errors.New("error saving tax"))

	err := CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0.0, repository)
	assert.Error(t, err, "error saving tax")

	repository.AssertExpectations(t)

	repository.AssertNumberOfCalls(t, "SaveTax", 2)
}
