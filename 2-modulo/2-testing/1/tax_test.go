package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expected float64
	}

	table := []calcTax{
		{amount: 500.0, expected: 5.0},
		{amount: 1000.0, expected: 10.0},
		{amount: 1500.0, expected: 10.0},
	}

	for _, test := range table {
		result := CalculateTax(test.amount)

		if result != test.expected {
			t.Errorf("Expected %f, got %f", test.expected, result)
		}
	}
}
