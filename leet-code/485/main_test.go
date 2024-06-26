package main

import "testing"

type TestCase struct {
	data     []int
	solution int
}

func TestFindMaxConsecutiveOnes(t *testing.T) {
	tests := []TestCase{
		{
			data:     []int{1, 1, 0, 1, 1, 1},
			solution: 3,
		},
		{
			data:     []int{1, 0, 1, 1, 0, 1},
			solution: 2,
		},
		{
			data:     []int{0},
			solution: 0,
		},
		{
			data:     []int{1, 1, 0, 1},
			solution: 2,
		},
	}
	for i, test := range tests {
		if res := findMaxConsecutiveOnes(test.data); res != test.solution {
			t.Errorf("Test %d: got %d, expected %d", i, res, test.solution)
		}
	}
}
