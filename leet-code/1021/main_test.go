package main

import "testing"

type TestCase struct {
	data     string
	solution string
}

func TestRemoveOutermostParentheses(t *testing.T) {
	tests := []TestCase{
		{
			"(()())(())", "()()()",
		},
		{
			"(()())(())(()(()))", "()()()()(())",
		},
		{
			"()()", "",
		},
	}
	for i, test := range tests {
		if res := removeOuterParentheses(test.data); res != test.solution {
			t.Errorf("Test %d: expected %s, got %s", i, test.solution, res)
		}
	}
}
