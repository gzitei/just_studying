package main

import "strings"

/*
	Runs in 5ms - beats 84.78%
	Uses 6.08MB - beats 95.65%
*/

func createWord(s string, r *[]string, i int) {
	if i >= len(s) {
		*r = append(*r, s)
		return
	}
	if 'a' <= s[i] && s[i] <= 'z' {
		createWord(s, r, i+1)
		str := s[:i] + strings.ToUpper(string(s[i])) + s[i+1:]
		createWord(str, r, i+1)
	} else if 'A' <= s[i] && s[i] <= 'Z' {
		createWord(s, r, i+1)
		str := s[:i] + strings.ToLower(string(s[i])) + s[i+1:]
		createWord(str, r, i+1)
	} else {
		createWord(s, r, i+1)
	}
}

func letterCasePermutation(s string) []string {
	res := []string{}
	createWord(s, &res, 0)
	return res
}

func main() {}
