package main

import (
	"testing"
)

var tests = []struct {
	input    string
	expected int
}{
	{input: `.....
.S-7.
.|.|.
.L-J.
.....`, expected: 4},
	{input: `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`, expected: 8},
}

func assert(t *testing.T, result, expected int) {
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestMain(t *testing.T) {
	assert(t, part1([]byte(tests[0].input)), tests[0].expected)
	// assert(t, part2([]byte(tests[1].input)), tests[1].expected)
}
