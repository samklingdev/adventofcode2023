package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

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
	games := strings.Split(string(content), "\n")

	maxRed := 12
	maxGreen := 13
	maxBlue := 14

	result := 0
	for _, game := range games {
		gameId := getGameID(game)
		redBalls := getBallsByColor(game, "red")
		blueBalls := getBallsByColor(game, "blue")
		greenBalls := getBallsByColor(game, "green")

		mRed := max(redBalls)
		mBlue := max(blueBalls)
		mGreen := max(greenBalls)

		if mRed > maxRed || mBlue > maxBlue || mGreen > maxGreen {
			continue
		} else {
			result += gameId
		}
	}
	return result
}

func part2(content []byte) int {
	games := strings.Split(string(content), "\n")
	result := 0

	for _, game := range games {
		redBalls := getBallsByColor(game, "red")
		blueBalls := getBallsByColor(game, "blue")
		greenBalls := getBallsByColor(game, "green")

		mRed := max(redBalls)
		mBlue := max(blueBalls)
		mGreen := max(greenBalls)

		result += mRed * mBlue * mGreen
	}
	return result
}

func getGameID(game string) int {
	re, _ := regexp.Compile(`Game (\d+)`)
	matches := re.FindStringSubmatch(game)
	if matches == nil {
		return 0
	}

	i, err := strconv.Atoi(matches[1])
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func getBallsByColor(game string, color string) []int {
	re, _ := regexp.Compile(`(\d+) ` + color)
	matches := re.FindAllStringSubmatch(game, -1)

	if matches == nil {
		return []int{0}
	}

	result := []int{}

	for _, match := range matches {
		i, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, i)
	}
	return result
}

func max(balls []int) int {
	sort.Ints(balls)
	return balls[len(balls)-1]
}
