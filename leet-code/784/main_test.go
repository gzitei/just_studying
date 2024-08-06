package main

import (
	"slices"
	"testing"
)

type TestCase struct {
	data     string
	solution []string
}

func TestLetterCasePermutations(t *testing.T) {
	var tests = []TestCase{
		{
			"a1b2", []string{"a1b2", "a1B2", "A1b2", "A1B2"},
		},
		{
			"3z4", []string{"3z4", "3Z4"},
		},
	}
	for i, test := range tests {
		res := letterCasePermutation(test.data)
		if len(res) != len(test.solution) {
			t.Errorf("Test %v: expected solution to contain %v elements, got %v", i, len(test.solution), len(res))
			continue
		}
		for _, sol := range test.solution {
			if !slices.Contains(res, sol) {
				t.Errorf("Test %v: %v not found in solution", i, sol)
			}
		}
	}
}
