package main

import "testing"

type TestCase struct {
	data     []string
	solution int
}

func TestMinOperations(t *testing.T) {
	tests := []TestCase{
		{
			[]string{"d1/", "d2/", "../", "d21/", "./"},
			2,
		}, {
			[]string{"d1/", "d2/", "./", "d3/", "../", "d31/"},
			3,
		}, {
			[]string{"d1/", "../", "../", "../"},
			0,
		},
	}
	for i, test := range tests {
		if res := minOperations(test.data); res != test.solution {
			t.Errorf("Test %d: expected %d got %d", i, test.solution, res)
		}
	}
}
