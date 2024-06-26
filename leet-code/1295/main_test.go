package main

import "testing"

type TestCase struct {
	data     []int
	solution int
}

func TestFindNumbers(t *testing.T) {
	tests := []TestCase{
		{
			data:     []int{12, 345, 2, 6, 7896},
			solution: 2,
		},
		{
			data:     []int{555, 901, 482, 1771},
			solution: 1,
		},
	}
	for i, test := range tests {
		if res := findNumbers(test.data); res != test.solution {
			t.Errorf("Test %d: got %d expected %d", i, res, test.solution)
		}
	}
}
