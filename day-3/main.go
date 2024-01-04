package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type Point struct {
	x, y int
}

func main() {
	content, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := part1(content)
	fmt.Println("Result part1: ", result)

	result = part2(content)
	fmt.Println("Result part2: ", result)
}

func part1(content []byte) int {
	symbols := getSymbols(content)
	engineParts := getEngineParts(content, symbols)

	result := 0
	for _, parts := range engineParts {
		for _, part := range parts {
			result += part
		}
	}

	return result
}

func part2(content []byte) int {
	symbols := getSymbols(content)
	engineParts := getEngineParts(content, symbols)

	result := 0
	for p, adjecents := range engineParts {
		if symbols[p] == '*' && len(adjecents) == 2 {
			result += adjecents[0] * adjecents[1]
		}
	}

	return result
}

func getSymbols(content []byte) map[Point]rune {
	var symbols = map[Point]rune{}

	lines := strings.Split(string(content), "\n")
	for y, line := range lines {
		for x, r := range line {
			if r != '.' && !unicode.IsDigit(r) {
				symbols[Point{x, y}] = r
			}
		}
	}
	return symbols
}

func getEngineParts(content []byte, symbols map[Point]rune) map[Point][]int {
	var engineParts = map[Point][]int{}

	re := regexp.MustCompile(`\d+`)
	directions := []Point{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	lines := strings.Split(string(content), "\n")
	for y, line := range lines {
		matches := re.FindAllStringIndex(line, -1)
		for _, match := range matches {
			start, end := match[0], match[1]
			keys := map[Point]struct{}{}
			// engine parts must be adjecent to a symbol to be valid
			for x := start; x < end; x++ {
				for _, direction := range directions {
					keys[Point{x + direction.x, y + direction.y}] = struct{}{}
				}
			}

			number := line[start:end]
			num, err := strconv.Atoi(number)
			if err != nil {
				log.Fatal(err)
			}
			for p := range keys {
				if _, ok := symbols[p]; ok {
					// save symbol point and all neighbouring numbers as engine parts
					engineParts[p] = append(engineParts[p], num)
				}
			}

		}
	}
	return engineParts
}
