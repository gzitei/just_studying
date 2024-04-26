package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
)

func delta(b, c float32) float32 {
	return (b * b) - (4 * -1 * c)
}

func root(b, delta float32) (float32, float32) {
	x1 := ((-1 * b) + float32(math.Sqrt(float64(delta)))) / (2 * -1)
	x2 := ((-1 * b) - float32(math.Sqrt(float64(delta)))) / (2 * -1)
	return x1, x2
}

func main() {
	fileName := os.Args[1]
	content, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	multipliers := []int{}
	races := make([]map[string]float32, 0)
	pieces := strings.Split(string(content), "\n")
	re := regexp.MustCompile(` {1,}`)
	times := re.Split(pieces[0], -1)
	record := re.Split(pieces[1], -1)
	for i := 1; i < len(times); i++ {
		var distance, time int
		fmt.Sscanf(record[i], "%d", &distance)
		fmt.Sscanf(times[i], "%d", &time)
		race := map[string]float32{}
		race["time"] = float32(time)
		race["record"] = float32(distance)
		races = append(races, race)
	}
	for _, rc := range races {
		delta := delta(rc["time"], -1*rc["record"])
		r1, r2 := root(rc["time"], delta)
		possibilities := math.Ceil(float64(r2-1)) - math.Floor(float64(r1+1)) + 1
		fmt.Println(rc, r1, r2, possibilities)
		multipliers = append(multipliers, int(possibilities))
	}
	fmt.Println(multipliers)
	result := 1
	for _, m := range multipliers {
		result *= m
	}
	fmt.Println("Resultado:", result)
}
