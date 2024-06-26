package main

/*
   Runs in 3ms - Beats 87.38%
   Uses 5.77MB - Beats 5.51%
*/

func removeDuplicates(nums []int) int {
	pos := 0
	dic := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if _, ok := dic[nums[i]]; !ok {
			dic[nums[i]] = 1
			if nums[pos] != nums[i] {
				nums[pos] = nums[i]
			}
			pos++
		}
	}
	return pos
}

func main() {
}
