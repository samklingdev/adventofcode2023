package main

import (
	"testing"
)

var tests = []struct {
	input    string
	expected int
}{
	{`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
	Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
	Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
	Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`, 8},
	{`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
	Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
	Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
	Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`, 2286},
}

func assert(t *testing.T, result, expected int) {
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestMain(t *testing.T) {
	assert(t, part1([]byte(tests[0].input)), tests[0].expected)
	assert(t, part2([]byte(tests[1].input)), tests[1].expected)
}
