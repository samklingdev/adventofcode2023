package main

import (
	"testing"
)

var tests = []struct {
	input    string
	expected int
}{
	{`1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet`, 142},
	{`two1nine
	eightwothree
	abcone2threexyz
	xtwone3four
	4nineeightseven2
	zoneight234
7pqrstsixteen`, 281},
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
