package main

import (
	"io/ioutil"
	"testing"
)

func TestMaxCalories(t *testing.T) {
	input, err := ioutil.ReadFile("d01_test.input")
	check(err)
	ans := MaxCalories(string(input))
	if ans != 24_000 {
		t.Errorf("MaxCalories() = %d; want 24,000", ans)
	}
}
