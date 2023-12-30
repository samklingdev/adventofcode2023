package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	times     [4]int
	distances [4]int
)

func init() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\r\n")
	fmt.Sscanf(lines[0], "Time: %d %d %d %d", &times[0], &times[1], &times[2], &times[3])
	fmt.Sscanf(lines[1], "Distance: %d %d %d %d", &distances[0], &distances[1], &distances[2], &distances[3])
}

func main() {
	// part1()
	part2()
}

func part2() {
	t := 60947882
	d := 475213810151650
	fmt.Printf("Time: %d | Distance: %d\n", t, d)
	c := 0
	for j := 1; j < t; j++ {
		if d < (j * (t - j)) {
			c++
		}
	}
	fmt.Printf("Count: %d\n", c)
}

func part1() {
	multi := 1
	for i, t := range times {
		d := distances[i]
		fmt.Printf("Time: %d | Distance: %d\n", t, d)
		c := 0
		for j := 1; j < t; j++ {
			if d < (j * (t - j)) {
				c++
			}
		}
		fmt.Printf("Count: %d\n", c)
		multi *= c
	}
	fmt.Printf("Multi: %d\n", multi)
}
