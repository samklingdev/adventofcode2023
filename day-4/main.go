package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := part1(content)
	fmt.Println("Part 1:", result)

	result = part2(content)
	fmt.Println("Part 2:", result)

}

func part1(content []byte) int {
	lines := strings.Split(string(content), "\n")

	totalPoints := 0

	for _, line := range lines {
		winners := getMatches(line)
		totalPoints += int(math.Pow(2, float64(winners-1)))
	}
	return totalPoints
}

func part2(content []byte) int {
	lines := strings.Split(string(content), "\n")
	cards := make([]int, len(lines))

	for i, line := range lines {
		winners := getMatches(line)
		cards[i]++
		for j := 1; j <= winners; j++ {
			cards[i+j] += cards[i]
		}
	}

	result := 0
	for _, card := range cards {
		result += card
	}
	return result
}

func getMatches(line string) int {

	cardParts := strings.Split(line, ":")
	numberParts := strings.Split(cardParts[1], "|")

	var card int
	fmt.Sscanf(cardParts[0], "Card %d", &card)

	winStr := strings.TrimSpace(numberParts[0])
	numStr := strings.TrimSpace(numberParts[1])

	winningNums := make(map[string]bool)
	for _, n := range strings.Split(winStr, " ") {
		if strings.TrimSpace(n) != "" {
			winningNums[n] = true
		}
	}
	count := 0
	numbers := strings.Split(numStr, " ")
	for _, n := range numbers {
		if winningNums[n] {
			count++
		}
	}
	return count
}
