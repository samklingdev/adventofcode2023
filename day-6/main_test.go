package main

import (
	"testing"
)

var tests = []struct {
	input    string
	expected int
}{
	{input: `Time:      7  15   30
Distance:  9  40  200`, expected: 288},
	{expected: 71503},
}

func assert(t *testing.T, result, expected int) {
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestMain(t *testing.T) {
	assert(t, part1([]byte(tests[0].input)), tests[0].expected)
	assert(t, part2([]byte(tests[0].input)), tests[1].expected)
}
