package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
	"time"
)

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

func CreateClusters(v Vector) []Breadcrumb {
	res := []Breadcrumb{}
	for i := v.start.y; i <= v.end.y; i++ {
		line := lines[i]
		if !strings.Contains(line, ".") {
			continue
		}
		for j := v.start.x; j <= v.end.x; j++ {
			if line[j] == '.' {
				p := Breadcrumb{x: j, y: i}
				res = append(res, p)
			}
		}
	}
	return res
}

func (p Breadcrumb) intercepts(v Vector) bool {
	coord := []int{v.start.y, v.end.y}
	slices.Sort(coord)
	return p.y >= coord[0] && p.y <= coord[1] && v.start.x-v.end.x >= 0
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
	var minX, maxX, minY, maxY int
	maxX, maxY, minY, maxY = -1, -1, -1, -1
loop:
	for _, direction := range "udlr" {
		v := Vector{start: start}
		d := direction
		vdir := direction
		var current, next Breadcrumb
		current = start
		for d != 'b' {
			next, d = Walk(current, d)
			if maxX == -1 {
				maxX = next.x
				minX = next.x
				maxY = next.y
				minY = next.y
			}
			if next.y > maxY {
				maxY = next.y
			}
			if next.y < minY {
				minY = next.y
			}
			if next.x > maxX {
				maxX = next.x
			}
			if next.x < minX {
				minX = next.x
			}
			v.end = next
			if next == errorPosition {
				break
			}
			if vdir != d {
				paths[direction] = append(paths[direction], v)
				if v.start.x == v.end.x && v.start.y != v.end.y {
					path = append(path, v)
				}
				vdir = d
				v.start = next
			}
			if next == start {
				paths[direction] = append(paths[direction], v)
				if v.start.x == v.end.x && v.start.y != v.end.y {
					path = append(path, v)
				}
				break loop
			}
			current = next
		}
	}

	rect.start.x = minX
	rect.start.y = minY
	rect.end.x = maxX
	rect.end.y = maxY
	response := 0
	clusters := CreateClusters(rect)
	for _, p := range clusters {
		inside := 0
		for _, vecs := range paths {
			intVects := 0
			for _, v := range vecs {
				if v.start.x < p.x && p.intercepts(v) {
					fmt.Println(p, "intercepts", v)
					intVects++
				}
			}
			if intVects%2 == 1 {
				inside++
			}
		}
		if inside%2 == 1 {
			fmt.Println(p)
			response++
		}
	}
	return response
}
