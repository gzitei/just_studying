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
	result := Challenge(content)
	fmt.Println("Result:", result)
	fmt.Println("Ran in", time.Since(start).Seconds(), "s")
}

type Record struct {
	row       string
	positions []int
}

type Map []Record

func parseData(line string) Record {
	split := strings.Split(line, " ")
	nums := strings.Split(split[1], ",")
	numbers := make([]int, len(nums))
	for i, n := range nums {
		fmt.Sscanf(n, "%d", &numbers[i])
	}
	return Record{
		row:       split[0],
		positions: numbers,
	}
}

func Challenge(content string) int {
	sum := 0
	lines := strings.Split(content, "\n")
	springMap := make(Map, len(lines))
	for i, line := range lines {
		springMap[i] = parseData(line)
	}
	fmt.Println(springMap)
	return sum
}
