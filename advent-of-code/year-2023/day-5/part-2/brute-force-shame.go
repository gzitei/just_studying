package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Seeds []int

type Range struct {
	input  int
	output int
	size   int
}

func (r Range) fits(n int) bool {
	lower := r.input
	higher := r.input + r.size
	return n >= lower && n < higher
}

func (r Range) transform(n int) int {
	result := n
	if r.fits(n) {
		diff := n - r.input
		result = r.output + diff
	}
	return result
}

func main() {
	st := time.Now()
	var closer int
	args := os.Args
	filePath := args[1]
	file, err := os.ReadFile(filePath)
	if err != nil {
		panic(err.Error())
	}
	content := string(file)
	pieces := strings.Split(content, "\n")
	var currRange []Range
	var ranges [][]Range
	seeds := make(Seeds, 0)
	for i, str := range pieces {
		if i == 0 {
			var number int
			idx := strings.Index(str, ":")
			nums := strings.Split(str[idx+2:], " ")
			for _, n := range nums {
				fmt.Sscanf(strings.TrimSpace(n), "%d", &number)
				seeds = append(seeds, number)
			}
		} else {
			var input, output, size int
			fmt.Sscanf(str, "%d %d %d", &output, &input, &size)
			if output == 0 && input == 0 && size == 0 {
				if len(currRange) > 0 {
					ranges = append(ranges, currRange)
				}
				currRange = make([]Range, 0)
			} else {
				currRange = append(currRange, Range{
					input:  input,
					output: output,
					size:   size,
				})
			}
		}
	}
	flg := -1
	for i := 0; i < len(seeds); i += 2 {
		rSeeds := make([]int, 0)
		for k := 0; k < seeds[i+1]; k++ {
			rSeeds = append(rSeeds, seeds[i]+k)
		}
		for _, seed := range rSeeds {
			curr := seed
			for i := range ranges {
				tests := ranges[i]
				for j := range tests {
					test := tests[j]
					seed = test.transform(seed)
					if seed != curr {
						curr = seed
						break
					}
				}
			}
			if flg < 0 || seed < closer {
				closer = seed
				flg++
				fmt.Println("Candidate:", closer)
			}
		}
	}
	end := time.Now()
	span := end.Sub(st).Minutes()
	fmt.Println("Closer location:", closer, "in", span, "minutes")
}
