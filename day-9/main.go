package main

import (
	"fmt"
	"log"
	"os"
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
	lines := strings.Split(string(content), "\n")

	result := 0

	for _, line := range lines {

		// Split the string into a slice of strings
		numbersStr := strings.Fields(line)

		// Create a slice to hold the integers
		numbers := make([]int, len(numbersStr))

		// Convert each string to an integer
		for i, numberStr := range numbersStr {
			number, err := strconv.Atoi(numberStr)
			if err != nil {
				log.Fatal(err)
			}
			numbers[i] = number
		}

		matrix := [][]int{numbers}

		for {
			diffs, sum := diffsAndSum(numbers)
			numbers = diffs
			matrix = append(matrix, diffs)
			if sum == 0 {
				break
			}
		}

		// loop through the matrix backwards
		prediction := 0
		for i := len(matrix) - 1; i >= 0; i-- {
			row := matrix[i]
			prediction += row[len(row)-1]
			matrix[i] = append(row, prediction)
		}
		result += prediction
	}
	return result
}

func part2(content []byte) int {
	lines := strings.Split(string(content), "\n")

	result := 0

	for _, line := range lines {

		// Split the string into a slice of strings
		numbersStr := strings.Fields(line)

		// Create a slice to hold the integers
		numbers := make([]int, len(numbersStr))

		// Convert each string to an integer
		for i, numberStr := range numbersStr {
			number, err := strconv.Atoi(numberStr)
			if err != nil {
				log.Fatal(err)
			}
			numbers[i] = number
		}

		matrix := [][]int{numbers}

		for {
			diffs, sum := diffsAndSum(numbers)
			numbers = diffs
			matrix = append(matrix, diffs)
			if sum == 0 {
				break
			}
		}

		// loop through the matrix backwards
		extrapolate := 0
		for i := len(matrix) - 1; i >= 0; i-- {
			row := matrix[i]
			extrapolate = row[0] - extrapolate
			matrix[i] = append([]int{extrapolate}, matrix[i]...)
		}
		result += extrapolate
	}
	return result
}

func diffsAndSum(nums []int) (result []int, sum int) {
	diffs := make([]int, len(nums)-1)
	for i := 0; i < len(nums)-1; i++ {
		diffs[i] = nums[i+1] - nums[i]
		sum += diffs[i]
	}
	return diffs, sum
}
