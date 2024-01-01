package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	content, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := part1(content)
	fmt.Println("Result part1: ", result)

	/*
		network, directions, err = handleFile("input.txt")
		if err != nil {
			log.Fatal(err)
		}

		result = part2(network, directions)
		fmt.Println("Result part2: ", result)
	*/
}

func parse(content []byte) (map[string][2]string, string, error) {
	network := make(map[string][2]string)

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

func part1(input []byte) int {
	network, directions, err := parse(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("network:", network)
	fmt.Println("directions:", directions)
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
