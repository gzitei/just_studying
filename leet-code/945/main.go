package main

import (
	"fmt"
	"slices"
)

// Faster than 40% (483 ms), less memory than 20% (9.88 MB)

func minIncrementForUnique(nums []int) int {
	steps := 0
	slices.Sort(nums)
	for i := 1; i < len(nums); i++ {
		for nums[i] <= nums[i-1] {
			nums[i]++
			steps++
		}
	}
	fmt.Println(nums)
	return steps
}

type TestCase struct {
	data []int
	sol  int
}

func main() {
	tests := []TestCase{
		{
			data: []int{1, 2, 2},
			sol:  1,
		},
		{
			data: []int{3, 2, 1, 2, 1, 7},
			sol:  6,
		},
	}
	failed := 0
	for i, test := range tests {
		if res := minIncrementForUnique(test.data); res != test.sol {
			failed++
			fmt.Println("Test", i, "failed: expected", test.sol, "got", res)
		}
	}
	if failed == 0 {
		fmt.Println("*** PASSED ***")
	} else {
		fmt.Println("*** FAILED ***")
	}
}
