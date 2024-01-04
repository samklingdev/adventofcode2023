package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := part1(content)
	fmt.Printf("Part 1: %d\n", result)

	result = part2(content)
	fmt.Printf("Part 2: %d\n", result)
}

func part1(content []byte) int {
	lines := strings.Split(string(content), "\n")

	var times []int
	for _, word := range strings.Fields(lines[0])[1:] { // Skip the first word ("Time:")
		num, err := strconv.Atoi(word)
		if err != nil {
			log.Fatal(err)
		}
		times = append(times, num)
	}

	var distances []int
	for _, word := range strings.Fields(lines[1])[1:] { // Skip the first word ("Distance:")
		num, err := strconv.Atoi(word)
		if err != nil {
			log.Fatal(err)
		}
		distances = append(distances, num)
	}

	multi := 1
	for i, t := range times {
		d := distances[i]
		c := 0
		for j := 1; j < t; j++ {
			if d < (j * (t - j)) {
				c++
			}
		}
		multi *= c
	}
	return multi
}

func part2(content []byte) int {
	lines := strings.Split(string(content), "\n")
	var t int
	var d int

	// remove whitespaces since the numbers divided by space should be treated as a single number
	fmt.Sscanf(strings.ReplaceAll(lines[0], " ", ""), "Time:%d", &t)
	fmt.Sscanf(strings.ReplaceAll(lines[1], " ", ""), "Distance:%d", &d)

	count := 0
	for j := 1; j < t; j++ {
		if d < (j * (t - j)) {
			count++
		}
	}
	return count
}
