package main

/*
   Runs in 1ms - Beats 77.67%
   Uses 2.65MB - Beats 98.06%
*/

func removeOuterParentheses(s string) string {
	opened := 0
	res := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			if opened >= 1 {
				res = append(res, s[i])
			}
			opened++
		case ')':
			if opened > 1 {
				res = append(res, s[i])
			}
			opened--
		}
	}
	return string(res)
}

func main() {
}

