package main

import (
	"testing"
)

func TestPart1(t *testing.T) {

	tests := []struct {
		expected int
		input    string
	}{
		{expected: 2, input: `RL

		AAA = (BBB, CCC)
		BBB = (DDD, EEE)
		CCC = (ZZZ, GGG)
		DDD = (DDD, DDD)
		EEE = (EEE, EEE)
		GGG = (GGG, GGG)
		ZZZ = (ZZZ, ZZZ)`},
		{expected: 6, input: `LLR

		AAA = (BBB, BBB)
		BBB = (AAA, ZZZ)
		ZZZ = (ZZZ, ZZZ)`},
	}

	for _, test := range tests {
		result := part1([]byte(test.input))
		if result != test.expected {
			t.Errorf("part1(%s) returned %d, expected %d", test.input, result, test.expected)
		}
	}
}

// func TestPart2(t *testing.T) {

// 	tests := []struct {
// 		expected int
// 		input    string
// 	}{
// 		{expected: 6, input: `LR

// 		11A = (11B, XXX)
// 		11B = (XXX, 11Z)
// 		11Z = (11B, XXX)
// 		22A = (22B, XXX)
// 		22B = (22C, 22C)
// 		22C = (22Z, 22Z)
// 		22Z = (22B, 22B)
// 		XXX = (XXX, XXX)`},
// 	}

// 	for _, test := range tests {
// 		result := part2([]byte(test.input))
// 		if result != test.expected {
// 			t.Errorf("part1(%s) returned %d, expected %d", test.input, result, test.expected)
// 		}
// 	}
// }
