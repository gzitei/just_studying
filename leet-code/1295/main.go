package main

import "math"

func findNumbers(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		j := 1.0
		for nums[i] != (nums[i] % int(math.Pow(10.0, j))) {
			j++
		}
		if int(j)%2 == 0 {
			count++
		}
	}
	return count
}

func main() {}
