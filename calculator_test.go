package calculator_test

import (
	"testing"

	"calculator"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	var want float64 = 4
	got := calculator.Add(2, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	var want float64 = 2
	get := calculator.Subtract(4, 2)
	if want != get {
		t.Errorf("want %f, got %f", want, get)
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	var want float64 = 4
	get := calculator.Multiply(2, 2)
	if want != get {
		t.Errorf("want %f, got %f", want, get)
	}
}
