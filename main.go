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
		println("Max Calories:", max)
	default:
		println("Invalid week.")
		os.Exit(1)
	}
}
