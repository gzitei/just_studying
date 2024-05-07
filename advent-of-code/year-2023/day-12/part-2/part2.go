package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
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

type Record struct {
	text    string
	changes []int
	req     []int
}

type Map []Record

func parseData(line string) Record {
	split := strings.Split(line, " ")
	nums := strings.Split(split[1], ",")
	req := make([]int, 0)
	changes := make([]int, 0)
	for i := 0; i < len(split[0]); i++ {
		if split[0][i] == '?' {
			changes = append(changes, i)
		}
	}
	for _, n := range nums {
		var num int
		fmt.Sscanf(n, "%d", &num)
		req = append(req, num)
	}
	r := Record{
		text:    "",
		changes: []int{},
		req:     []int{},
	}
	for range 5 {
		r.text += split[0]
		r.changes = slices.Concat(r.changes, changes)
		r.req = slices.Concat(r.req, req)
	}
	return r
}

func (r Record) combinations() int {
	count := 0
	arr := []byte{'.', '#'}
	str := r.text
	changes := r.changes
	req := r.req
	n := len(changes)
	arrBytes := []byte(str)
	for i := 0; i < int(math.Pow(2, float64(n))); i++ {
		itWorks := true
		num := fmt.Sprintf("%0"+fmt.Sprintf("%d", n)+"b\n", i)
		for j := 0; j < len(changes); j++ {
			idx := changes[j]
			arrBytes[idx] = arr[int(num[j])-48]
		}
		s := string(arrBytes)
		exp := ""
		for j := 0; j < len(req); j++ {
			if j == 0 {
				exp += `^\.*#{` + fmt.Sprintf("%d", req[j]) + `}\.{1,}`
			} else if j+1 == len(req) {
				exp += `#{` + fmt.Sprintf("%d", req[j]) + `}\.*$`
			} else {
				exp += `#{` + fmt.Sprintf("%d", req[j]) + `}\.{1,}`
			}
		}
		re := regexp.MustCompile(exp)
		itWorks = re.MatchString(s) && itWorks
		if itWorks {
			count++
		}
	}
	return count
}

func Challenge(content string) int {
	sum := 0
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		record := parseData(line)
		sum += record.combinations()
	}
	return sum
}
