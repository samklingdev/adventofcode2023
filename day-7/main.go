package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Hand struct {
	cards []rune
	bid   int
}

var orderPart1 = "AKQJT98765432"
var orderPart2 = "AKQT98765432J"

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := part1(content)
	fmt.Println("Part 1 result:", result)

	result = part2(content)
	fmt.Println("Part 2 result:", result)
}

func part1(content []byte) int {
	var cards string
	lines := strings.Split(string(content), "\n")
	hands := make([]Hand, len(lines))

	for i, line := range lines {
		fmt.Sscanf(line, "%s %d", &cards, &hands[i].bid)
		hands[i].cards = []rune(cards)
	}

	sort.Slice(hands, func(i, j int) bool {
		return sortHandsPart1(hands[i], hands[j])
	})

	result := 0
	for i, h := range hands {
		result += h.bid * (i + 1)
	}

	return result
}

func part2(content []byte) int {
	var cards string
	lines := strings.Split(string(content), "\n")
	hands := make([]Hand, len(lines))

	for i, line := range lines {
		fmt.Sscanf(line, "%s %d", &cards, &hands[i].bid)
		hands[i].cards = []rune(cards)
	}

	sort.Slice(hands, func(i, j int) bool {
		return sortHandsPart2(hands[i], hands[j])
	})

	result := 0
	for i, h := range hands {
		result += h.bid * (i + 1)
	}
	return result
}

func sortHandsPart2(a Hand, b Hand) bool {
	cardsLen := len(a.cards)
	aMap, bMap := map[rune]int{}, map[rune]int{}
	aStr, bStr := []int{}, []int{}

	for i := 0; i < cardsLen; i++ {
		aMap[a.cards[i]]++
		bMap[b.cards[i]]++
	}

	for k, v := range aMap {
		if k != 'J' {
			aStr = append(aStr, v)
		}
	}
	for k, v := range bMap {
		if k != 'J' {
			bStr = append(bStr, v)
		}
	}

	sort.Slice(aStr, func(i, j int) bool { return aStr[i] > aStr[j] })
	sort.Slice(bStr, func(i, j int) bool { return bStr[i] > bStr[j] })

	aJokers, ok := aMap['J']
	if ok {
		if len(aStr) == 0 {
			aStr = append(aStr, 5)
		} else {
			aStr[0] += aJokers
		}
	}
	bJokers, ok := bMap['J']
	if ok {
		if len(bStr) == 0 {
			bStr = append(bStr, 5)
		} else {
			bStr[0] += bJokers
		}
	}

	aStrength, bStrength := getStrength(aStr), getStrength(bStr)

	if aStrength < bStrength {
		return true
	} else if aStrength > bStrength {
		return false
	} else {
		for i := 0; i < len(a.cards); i++ {
			aIndex := strings.Index(orderPart2, string(a.cards[i]))
			bIndex := strings.Index(orderPart2, string(b.cards[i]))
			if aIndex > bIndex {
				return true
			} else if aIndex < bIndex {
				return false
			}
		}
	}
	return false
}

func sortHandsPart1(a Hand, b Hand) bool {

	aMap, bMap := map[rune]int{}, map[rune]int{}
	aStr, bStr := []int{}, []int{}

	for i := 0; i < len(a.cards); i++ {
		aMap[a.cards[i]]++
		bMap[b.cards[i]]++
	}

	for _, v := range aMap {
		aStr = append(aStr, v)
	}
	for _, v := range bMap {
		bStr = append(bStr, v)
	}

	sort.Slice(aStr, func(i, j int) bool { return aStr[i] > aStr[j] })
	sort.Slice(bStr, func(i, j int) bool { return bStr[i] > bStr[j] })

	aStrength, bStrength := getStrength(aStr), getStrength(bStr)

	if aStrength < bStrength {
		return true
	} else if aStrength > bStrength {
		return false
	} else {
		for i := 0; i < len(a.cards); i++ {
			aIndex := strings.Index(orderPart1, string(a.cards[i]))
			bIndex := strings.Index(orderPart1, string(b.cards[i]))
			if aIndex > bIndex {
				return true
			} else if aIndex < bIndex {
				return false
			}
		}
	}
	return false
}

func getStrength(v []int) int {
	strValue := fmt.Sprint(v)
	switch strValue {
	case "[5]":
		return 6
	case "[4 1]":
		return 5
	case "[3 2]":
		return 4
	case "[3 1 1]":
		return 3
	case "[2 2 1]":
		return 2
	case "[2 1 1 1]":
		return 1
	default:
		return 0
	}
}
