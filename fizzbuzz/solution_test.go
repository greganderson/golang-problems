package fizzbuzz 

import (
    "fmt"
    "testing"
)

func TestFizzBuzz(t *testing.T) {
    tests := []struct {
        input    int
        expected []string
    }{
        {1, []string{"1"}},
        {3, []string{"1", "2", "fizz"}},
        {5, []string{"1", "2", "fizz", "4", "buzz"}},
        {15, []string{"1", "2", "Fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz"}},
    }

    for _, test := range tests {
        t.Run(fmt.Sprintf("n=%d", test.input), func(t *testing.T) {
            got := Fizzbuzz(test.input)
            for i, v := range got {
                if v != test.expected[i] {
                    t.Errorf("For input %d, expected %v, but got %v", test.input, test.expected, got)
                    break
                }
            }
        })
    }
}

