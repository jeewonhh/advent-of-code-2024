package test

import (
	solution "advent-of-code-2024/src/main"
	"os"
	"strings"
	"testing"
)

type Day04Tester struct {
	Tester
	solve  solution.Day04
	puzzle []string
}

func NewDay04Tester(year int, day int) *Day04Tester {
	return &Day04Tester{
		Tester: NewTester(year, day),
		solve:  solution.Day04{},
		puzzle: []string{
			"MMMSXXMASM",
			"MSAMXMSMSA",
			"AMXSXMAAMM",
			"MSAMASMSMX",
			"XMASAMXAMM",
			"XXAMMXXAMA",
			"SMSMSASXSS",
			"SAXAMASAAA",
			"MAMMMXMMMM",
			"MXMXAXMASX",
		},
	}
}

func (tester Day04Tester) TestPart01Mock(t *testing.T) {
	answer := tester.solve.Part01(tester.puzzle)
	want := 18
	if answer != want {
		t.Errorf("got %d, want %d", answer, want)
	}
}

func (tester Day04Tester) TestPart01(t *testing.T) {
	bytes, _ := os.ReadFile(tester.input_file)
	puzzle := strings.Split(string(bytes), "\n")
	answer := tester.solve.Part01(puzzle)
	want := 2493
	if answer != want {
		t.Errorf("got %d, want %d", answer, want)
	}
}

//func (tester Day04Tester) TestPart02Mock(t *testing.T) {
//	answer := tester.solve.Part02(tester.records)
//	want := 4
//	if answer != want {
//		t.Errorf("got %d, want %d", answer, want)
//	}
//}
//
//func (tester Day04Tester) TestPart02(t *testing.T) {
//	records := utils.Parse2DInts(tester.input_file)
//	answer := tester.solve.Part02(records)
//	want := 301
//	if answer != want {
//		t.Errorf("got %d, want %d", answer, want)
//	}
//}

func TestDay04(t *testing.T) {
	tester := NewDay04Tester(2024, 04)

	t.Run("Part1 (Mock)", tester.TestPart01Mock)
	t.Run("Part1", tester.TestPart01)
	//t.Run("Part2 (Mock)", tester.TestPart02Mock)
	//t.Run("Part2", tester.TestPart02)
}
