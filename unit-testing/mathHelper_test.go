package main

import (
	"fmt"
	"testing"
)

// go test -v
func TestSum(t *testing.T) {
	res := Sum(1, 2)

	if res != 3 {
		t.Errorf("Sum(1, 2) expected %v, got %v", 3, res)
	}
}

func TestSubtract(t *testing.T) {
	/*
		use t.Run to name the test
		we declare an inner t *testing.T to avoid pollution - even though this
		is a singular test and there would probably be no pollution
	*/
	t.Run("Subtract 3 from 5", func(t *testing.T) {
		res := Subtract(5, 3)
		if res != 2 {
			t.Errorf("Subtract(5, 3) expected %v, got %v", 2, res)
		}
	})
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{1, 1, 1},
		{2, 2, 4},
		{3, 1, 3},
		{6, 8, 48},
		{150, 10, 1500},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("Multiply Test: %v * %v want %v", test.a, test.b, test.want)
		t.Run(testName, func(t *testing.T) {
			result := Multiply(test.a, test.b)

			if result != test.want {
				t.Errorf("%v * %v should be %v, got %v instead", test.a, test.b, test.want, result)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{9, 3, 3},
		{15, 3, 5},
		{7, 2, 3},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("Divide Test: %v * %v want %v", test.a, test.b, test.want)
		t.Run(testName, func(t *testing.T) {
			result := Divide(test.a, test.b)

			if result != test.want {
				t.Errorf("%v / %v should be %v, got %v instead", test.a, test.b, test.want, result)
			}
		})
	}
}

//go test -v -bench=.
func BenchmarkSum(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Sum(2, 3)
	}
}
