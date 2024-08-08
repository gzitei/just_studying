package main

/*
   Runs in 58 ms - beats 9.05%
   Uses 6.56 MB - beats 96.56%
*/

func moveZeroes(nums []int) {
	for i := len(nums); i > 0; i-- {
		if nums[i-1] == 0 {
			for j := i; j < len(nums); j++ {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}
}

func main() {}
