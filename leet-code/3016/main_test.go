package main

import "testing"

type TestCase struct {
	data     string
	solution int
}

func TestMinimumPushes(t *testing.T) {
	var tests = []TestCase{
		{
			"abcde",
			5,
		},
		{

			"xyzxyzxyzxyz",
			12,
		},
		{
			"aabbccddeeffgghhiiiiii",
			24,
		},
	}
	for i, test := range tests {
		if res := minimumPushes(test.data); res != test.solution {
			t.Errorf("Test %d: expected %d got %d", i, test.solution, res)
		}
	}
}
