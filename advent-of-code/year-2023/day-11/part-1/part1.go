package main

import (
	"fmt"
	"math"
	"os"
	"slices"
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

type Galaxy map[string]int

var columns, rows []int

func (g Galaxy) CorrectPosition() Galaxy {
	newPosition := Galaxy{
		"x": g["x"],
		"y": g["y"],
	}
	for _, r := range rows {
		if r < g["y"] {
			newPosition["y"] += 1
		}
	}
	for _, c := range columns {
		if c < g["x"] {
			newPosition["x"] += 1
		}
	}
	return newPosition
}

func (g1 Galaxy) ClosestDistance(g2 Galaxy) int {
	g1 = g1.CorrectPosition()
	g2 = g2.CorrectPosition()
	hDistance := int(math.Abs(float64((g1["x"] - g2["x"]))))
	vDistance := int(math.Abs(float64((g1["y"] - g2["y"]))))
	dist := hDistance + vDistance
	return int(dist)
}

func parseText(lines []string) []Galaxy {
	galaxies := []Galaxy{}
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		count := 0
		for j := 0; j < len(line); j++ {
			if line[j] == '#' {
				count++
				galaxies = append(galaxies, Galaxy{"x": j, "y": i})
				columns = slices.DeleteFunc(columns, func(el int) bool {
					return el == j
				})
			}
		}
		if count == 0 {
			rows = append(rows, i)
		}
	}
	return galaxies
}

func Challenge(content string) int {
	sum := 0
	lines := strings.Split(content, "\n")
	for i := range lines[0] {
		columns = append(columns, i)
	}
	galaxies := parseText(lines)
	for i := 0; i < len(galaxies); i++ {
		galaxies[i].CorrectPosition()
	}
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += galaxies[i].ClosestDistance(galaxies[j])
		}
	}
	return sum
}
