package main

/*
   Runs in 2 ms - beats 91.70%
   Uses 3.71 MB - beats 49.30%
*/

import "fmt"

type Msg struct {
	div int
	msg string
}

var rules = []Msg{
	{3, "Fizz"}, {5, "Buzz"},
}

func fizzBuzz(n int) []string {
	r := make([]string, 0, n)
	for i := 0; i < n; i++ {
		msg := ""
		for _, rule := range rules {
			if (i+1)%rule.div == 0 {
				msg += rule.msg
			}
		}
		if msg == "" {
			msg = fmt.Sprintf("%d", i+1)
		}
		r = append(r, msg)
	}
	return r
}

func main() {}
