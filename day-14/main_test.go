package main

import (
	"testing"
)

var tests = []struct {
	input    string
	expected int
}{
	{input: ``, expected: 0},
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
