package main

import "testing"

type TestCase struct {
	data     []int
	solution int
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []TestCase{
		{
			data:     []int{1, 1, 2},
			solution: 2,
		},
		{
			data:     []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			solution: 5,
		},
	}
	for i, test := range tests {
		if res := removeDuplicates(test.data); res != test.solution {
			t.Errorf("Test %d: got %d expected %d", i, res, test.solution)
		}
	}
}
