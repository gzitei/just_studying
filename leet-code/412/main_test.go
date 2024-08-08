package main

import (
	"slices"
	"testing"
)

type TestCase struct {
	data     int
	solution []string
}

func TestFizzBuzz(t *testing.T) {
	tests := []TestCase{
		{
			3,
			[]string{"1", "2", "Fizz"},
		},
		{
			5,
			[]string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			15,
			[]string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}
	for i, test := range tests {
		res := fizzBuzz(test.data)
		if len(res) != len(test.solution) {
			t.Errorf("Test %d: Expected solution to have %d elements, got %d", i, len(test.solution), len(res))
		}

		for _, s := range test.solution {
			if !slices.Contains(res, s) {
				t.Errorf("Test %d: Expected solution to contain %s, got %v", i, s, res)
			}
		}
	}
}
