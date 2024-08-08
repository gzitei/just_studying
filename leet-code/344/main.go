package main

/*
   Runs in 25 ms - beats 89.06%
   Uses 6.66 MB - beats 71.25%
*/

func reverseString(s []byte) {
	for i := len(s); i > len(s)-i; i-- {
		s[i-1], s[len(s)-i] = s[len(s)-i], s[i-1]
	}
}

func main() {}
