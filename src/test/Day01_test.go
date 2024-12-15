package test

import (
	"advent-of-code-2024/src/main"
	"testing"
)

type Day01Tester struct {
	Tester
	left  []int
	right []int
	solve solution.Day01
}

//3   4
//4   3
//2   5
//1   3
//3   9
//3   3

func NewDay01Tester(year int, day int) *Day01Tester {
	return &Day01Tester{
		Tester: NewTester(year, day),
		left:   []int{3, 4, 2, 1, 3, 3},
		right:  []int{4, 3, 5, 3, 9, 3},
		solve:  solution.Day01{},
	}
}

func (tester Day01Tester) TestPart01Mock(t *testing.T) {
	answer, _ := tester.solve.Part01(tester.left, tester.right)
	want := 11
	if answer != want {
		t.Errorf("got %d, want %d", answer, want)
	}
}

func (tester Day01Tester) TestPart01(t *testing.T) {
	left, right := tester.Tester.ParseInput()
	answer, _ := tester.solve.Part01(left, right)
	want := 2367773
	if answer != want {
		t.Errorf("got %d, want %d", answer, want)
	}
}

func (tester Day01Tester) TestPart02Mock(t *testing.T) {
	answer, _ := tester.solve.Part02(tester.left, tester.right)
	want := 31
	if answer != want {
		t.Errorf("got %d, want %d", answer, want)
	}
}

func (tester Day01Tester) TestPart02(t *testing.T) {
	left, right := tester.Tester.ParseInput()
	answer, _ := tester.solve.Part02(left, right)
	want := 21271939
	if answer != want {
		t.Errorf("got %d, want %d", answer, want)
	}
}

func TestDay01(t *testing.T) {
	tester := NewDay01Tester(2024, 01)

	t.Run("Part1 (Mock)", tester.TestPart01Mock)
	t.Run("Part2 (Mock)", tester.TestPart02Mock)
	t.Run("Part1", tester.TestPart01)
	t.Run("Part2", tester.TestPart02)
}
