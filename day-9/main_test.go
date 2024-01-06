package main

import (
	"testing"
)

var tests = []struct {
	input    string
	expected int
}{
	{input: `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`, expected: 114},
	{expected: 2},
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
