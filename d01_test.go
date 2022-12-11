package main

import (
	_ "embed"
	"testing"
)

//go:embed d01_test.input
var d01TestInput string

func TestMaxCalories(t *testing.T) {
	ans := MaxCalories(d01TestInput)
	if ans != 24_000 {
		t.Errorf("MaxCalories() = %d; want 24,000", ans)
	}
}

func TestTopThreeCalories(t *testing.T) {
	ans := TopThreeCalories(d01TestInput)
	if ans != 45_000 {
		t.Errorf("TopThreeCalories() = %d; want 24,000", ans)
	}
}
