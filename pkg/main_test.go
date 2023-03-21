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
		{"1 + 2", 3},
		{"1 - 2", -1},
		{"2 * 3", 6},
		{"6 / 3", 2},
		{"1 + 2 * 3", 7},
		{"1 + 2 * 5", 11},
		{"(1 + 2) * 3", 9},
		{"(1+2) * 5", 15},
		{"(1+2) * (3+2)", 15},
		{"1 / (3 - 4)", -1},
		{"6 / (2 + 1)", 2},
		{"8 / 2 * (2 + 2)", 16},
		{"2 * (3 + 4 * 5) - 6", 40},
		{"(5 - 3) * (4 + 6 / 2)", 14},
		{"9 + 24 / (7 - 3)", 15},
		{"(3.5 + 2.5) * 4", 24},
		{"(4 - 2) * (4.5 + 1.5)", 12},
		{"(3.1 + 4.9) * 2", 16},
		{"7 / 3.5", 2},
	}

	for _, tc := range testCases {
		result, err := Calculate(tc.input)
		if err != nil {
			t.Errorf("Error while evaluating: %v", err)
		}

		if result != tc.expected {
			t.Errorf("Expected %f, got %f", tc.expected, result)
		}
		fmt.Printf("Input: %s\nExpected: %f\nResult: %f\n\n", tc.input, tc.expected, result)
	}
}
