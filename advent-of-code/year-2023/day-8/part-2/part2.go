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
	part2(content)
	fmt.Println("Ran in", time.Since(start).Seconds(), "s")
}

func nextPoint(el string, instr map[string][2]string, direction rune, flg *int) string {
	ready := true
	var dir int
	if direction == 'L' {
		dir = 0
	} else {
		dir = 1
	}
	nextPoint := instr[el][dir]
	ready = ready && nextPoint[2] == 'Z'
	if ready {
		*flg++
	}
	return nextPoint
}

func greatestCommonDivisor(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func leastCommonMultiplier(slc []int) int {
	n := slc[0]
	for i := 1; i < len(slc); i++ {
		n = (n * slc[i]) / greatestCommonDivisor(n, slc[i])
	}
	return n
}

func part2(content string) {
	steps := 0
	content = strings.Trim(content, "\n")
	pieces := strings.Split(content, "\n\n")
	directions := pieces[0]
	mapping := strings.Split(pieces[1], "\n")
	instr := map[string][2]string{}
	currentPoints := make([]string, 0)
	for _, str := range mapping {
		var start, pointA, pointB string
		fmt.Sscanf(str, "%s = (%s%s)", &start, &pointA, &pointB)
		pointA = pointA[:3]
		pointB = strings.TrimSpace(pointB)[:3]
		instr[start] = [2]string{pointA, pointB}
		if start[2] == 'A' {
			currentPoints = append(currentPoints, start)
		}
	}
	flg := -1
	freq := make([]int, 0, len(currentPoints))
	for _, el := range currentPoints {
		currentPoint := el
	loop:
		for {
			for _, i := range directions {
				steps++
				currentPoint = nextPoint(currentPoint, instr, i, &flg)
				if flg == 0 {
					fmt.Println(steps, "steps from", el, "to", currentPoint)
					freq = append(freq, steps)
					steps = 0
					flg = -1
					break loop
				}
			}
		}
	}
	sync := leastCommonMultiplier(freq)
	fmt.Println("It takes", sync, "to reach all positions Z from all positions A")
}
