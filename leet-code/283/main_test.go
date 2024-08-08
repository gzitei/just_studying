package main

import "testing"

type TestCase struct {
	data     []int
	solution []int
}

func TestMoveZeroes(t *testing.T) {
	tests := []TestCase{
		{
			[]int{0, 1, 0, 3, 12},
			[]int{1, 3, 12, 0, 0},
		},
		{
			[]int{0},
			[]int{0},
		},
		{
			[]int{0, 0, 1},
			[]int{1, 0, 0},
		},
	}
	for i, test := range tests {
		moveZeroes(test.data)
	loop:
		for j := 0; j < len(test.solution); j++ {
			if test.data[j] != test.solution[j] {
				t.Errorf("Test %d: expected %v, got %v", i, test.solution, test.data)
				break loop
			}
		}
	}
}
