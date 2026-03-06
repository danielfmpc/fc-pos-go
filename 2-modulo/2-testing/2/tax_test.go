package tax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	result, err := CalculateTax(1000.5)

	assert.Nil(t, err)
	assert.Equal(t, 10.0, result)

	result, err = CalculateTax(0)
	assert.Error(t, err, "amount must be greater than 0")
	assert.Equal(t, 0.0, result)
	assert.Contains(t, err.Error(), "be greater than 0")
}

// func TestCalculateTaxBatch(t *testing.T) {
// 	type calcTax struct {
// 		amount, expected float64
// 	}

// 	table := []calcTax{
// 		{amount: 500.0, expected: 5.0},
// 		{amount: 1000.0, expected: 10.0},
// 		{amount: 1500.0, expected: 10.0},
// 	}

// 	for _, test := range table {
// 		result, err := CalculateTax(test.amount)

// 		if result != test.expected {
// 			t.Errorf("Expected %f, got %f", test.expected, result)
// 		}
// 	}
// }
