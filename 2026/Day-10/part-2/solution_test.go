package part2_test

import (
	"testing"

	"day-10/reader"

	part2 "day-10/part-2"
)

func TestPart2(t *testing.T) {
	expected := 0
	got := part2.Solve(reader.ReadInput("../input/test.txt"))
	if got != expected {
		t.Errorf("Test failed with expected value: %v, test output: %v", expected, got)
	}
}
