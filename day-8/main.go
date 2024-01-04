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

	result := part1(content)
	fmt.Println("Result part1: ", result)

	result = part2(content)
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

	start := "A"
	end := "Z"

	startNodes := []string{}
	for k := range network {
		if strings.HasSuffix(k, start) {
			startNodes = append(startNodes, k)
		}
	}

	repeatCycles := []int{}

	// keep going until we got all repeat cycles
	for _, node := range startNodes {
		steps := 0
		for !strings.HasSuffix(node, end) {
			i := steps % len(directions)
			d := directions[i]
			if d == 'L' {
				node = network[node][0]
			} else {
				node = network[node][1]
			}
			steps++
		}
		repeatCycles = append(repeatCycles, steps)
	}
	lcm := LCM(repeatCycles[0], repeatCycles[1], repeatCycles[2:]...)
	return lcm
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

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
