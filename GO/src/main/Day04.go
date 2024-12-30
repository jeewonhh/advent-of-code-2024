package solution

import (
	"fmt"
	"strings"
)

type Day04 struct {
}

const XMAS = "XMAS"
const SAMX = "SAMX"
const LEN = 4

const MAS = "MAS"
const SAM = "SAM"
const MASLEN = 3

func lookEverywhere(puzzle []string, row int, col int) int {
	count := 0
	rows := len(puzzle)
	cols := len(puzzle[0])

	// horizontally (to right)
	if col < cols-LEN+1 {
		if col+LEN > cols {
			fmt.Println(col + LEN)
		}
		word := puzzle[row][col : col+LEN]
		if word == XMAS || word == SAMX {
			count++
		}
	}
	// vertically (downwards)
	if row < rows-LEN+1 {
		var sb strings.Builder
		for i := 0; i < LEN; i++ {
			sb.WriteRune(rune(puzzle[row+i][col]))
		}
		if sb.String() == XMAS || sb.String() == SAMX {
			count++
		}
	}
	// diagonally (up and right)
	if row < rows-LEN+1 && col > LEN-2 {
		var sb strings.Builder
		for i := 0; i < LEN; i++ {
			sb.WriteRune(rune(puzzle[row+i][col-i]))
		}
		if sb.String() == XMAS || sb.String() == SAMX {
			count++
		}
	}
	// diagonally (down and right)
	if row < rows-LEN+1 && col < cols-LEN+1 {
		var sb strings.Builder
		for i := 0; i < LEN; i++ {
			sb.WriteRune(rune(puzzle[row+i][col+i]))
		}
		if sb.String() == XMAS || sb.String() == SAMX {
			count++
		}
	}

	return count
}

func lookAtBox(puzzle []string, row int, col int) int {
	rows := len(puzzle)
	cols := len(puzzle[0])

	if row > rows-MASLEN || col > cols-MASLEN {
		return 0
	}

	var sb strings.Builder
	for i := 0; i < MASLEN; i++ {
		sb.WriteRune(rune(puzzle[row+i][col+i]))
	}
	first := sb.String()
	if first != MAS && first != SAM {
		return 0
	}
	sb.Reset()
	for i := 0; i < MASLEN; i++ {
		sb.WriteRune(rune(puzzle[row+MASLEN-i-1][col+i]))
	}
	second := sb.String()
	if second != MAS && second != SAM {
		return 0
	}
	return 1
}

func (d Day04) Part01(puzzle []string) int {
	count := 0
	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle[0]); j++ {
			count += lookEverywhere(puzzle, i, j)
		}
	}
	return count
}

func (d Day04) Part02(puzzle []string) int {
	count := 0

	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle[0]); j++ {
			count += lookAtBox(puzzle, i, j)
		}
	}
	return count
}
