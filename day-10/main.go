package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

/*
┌─┐
│ │
└─┘
*/
var pipes = map[rune][]int{
	'|': {0, 1, 0, 1}, // │
	'-': {1, 0, 1, 0}, // ─
	'L': {0, 1, 1, 0}, // └
	'J': {1, 1, 0, 0}, // ┘
	'7': {1, 0, 0, 1}, // ┐
	'F': {0, 0, 1, 1}, // ┌
	'.': {0, 0, 0, 0}, // No pipe
	'S': {1, 1, 1, 1}, // Start, all directions
	'X': {0, 0, 0, 0}, // Visited
}

/*
| is a vertical pipe connecting north and south.
| = [0,1,0,1]

- is a horizontal pipe connecting east and west.
- = [1,0,1,0]

L is a 90-degree bend connecting north and east.
L = [0,1,1,0]

J is a 90-degree bend connecting north and west.
J = [0,0,1,1]

7 is a 90-degree bend connecting south and west.
7 = [1,1,0,0]

F is a 90-degree bend connecting south and east.
F = [1,0,0,1]

. is ground; there is no pipe in this tile.
. = [0,0,0,0]

S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.
S = [1,1,1,1]
*/

type Matrix [][]rune

func (m *Matrix) walkPipe(x, y int, dir rune) bool {
	pipe := pipes[(*m)[y][x]]

	initX, initY := x, y

	switch dir {
	case '>':
		x++
		if x >= len((*m)[y]) {
			return false
		}
		destPipe := pipes[(*m)[y][x]]
		if pipe[2] == 0 || destPipe[0] == 0 {
			return false
		}
	case '<':
		x--
		if x < 0 {
			return false
		}
		destPipe := pipes[(*m)[y][x]]
		if pipe[0] == 0 || destPipe[2] == 0 {
			return false
		}
	case '^':
		y--
		if y < 0 {
			return false
		}
		destPipe := pipes[(*m)[y][x]]
		if pipe[1] == 0 || destPipe[3] == 0 {
			return false
		}
	case 'v':
		y++
		if y >= len(*m) {
			return false
		}
		destPipe := pipes[(*m)[y][x]]
		if pipe[3] == 0 || destPipe[1] == 0 {
			return false
		}
	}
	if (*m)[initY][initX] != 'S' {
		(*m)[initY][initX] = 'X'
	}
	return true
}

func (m *Matrix) render(col, row int) {
	for y, line := range *m {
		for x, char := range line {
			if col == x && row == y {
				fmt.Printf("\033[31m%c\033[0m", char)
			} else {
				fmt.Printf("%c", char)
			}
		}
		fmt.Println()
	}
}

func (m *Matrix) findStart() (int, int) {
	for y, line := range *m {
		for x, char := range line {
			if char == 'S' {
				return x, y
			}
		}
	}
	return 0, 0
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := part1(content)
	fmt.Println("Result part1: ", result)

	// result = part2(content)
	// fmt.Println("Result part2: ", result)
}

func part1(content []byte) int {
	lines := strings.Split(string(content), "\n")
	m := make(Matrix, len(lines))
	for i, s := range lines {
		m[i] = []rune(s)
	}

	x, y := m.findStart()
	m.render(x, y)

	steps := 0
	for {
		// fmt.Scanln() // wait for key press
		for _, dir := range []rune{'^', '>', 'v', '<'} {
			if m.walkPipe(x, y, dir) {
				switch dir {
				case '>':
					x++
				case '<':
					x--
				case '^':
					y--
				case 'v':
					y++
				}
				// m.render(x, y)
				steps++
				if m[y][x] == 'S' {
					return steps / 2
				}
				break
			}
		}
	}
}
