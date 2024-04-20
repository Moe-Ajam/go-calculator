package calculator_test

import (
	"math"
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

	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{4, 2, 2},
		{5, 2, 3},
		{2, 0, 2},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if got != tc.want {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{4, 2, 8},
		{2, 2, 4},
		{3, 2, 6},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if got != tc.want {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{4, 2, 2},
		{5, 2, 2.5},
		{10, 2, 5},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			t.Fatalf("want no error for valid input, got %v", err)
		}
		if !closeEnough(got, tc.want, 0.001) {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestDivideInvalid(t *testing.T) {
	t.Parallel()

	_, err := calculator.Divide(1, 0)

	if err == nil {
		t.Errorf("want error for invalid input got nil")
	}
}

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestSqrt(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a    float64
		want float64
	}

	testCases := []testCase{
		{4, 2},
		{9, 3},
	}

	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)

		if err != nil {
			t.Fatalf("want no erro for valid input, got %v", err)
		}

		if !closeEnough(got, tc.want, 0.1) {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestSqrtInvalid(t *testing.T) {
	t.Parallel()

	_, err := calculator.Sqrt(-1)

	if err == nil {
		t.Errorf("want error on invalid input, got nil")
	}
}
