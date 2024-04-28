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
	content := string(readFile(fileName))
	part1(content)
	fmt.Println("Ran in", time.Since(start).Seconds(), "s")
}

func part1(content string) {
	steps := 0
	content = strings.Trim(content, "\n")
	pieces := strings.Split(content, "\n\n")
	directions := pieces[0]
	mapping := strings.Split(pieces[1], "\n")
	instr := map[string][2]string{}
	for _, str := range mapping {
		var start, pointA, pointB string
		fmt.Sscanf(str, "%s = (%s%s)", &start, &pointA, &pointB)
		pointA = pointA[:3]
		pointB = strings.TrimSpace(pointB)[:3]
		fmt.Println(str, pointA, "|", pointB)
		instr[start] = [2]string{pointA, pointB}
	}
	currentPoint := "AAA"
	fmt.Println(instr)
loop:
	for {
		for _, i := range directions {
			var ok bool
			steps++
			fmt.Print(steps, " ", currentPoint, " ", "=> ")
			destination, ok := instr[currentPoint]
			if i == 'L' {
				currentPoint = destination[0]
			} else {
				currentPoint = destination[1]
			}
			fmt.Println(currentPoint)
			if currentPoint == "ZZZ" || !ok {
				break loop
			}
		}
	}
	fmt.Println("It took", steps, "to reach ZZZ from AAA")
}
