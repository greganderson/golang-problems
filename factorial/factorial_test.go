package main

import (
	"fmt"
	"testing"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		input	 int
		expected int 
	}{
		{ 5, 120 },
		{ 4, 24 },
		{ 6, 720 },
		{ 8, 40320 },
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("n = %d", test.input), func(t *testing.T) {
			got := factorial(test.input)
			if got != test.expected {
				t.Errorf("For input %d, expected %d but got %d", test.input, test.expected, got)
			}
		})
	}
}
