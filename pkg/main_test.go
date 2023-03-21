package main

import (
	"fmt"
	"testing"
)

func TestEvaluateExpression(t *testing.T) {
	testCases := []struct {
		input    string
		expected float64
	}{
		{"5 + 3", 8},
		{"1 + 2 * 5", 11},
		{"(1+2) * 5", 15},
		{"(1+2) * (3+2)", 15},
		{"1 / (3 - 4)", -1},
	}

	for _, tc := range testCases {
		result, err := Calculate(tc.input)
		if err != nil {
			t.Errorf("Error while evaluating: %v", err)
		}

		if result != tc.expected {
			t.Errorf("Expected %f, got %f", tc.expected, result)
		}
		fmt.Printf("Input: %s = Expected: %f => Result: %f\n", tc.input, tc.expected, result)
	}
}
