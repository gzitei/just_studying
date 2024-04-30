package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
	"time"
)

type Breadcrumb struct {
	x int
	y int
}

type Path []Breadcrumb

type Pipe map[rune]rune

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

var pipes map[rune]Pipe = map[rune]Pipe{
	'|': {
		'u': 'u',
		'd': 'd',
		'l': 'b',
		'r': 'b',
	},
	'-': {
		'u': 'b',
		'd': 'b',
		'l': 'l',
		'r': 'r',
	},
	'L': {
		'u': 'b',
		'd': 'r',
		'l': 'u',
		'r': 'b',
	},
	'J': {
		'u': 'b',
		'd': 'l',
		'l': 'b',
		'r': 'u',
	},
	'7': {
		'u': 'l',
		'd': 'b',
		'l': 'b',
		'r': 'd',
	},
	'F': {
		'u': 'r',
		'd': 'b',
		'l': 'd',
		'r': 'b',
	},
}

var errorPosition Breadcrumb = Breadcrumb{
	x: -1,
	y: -1,
}

func Walk(position Breadcrumb, direction rune) (Breadcrumb, rune) {
	var newPosition Breadcrumb
	switch direction {
	case 'u':
		{
			if position.y-1 < 0 {
				return errorPosition, 'b'
			}
			newPosition.y = position.y - 1
			newPosition.x = position.x
		}
	case 'd':
		{
			if position.y+1 > maxV {
				return errorPosition, 'b'
			}
			newPosition.y = position.y + 1
			newPosition.x = position.x
		}
	case 'l':
		{
			if position.x-1 < 0 {
				return errorPosition, 'b'
			}
			newPosition.y = position.y
			newPosition.x = position.x - 1
		}
	case 'r':
		{
			if position.x+1 > maxH {
				return errorPosition, 'b'
			}
			newPosition.y = position.y
			newPosition.x = position.x + 1
		}
	default:
		{
			return errorPosition, 'b'
		}
	}
	newLine := lines[newPosition.y]
	newEl := newLine[newPosition.x]
	newPipe, ok := pipes[rune(newEl)]
	if !ok && newEl != 'S' {
		return errorPosition, 'b'
	}
	return newPosition, newPipe[direction]
}

var (
	maxH, maxV int
	lines      []string
)

func (s Breadcrumb) distance(p Breadcrumb) int {
	vertical := math.Abs(float64(p.y - s.y))
	horizontal := math.Abs(float64(p.x - s.x))
	return int(vertical) + int(horizontal)
}

func Challenge(content string) int {
	paths := map[rune]Path{}
	var start Breadcrumb
	distance := -1
	lines = strings.Split(content, "\n")
	start.y = slices.IndexFunc(lines, func(s string) bool {
		return strings.Contains(s, "S")
	})
	start.x = strings.Index(lines[start.y], "S")
	maxH = len(lines[0]) - 1
	maxV = len(lines) - 1
	for _, direction := range "udlr" {
		paths[direction] = Path{}
		dist := distance
		d := direction
		var current, next Breadcrumb
		current = start
		for d != 'b' {
			next, d = Walk(current, d)
			if next == errorPosition {
				break
			}
			newD := start.distance(next)
			paths[direction] = append(paths[direction], next)
			if dist == -1 || newD > dist {
				dist = newD
			}
			if next == start {
				if dist > distance {
					distance = dist
				}
				break
			}
			current = next
		}
	}
	greaterDistance := -1
	for _, v := range paths {
		idx := int(len(v) / 2)
		if idx > greaterDistance {
			greaterDistance = idx
		}
	}
	return greaterDistance
}
