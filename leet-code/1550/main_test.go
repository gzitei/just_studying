package main

import "testing"

type TestCase struct {
	data     []int
	solution bool
}

func TestThreeConsecutiveOdds(t *testing.T) {
	tests := []TestCase{
		{
			data:     []int{2, 6, 4, 1},
			solution: false,
		},
		{
			data:     []int{1, 2, 34, 3, 4, 5, 7, 23, 12},
			solution: true,
		},
	}
	for i, test := range tests {
		if res := threeConsecutiveOdds(test.data); res != test.solution {
			t.Errorf("Test %d: expected %v got %v", i, test.solution, res)
		}
	}
}
