package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	firstPart()
	secondPart()
}

func firstPart() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\r\n")

	totalPoints := 0

	for _, line := range lines {

		points := getPoints(line)
		totalPoints += points
	}
	fmt.Println("Total points:", totalPoints)
}

type Card struct {
	matches int
	copies  int
}

func secondPart() {
	content, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\r\n")
	cardPoints := make(map[int]*Card)

	for i, line := range lines {
		matches := getMatches(line)
		cardPoints[i] = &Card{matches, 0}
		if matches > 0 {
			for ii := i + 1; ii <= matches; ii++ {
				if cardPoints[ii] != nil {
					cardPoints[ii].copies++
				}
			}
		}
	}

	for _, card := range cardPoints {
		// I dont know what to do here
		fmt.Printf("%v", card)
	}
}

func getPoints(line string) int {
	points := 0

	count := getMatches(line)
	if count == 0 {
		points = 0
	} else if count == 1 {
		points = 1
	} else {
		points = int(math.Pow(2, float64(count-1)))
	}
	return points
}

func getMatches(line string) int {
	count := 0

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

	numbers := strings.Split(numStr, " ")
	for _, n := range numbers {
		if winningNums[n] {
			count++
		}
	}
	return count
}
