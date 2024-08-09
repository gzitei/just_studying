package main

import "math"

func titleToNumber(columnTitle string) int {
	res := 0
	for i := 0; i < len(columnTitle); i++ {
		res += (int(columnTitle[i]) - 'A' + 1) * int(math.Pow(26, float64(len(columnTitle)-i-1)))
	}
	return res
}

func main() {}
