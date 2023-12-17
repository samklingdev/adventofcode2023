package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Number struct {
	Point
	value string
}

type Special struct {
	Point
	value           rune
	adjacentNumbers []Number
}

var nums = map[Point]Number{}
var specials = map[Point]Special{}

func main() {
	generateMaps()
	firstPart()
	secondPart()
}

func generateMaps() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\r\n")
	for y, line := range lines {
		digitsReg := regexp.MustCompile(`\d+`)
		specialReg := regexp.MustCompile(`[^0-9.]`)
		numbersMatches := digitsReg.FindAllStringIndex(line, -1)
		specialMatches := specialReg.FindAllStringIndex(line, -1)

		for _, match := range numbersMatches {
			start, end := match[0], match[1]
			number := line[start:end]
			fmt.Printf("Number: %s, Position: %d,%d\n", number, start, y)
			nums[Point{start, y}] = Number{Point{start, y}, number}
		}

		for _, match := range specialMatches {
			start := match[0]
			char := line[start]
			fmt.Printf("Special: %s, Position: %d,%d\n", string(char), start, y)
			specials[Point{start, y}] = Special{Point{start, y}, rune(char), []Number{}}
		}

	}
}

func firstPart() {
	res := 0
	for _, num := range nums {
		if hasAdjacentSymbol(num) {
			n, err := strconv.Atoi(num.value)
			if err != nil {
				log.Fatal(err)
			}
			res += n
		}
	}
	fmt.Println("Result: ", res)
}

func secondPart() {
	res := 0
	for _, s := range specials {
		fmt.Printf("Special: %s, Adjacent numbers: %v\n", string(s.value), s.adjacentNumbers)
		if s.value == '*' && len(s.adjacentNumbers) == 2 {
			num1, err := strconv.Atoi(s.adjacentNumbers[0].value)
			if err != nil {
				log.Fatal(err)
			}
			num2, err := strconv.Atoi(s.adjacentNumbers[1].value)
			if err != nil {
				log.Fatal(err)
			}
			res += num1 * num2
		}
	}
	fmt.Println("Result: ", res)
}

func isSymbolAt(x, y int, num Number) bool {
	point := Point{x, y}
	if s, ok := specials[point]; ok {
		fmt.Printf("Special: %+v", s)
		s.adjacentNumbers = append(s.adjacentNumbers, num)
		specials[point] = s // Put the modified value back in the map
		return true
	}
	return false
}

func hasAdjacentSymbol(num Number) bool {
	for i := 0; i < len(num.value); i++ {
		// Check the eight possible adjacent positions
		if isSymbolAt(num.x-1, num.y-1, num) ||
			isSymbolAt(num.x, num.y-1, num) ||
			isSymbolAt(num.x+1, num.y-1, num) ||
			isSymbolAt(num.x-1, num.y, num) ||
			isSymbolAt(num.x+1, num.y, num) ||
			isSymbolAt(num.x-1, num.y+1, num) ||
			isSymbolAt(num.x, num.y+1, num) ||
			isSymbolAt(num.x+1, num.y+1, num) {
			return true
		}
		num.x++ // Move to the next character in the number
	}
	return false
}
