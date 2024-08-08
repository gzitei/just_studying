package main

/*
   Runs in 13 ms - beats 68.71%
   Uses 6.00 MB - beats 99.37%
*/

import "slices"

func singleNumber(nums []int) int {
	slices.Sort(nums)
	r := nums[len(nums)-1]
	for i := 0; i < len(nums)-1; i += 2 {
		if nums[i] != nums[i+1] {
			r = nums[i]
			break
		}
	}
	return r
}

func main() {}
