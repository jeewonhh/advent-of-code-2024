package test

import (
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
		input_file: fmt.Sprintf("../../../inputs/puzzle-%02d.txt", day),
	}
}
