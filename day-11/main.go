package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	content, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := part1(content)
	fmt.Println("Result part1: ", result)

	// result = part2(content)
	// fmt.Println("Result part2: ", result)
}

func part1(content []byte) int {
	return 0
}

func part2(content []byte) int {
	return 0
}
