package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
)

func delta(b, c float64) float64 {
	return (b * b) - (4 * -1 * c)
}

func root(b, delta float64) (float64, float64) {
	x1 := ((-1 * b) + float64(math.Sqrt(delta))) / (2 * -1)
	x2 := ((-1 * b) - float64(math.Sqrt(delta))) / (2 * -1)
	return x1, x2
}

func main() {
	fileName := os.Args[1]
	content, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	multipliers := []int64{}
	races := make([]map[string]float64, 0)
	pieces := strings.Split(string(content), "\n")
	re := regexp.MustCompile(` {1,}`)
	times := strings.Split(re.ReplaceAllLiteralString(pieces[0], ""), ":")
	record := strings.Split(re.ReplaceAllLiteralString(pieces[1], ""), ":")
	for i := 1; i < len(times); i++ {
		var distance, time int64
		fmt.Sscanf(record[i], "%d", &distance)
		fmt.Sscanf(times[i], "%d", &time)
		race := map[string]float64{}
		race["time"] = float64(time)
		race["record"] = float64(distance)
		races = append(races, race)
	}
	fmt.Println(races)
	for _, rc := range races {
		delta := delta(rc["time"], -1*rc["record"])
		r1, r2 := root(rc["time"], delta)
		possibilities := math.Ceil(float64(r2-1)) - math.Floor(float64(r1+1)) + 1
		fmt.Println(rc, r1, r2, possibilities)
		multipliers = append(multipliers, int64(possibilities))
	}
	fmt.Println(multipliers)
	var result int64
	result = 1
	for _, m := range multipliers {
		result *= m
	}
	fmt.Println("Resultado:", result)
}
