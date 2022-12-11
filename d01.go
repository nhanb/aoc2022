package main

import (
	_ "embed"
	"sort"
	"strconv"
	"strings"
)

//go:embed d01.input
var d01 string

func MaxCalories(inp string) int {
	inp = strings.TrimSpace(inp)
	max := 0
	for _, elf := range strings.Split(inp, "\n\n") {
		sum := 0
		for _, line := range strings.Split(elf, "\n") {
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

func TopThreeCalories(inp string) int {
	inp = strings.TrimSpace(inp)
	var sums []int

	for _, elf := range strings.Split(inp, "\n\n") {
		sum := 0
		for _, line := range strings.Split(elf, "\n") {
			load, err := strconv.Atoi(line)
			check(err)
			sum += load
		}
		sums = append(sums, sum)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sums)))
	return sums[0] + sums[1] + sums[2]
}
