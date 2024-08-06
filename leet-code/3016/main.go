package main

/*
	Runs in 19ms, beats 91.67%
	Uses 6.66MB, beats 95.83%
*/

import (
	"slices"
)

func minimumPushes(word string) int {
	n := 0
	s := make([]int, 26)
	for i := 0; i < len(word); i++ {
		s[word[i]-'a'] += 1
	}

	slices.SortFunc(s, func(a, b int) int {
		return b - a
	})

	for i := 0; i < len(s); i++ {
		if s[i] == 0 {
			break
		}
		n += ((i / 8) + 1) * s[i]
	}

	return n
}

func main() {}
