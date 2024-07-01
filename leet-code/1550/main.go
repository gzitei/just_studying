package main

/*
   Runs in 2m - Beats 83.33%
   Uses 2.66MB - Beats 93.33%
*/

func threeConsecutiveOdds(arr []int) bool {
	count := 0
	for i := 0; i < len(arr) && count < 3; i++ {
		if arr[i]%2 == 1 {
			count++
		} else {
			count = 0
		}
	}
	return count == 3
}

func main() {
}

