package main

import (
	"testing"
)

type TestCase struct {
	board    Board
	solution Board
}

func TestRemoveIslands(t *testing.T) {
	test := TestCase{
		board: Board{
			{1, 0, 0, 0, 0, 0},
			{0, 1, 0, 1, 1, 1},
			{0, 0, 1, 0, 1, 0},
			{1, 1, 0, 0, 1, 0},
			{1, 0, 1, 1, 0, 0},
			{1, 0, 0, 0, 0, 1},
		},
		solution: Board{
			{1, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 1},
			{0, 0, 0, 0, 1, 0},
			{1, 1, 0, 0, 1, 0},
			{1, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 1},
		},
	}
	solution := test.solution
	res := removeIslands(test.board)
	success := len(res) == len(solution)
	if !success {
		t.Errorf("Expected board len to be %d got %d", len(solution), len(res))
		t.FailNow()
	}
	for i := 0; i < len(solution) && success; i++ {
		success = success && len(solution[i]) == len(res[i])
		if !success {
			t.Errorf("Expected board's %d row len to be %d got %d", i, len(solution[i]), len(res[i]))
			t.FailNow()
		}
		for j := 0; j < len(solution[i]) && success; j++ {
			success = success && solution[i][j] == res[i][j]
			if !success {
				t.Errorf("Element at %d x %d: expected %d, got %d", i, j, solution[i][j], res[i][j])
			}

		}
	}
}
