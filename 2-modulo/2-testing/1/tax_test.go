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

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.0)
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1.0, -2, -2.5, -10.0, 0.0, 1.0, 100.0, 1000.0, 20000.0}
	for _, s := range seed {
		f.Add(s)
	}

	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Received %f, but expected 0", result)
		}
		if amount > 20000 && result != 20 {
			t.Errorf("Received %f, but expected 20", result)
		}
	})
}
