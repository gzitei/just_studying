package main

// Faster than 76.37% (5 ms) and less memory than 60.75%

func lengthOfLongestSubstring(s string) int {
	maxL := 0
	if len(s) < 2 {
		maxL = len(s)
	} else {
		mp := map[byte]int{}
		startPos := 0
		for i := 0; i < len(s); i++ {
			v, found := mp[s[i]]
			if v >= startPos && found {
				startPos = v + 1
			}
			mp[s[i]] = i
			size := i + 1 - startPos
			if size > maxL {
				maxL = size
			}
		}
	}
	return maxL
}
