package main

import "testing"

type TestCase struct {
	data     []int
	solution int
}

func TestSingleNumber(t *testing.T) {
	tests := []TestCase{
		{
			[]int{2, 2, 1},
			1,
		},
		{
			[]int{4, 1, 2, 1, 2},
			4,
		},
		{
			[]int{1},
			1,
		},
	}
	for i, test := range tests {
		if res := singleNumber(test.data); res != test.solution {
			t.Errorf("Test %d: expected %d, got %d", i, test.solution, res)
		}
	}
}
