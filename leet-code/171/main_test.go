package main

import "testing"

type TestCase struct {
	data     string
	solution int
}

func TestTitleToNumber(t *testing.T) {
	tests := []TestCase{
		{
			"A", 1,
		},
		{
			"AB", 28,
		},
		{
			"ZY", 701,
		},
		{
			"FXSHRXW",
			2147483647,
		},
	}
	for i, test := range tests {
		if res := titleToNumber(test.data); res != test.solution {
			t.Errorf("Test %d: expected %d, got %d", i, test.solution, res)
		}
	}
}
