package main

import (
	"fmt"
	"slices"
)

// faster than 77.53% (3 ms) and lest memory than 80.18% (3.5 MB)
func singleNumber(nums []int) int {
	// make sure nums are ordered
	slices.Sort(nums)
	// there are 3n + 1, so the last element won't be checked
	// we are covered if element is not found in the for loop
	r := nums[len(nums)-1]
loop:
	// traverse the slice by groups of 3
	for i := 0; i < len(nums)-2; i += 3 {
		a := nums[i]
		b := nums[i+1]
		c := nums[i+2]
		// if one of the group's element is differt, we found our candidate
		switch {
		case a == b && c != a:
			{
				r = c
				break loop
			}
		case a == c && b != a:
			{
				r = b
				break loop
			}
		case b == c && a != b:
			{
				r = a
				break loop
			}
		}
	}
	return r
}

type test []int

func main() {
	t1 := []int{2, 2, 3, 2}
	t2 := []int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}
	testes := []test{t1, t2}
	for _, t := range testes {
		r := singleNumber(t)
		fmt.Println(t, "=>", r)
	}
}
