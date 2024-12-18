package test

import (
	"advent-of-code-2024/src/main"
	"os"
	"testing"
)

type Day03Tester struct {
	Tester
	solve solution.Day03
}

func NewDay03Tester(year int, day int) *Day03Tester {
	return &Day03Tester{
		Tester: NewTester(year, day),
		solve:  solution.Day03{},
	}
}

func (tester Day03Tester) TestPart01Mock(t *testing.T) {
	answer := tester.solve.Part01("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))\n")
	want := 161
	if answer != want {
		t.Errorf("got %d, want %d", answer, want)
	}
}

func (tester Day03Tester) TestPart01(t *testing.T) {
	memory, _ := os.ReadFile(tester.input_file)
	answer := tester.solve.Part01(string(memory))
	want := 187833789
	if answer != want {
		t.Errorf("got %d, want %d", answer, want)
	}
}

func (tester Day03Tester) TestPart02Mock(t *testing.T) {
	answer := tester.solve.Part02("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))\n")
	want := 48
	if answer != want {
		t.Errorf("got %d, want %d", answer, want)
	}
}

func (tester Day03Tester) TestPart02(t *testing.T) {
	memory, _ := os.ReadFile(tester.input_file)
	answer := tester.solve.Part02(string(memory))
	want := 94455185
	if answer != want {
		t.Errorf("got %d, want %d", answer, want)
	}
}

func TestDay03(t *testing.T) {
	tester := NewDay03Tester(2024, 03)

	t.Run("Part1 (Mock)", tester.TestPart01Mock)
	t.Run("Part1", tester.TestPart01)
	t.Run("Part2 (Mock)", tester.TestPart02Mock)
	t.Run("Part2", tester.TestPart02)
}
