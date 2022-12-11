package main

import (
	"os"
)

func main() {
	if len(os.Args) != 2 {
		println("Usage: go run w01")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "1":
		max := MaxCalories(d01)
		sum := TopThreeCalories(d01)
		println("Max calories:", max)
		println("Sum of top three:", sum)
	default:
		println("Invalid week.")
		os.Exit(1)
	}
}
