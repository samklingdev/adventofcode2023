package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type Network = map[string][2]string

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// result := part1(content)
	// fmt.Println("Result part1: ", result)

	result := part2(content)
	fmt.Println("Result part2: ", result)
}

func part1(content []byte) int {
	network, directions, err := parse(content)
	if err != nil {
		log.Fatal(err)
	}

	start := "AAA"
	end := "ZZZ"

	steps := 0
	// keep going until we hit ZZZ
	for start != end {
		i := steps % len(directions)
		d := directions[i]
		if start == end {
			break
		}
		steps++
		if d == 'L' {
			start = network[start][0]
		} else {
			start = network[start][1]
		}
	}
	return steps
}

func part2(content []byte) int {
	network, directions, err := parse(content)
	if err != nil {
		log.Fatal(err)
	}

	nodes := []string{}
	for k := range network {
		if strings.HasSuffix(k, "A") {
			nodes = append(nodes, k)
		}
	}

	steps := 0

	// keep going until all nodes end with Z
	for !allEndsWithZ(nodes) {
		i := steps % len(directions)
		d := directions[i]

		// fmt.Println("Nodes:", nodes)
		// fmt.Println("Dir:", d)
		// wait for keyboard input
		// fmt.Scanln()

		steps++
		fmt.Println("Step:", steps)
		if d == 'L' {
			for i, n := range nodes {
				nodes[i] = network[n][0]
			}
		} else {
			for i, n := range nodes {
				nodes[i] = network[n][1]
			}
		}
	}
	return steps
}

func allEndsWithZ(nodes []string) bool {
	for _, n := range nodes {
		if !strings.HasSuffix(n, "Z") {
			return false
		}
	}
	return true
}

func parse(content []byte) (Network, string, error) {
	network := make(Network)

	blocks := strings.Split(string(content), "\n\n")
	if len(blocks) != 2 {
		return nil, "", fmt.Errorf("file does not contain two blocks")
	}

	directions := blocks[0]

	for _, line := range strings.Split(blocks[1], "\n") {
		var key, left, right string
		re := regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)
		match := re.FindStringSubmatch(line)
		if len(match) != 4 {
			return nil, "", fmt.Errorf("invalid line: %s", line)
		}
		key, left, right = match[1], match[2], match[3]
		network[key] = [2]string{left, right}
	}

	return network, directions, nil
}
