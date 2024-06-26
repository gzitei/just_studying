package main

/*
   Runs in 21ms - Beats 97.99%
   Uses 6.70MB - Beats 85.40%
*/

func findMaxConsecutiveOnes(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	start, end, max := -1, -1, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 && start == -1 {
			start = i
		}
		if nums[i] == 1 && start != -1 {
			end = i
		}
		if nums[i] != 1 || i == len(nums)-1 {
			d := 0
			if start >= 0 && end >= 0 {
				d = end - start + 1
			}
			if d > max {
				max = d
			}
			start, end = -1, -1
		}
	}
	return max
}

func main() {}
