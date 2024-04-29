package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func readFile(fileName string) []byte {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return bytes
}

func main() {
	start := time.Now()
	args := os.Args
	fileName := args[1]
	content := strings.Trim(string(readFile(fileName)), "\n")
	result := part1(content)
	fmt.Println("Result:", result)
	fmt.Println("Ran in", time.Since(start).Seconds(), "s")
}

func part1(content string) int {
	sum := 0
	pieces := strings.Split(content, "\n")
	data := make([][]int, 0, len(pieces))
	for _, piece := range pieces {
		splt := strings.Split(strings.TrimSpace(piece), " ")
		arr := make([]int, len(splt))
		for i, v := range splt {
			fmt.Sscanf(v, "%d", &arr[i])
		}
		data = append(data, arr)
	}
	for _, arr := range data {
		res := make([]int, 0)
		it := make([][]int, 0)
		current := arr
	loop:
		for {
			test := make([]int, 0)
			row := make([]int, 0, len(current)-1)
			it = append(it, current)
			for i := 1; i < len(current); i++ {
				n := current[i] - current[i-1]
				row = append(row, n)
				if n != 0 {
					test = append(test, n)
				}
			}
			if len(test) == 0 {
				break loop
			}
			current = row
		}
		base := 0
		for i := len(it) - 1; i > -1; i-- {
			curr := it[i]
			x := curr[0] - base
			res = append(res, x)
			base = x
		}
		fmt.Println(arr, "=>", res)
		sum += res[len(res)-1]
	}
	return sum
}
