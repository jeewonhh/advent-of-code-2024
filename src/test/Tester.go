package test

import (
	"advent-of-code-2024/src/utils"
	"fmt"
	"testing"
)

type Tester struct {
	year       int
	day        int
	input_file string
	t          *testing.T
}

func NewTester(year int, day int) Tester {
	return Tester{
		year:       year,
		day:        day,
		input_file: fmt.Sprintf("../../inputs/puzzle-%02d.txt", day),
	}
}

func (tester *Tester) ParseInput() ([]int, []int) {
	switch tester.day {
	case 1:
		return utils.ParseTwoIntColumns(tester.input_file)
	default:
		return []int{}, []int{}
	}
}
