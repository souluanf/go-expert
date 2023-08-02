package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 0.05
	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}
