package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var words = map[string]rune{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

func main() {
	firstPart()
	secondPart()
}

func firstPart() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var result int
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		numChars := []rune{}
		for _, c := range line {
			if unicode.IsDigit(c) {
				numChars = append(numChars, c)
			}
		}
		first := numChars[0]
		last := numChars[len(numChars)-1]
		intValue, err := strconv.Atoi(fmt.Sprintf("%c%c", first, last))
		if err != nil {
			log.Fatal(err)
		}
		result += intValue
	}
	fmt.Println("Sum: ", result)
}

func secondPart() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	var result int
	for _, line := range lines {
		numChars := []rune{}
		lineLength := len(line)
		for index, char := range line {
			if unicode.IsDigit(char) {
				numChars = append(numChars, char)
			} else {
				for key, val := range words {
					if index+len(key) <= lineLength && line[index:index+len(key)] == key {
						numChars = append(numChars, val)
					}
				}
			}
		}
		first := numChars[0]
		last := numChars[len(numChars)-1]
		intValue, err := strconv.Atoi(fmt.Sprintf("%c%c", first, last))
		if err != nil {
			log.Fatal(err)
		}
		result += intValue
	}
	fmt.Println("Sum: ", result)
}
