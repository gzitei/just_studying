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
		res := 0
		current := arr
	loop:
		for {
			test := make([]int, 0)
			row := make([]int, 0, len(current)-1)
			res += current[len(current)-1]
			fmt.Print(current, " => ", res, " ")
			for i := 1; i < len(current); i++ {
				n := current[i] - current[i-1]
				row = append(row, n)
				if n != 0 {
					test = append(test, n)
				}
			}
			if len(test) == 0 {
				fmt.Println(" =>", res)
				break loop
			}
			current = row
		}
		sum += res
	}
	return sum
}
