package calculator_test

import (
	"testing"

	"calculator"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64
		want float64
	}

	testcases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 2, b: 3, want: 5},
		{a: 4, b: 3, want: 7},
	}

	for _, tc := range testcases {
		got := calculator.Add(tc.a, tc.b)
		if got != tc.want {
			t.Errorf("want %f, got %f", tc.want, got)
		}
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
