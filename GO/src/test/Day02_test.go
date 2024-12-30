package test

import (
	"advent-of-code-2024/src/main"
	"advent-of-code-2024/src/utils"
	"testing"
)

type Day02Tester struct {
	Tester
	solve   solution.Day02
	records [][]int
}

//7 6 4 2 1
//1 2 7 8 9
//9 7 6 2 1
//1 3 2 4 5
//8 6 4 4 1
//1 3 6 7 9

func NewDay02Tester(year int, day int) *Day02Tester {
	return &Day02Tester{
		Tester: NewTester(year, day),
		solve:  solution.Day02{},
		records: [][]int{
			{7, 6, 4, 2, 1},
			{1, 2, 7, 8, 9},
			{9, 7, 6, 2, 1},
			{1, 3, 2, 4, 5},
			{8, 6, 4, 4, 1},
			{1, 3, 6, 7, 9},
		},
	}
}

func (tester Day02Tester) TestPart01Mock(t *testing.T) {
	answer := tester.solve.Part01(tester.records)
	want := 2
	if answer != want {
		t.Errorf("got %d, want %d", answer, want)
	}
}

func (tester Day02Tester) TestPart01(t *testing.T) {
	records := utils.Parse2DInts(tester.input_file)
	answer := tester.solve.Part01(records)
	want := 230
	if answer != want {
		t.Errorf("got %d, want %d", answer, want)
	}
}

func (tester Day02Tester) TestPart02Mock(t *testing.T) {
	answer := tester.solve.Part02(tester.records)
	want := 4
	if answer != want {
		t.Errorf("got %d, want %d", answer, want)
	}
}

func (tester Day02Tester) TestPart02(t *testing.T) {
	records := utils.Parse2DInts(tester.input_file)
	answer := tester.solve.Part02(records)
	want := 301
	if answer != want {
		t.Errorf("got %d, want %d", answer, want)
	}
}

func TestDay02(t *testing.T) {
	tester := NewDay02Tester(2024, 02)

	t.Run("Part1 (Mock)", tester.TestPart01Mock)
	t.Run("Part1", tester.TestPart01)
	t.Run("Part2 (Mock)", tester.TestPart02Mock)
	t.Run("Part2", tester.TestPart02)
}
