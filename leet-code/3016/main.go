package main

/*
	Runs in 178ms, beats 8.33%
	Uses 6.68MB, beats 95.83%
*/

import (
	"slices"
)

func minimumPushes(word string) (n int) {
	n = 0
	size := 0

	var strokes = map[byte]int{}
	for i := 0; i < len(word); i++ {
		if strokes[word[i]] == 0 {
			size++
		}
		strokes[word[i]] += 1
	}

	if size <= 8 {
		n = len(word)
		return
	}

	var clicks = make([]int, 0, size)
	for _, v := range strokes {
		clicks = append(clicks, v)
	}

	slices.Sort(clicks)

	for i := len(clicks) - 1; i >= 0; i-- {
		n += (((len(clicks) - 1 - i) / 8) + 1) * clicks[i]
	}
	return
}

func main() {}
