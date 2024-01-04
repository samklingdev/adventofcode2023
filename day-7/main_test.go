package main

import (
	"testing"
)

var tests = []struct {
	input    string
	expected int
}{
	{input: `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`, expected: 6440},
	{expected: 5905},
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
