package main

import (
	"io/ioutil"
	"testing"
)

func TestMaxCalories(t *testing.T) {
	input, err := ioutil.ReadFile("d01_test.input")
	check(err)
	str := string(input)

	ans := MaxCalories(str)
	if ans != 24_000 {
		t.Errorf("MaxCalories() = %d; want 24,000", ans)
	}
}

func TestTopThreeCalories(t *testing.T) {
	input, err := ioutil.ReadFile("d01_test.input")
	check(err)
	str := string(input)

	ans := TopThreeCalories(str)
	if ans != 45_000 {
		t.Errorf("TopThreeCalories() = %d; want 24,000", ans)
	}
}
