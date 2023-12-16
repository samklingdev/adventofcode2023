package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	firstPart()
}

type Num struct {
	y   int
	x   int
	len int
	n   int
}

type SpecialChar struct {
	y    int
	x    int
	char rune
}

var nums = []Num{}
var specialChars = []SpecialChar{}

func firstPart() {
	content, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	// result := 0
	currNumStr := []rune{}

	lines := strings.Split(string(content), "\n")
	for y, line := range lines {
		for x, char := range line {
			// check if char is a digit
			if unicode.IsDigit(char) {
				// check if we got a character anywhere around it
				currNumStr = append(currNumStr, char)
			} else if len(currNumStr) > 0 {
				n, err := strconv.Atoi(string(currNumStr))
				if err != nil {
					log.Fatal(err)
				}
				num := Num{
					x:   x - len(currNumStr),
					y:   y,
					len: len(currNumStr),
					n:   n,
				}
				nums = append(nums, num)
				currNumStr = []rune{}
			} else if char != 13 {
				specialChar := SpecialChar{
					x:    x,
					y:    y,
					char: char,
				}
				specialChars = append(specialChars, specialChar)
			}
		}
	}
	/*
		for each num
		*****
		*123*
		*****

	*/
	fmt.Printf("Nums: %+v\n", nums)
	fmt.Printf("Special Chars: %+v\n", specialChars)
}
