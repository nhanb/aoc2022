package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed d01.input
var d01 string

func MaxCalories(inp string) int {
	inp = strings.TrimSpace(inp)
	max := 0
	for _, reindeer := range strings.Split(inp, "\n\n") {
		sum := 0
		for _, line := range strings.Split(reindeer, "\n") {
			load, err := strconv.Atoi(line)
			check(err)
			sum += load
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
