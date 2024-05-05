package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
	"time"
)

var pipeInterval = map[int]int{}

var pointsInPath = []Breadcrumb{}

type Vector struct {
	start, end Breadcrumb
}

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

func (p Breadcrumb) touchsAny(c Path) bool {
	res := false
	for _, p2 := range c {
		if p.touchs(p2) {
			res = true
			break
		}
	}
	return res
}

func (s Breadcrumb) touchs(p Breadcrumb) bool {
	xp := math.Abs(float64(s.x - p.x))
	yp := math.Abs(float64(s.y - p.y))
	return (xp+yp <= 1)
}

func (p Breadcrumb) intercepts(v Vector) bool {
	//	if p.x > pipeInterval[p.y] {
	//		return false
	//	}
	isVertical := v.start.x == v.end.x && v.start.y != v.end.y
	makesVertical := map[byte]byte{
		'L': '7',
		'7': 'L',
		'F': 'J',
		'J': 'F',
	}
	ymax := int(math.Max(float64(v.start.y), float64(v.end.y)))
	ymin := int(math.Min(float64(v.start.y), float64(v.end.y)))
	xmax := int(math.Max(float64(v.start.x), float64(v.end.x)))
	startByte := lines[v.start.y][v.start.x]
	endByte := lines[v.end.y][v.end.x]
	expectedEndByte := makesVertical[startByte]
	isLeftSide := xmax < p.x
	isCross := p.y >= ymin && p.y <= ymax
	if isVertical {
		return isLeftSide && isCross
	} else {
		return isLeftSide && p.y == ymax && endByte == expectedEndByte
	}
}

var (
	rect Vector
	path = []Vector{}
)

func Challenge(content string) int {
	paths := map[rune][]Vector{}
	var start Breadcrumb
	lines = strings.Split(content, "\n")
	start.y = slices.IndexFunc(lines, func(s string) bool {
		return strings.Contains(s, "S")
	})
	start.x = strings.Index(lines[start.y], "S")
	maxH = len(lines[0]) - 1
	maxV = len(lines) - 1
loop:
	for _, direction := range "udlr" {
		v := Vector{start: start}
		d := direction
		vdir := direction
		var current, next Breadcrumb
		current = start
		for d != 'b' {
			pointsInPath = append(pointsInPath, current)
			next, d = Walk(current, d)
			if pipeInterval[next.y] < next.x {
				pipeInterval[next.y] = next.x
			}
			v.end = next
			if next == errorPosition {
				break
			}
			if vdir != d {
				paths[direction] = append(paths[direction], v)
				vdir = d
				v.start = next
			}
			if next == start {
				paths[direction] = append(paths[direction], v)
				path = paths[direction]
				break loop
			}
			current = next
		}
	}
	response := 0
	for r := 0; r < len(lines); r++ {
		for c := 0; c < len(lines[r]); c++ {
			count := 0
			p := Breadcrumb{x: c, y: r}
			if slices.Contains(pointsInPath, p) {
				continue
			}
			for _, v := range path {
				if p.intercepts(v) {
					count++
					arrBytes := []byte(lines[r])
					arrBytes[c] = 'O'
					lines[r] = string(arrBytes)
				}
			}
			if count%2 != 0 {
				response++
				arrBytes := []byte(lines[r])
				arrBytes[c] = 'I'
				lines[r] = string(arrBytes)
			}
		}
	}
	for _, line := range lines {
		fmt.Println(line)
	}
	return response
}
