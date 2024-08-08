package main

import (
	"testing"
)

type TestCase struct {
	data     []byte
	solution string
}

func TestReverseString(t *testing.T) {
	tests := []TestCase{
		{[]byte("hello"), "olleh"}, {[]byte("Hannah"), "hannaH"}, {[]byte("viola"), "aloiv"},
	}
	for i, test := range tests {
		if reverseString(test.data); string(test.data) != test.solution {
			t.Errorf("Test %d: expected %s, got %s", i, test.solution, string(test.data))
		}
	}
}
