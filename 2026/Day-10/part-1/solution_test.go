package part1_test

import (
	"testing"

	"day-10/reader"

	part1 "day-10/part-1"
)

func TestPart1(t *testing.T) {
	expected := 0
	got := part1.Solve(reader.ReadInput("../input/test.txt"))
	if got != expected {
		t.Errorf("Test failed with expected value: %v, test output: %v", expected, got)
	}
}
