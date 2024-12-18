package test

import (
	"advent-of-code-2024/src/main"
	"os"
	"testing"
)

type Day03Tester struct {
	Tester
	solve  solution.Day03
	memory string
}

func NewDay03Tester(year int, day int) *Day03Tester {
	return &Day03Tester{
		Tester: NewTester(year, day),
		solve:  solution.Day03{},
		memory: "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
	}
}

func (tester Day03Tester) TestPart01Mock(t *testing.T) {
	answer := tester.solve.Part01(tester.memory)
	want := 161
	if answer != want {
		t.Errorf("got %d, want %d", answer, want)
	}
}

func (tester Day03Tester) TestPart01(t *testing.T) {
	memory, _ := os.ReadFile(tester.input_file)
	answer := tester.solve.Part01(string(memory))
	want := 161
	if answer != want {
		t.Errorf("got %d, want %d", answer, want)
	}
}

//func (tester Day03Tester) TestPart02Mock(t *testing.T) {
//	answer := tester.solve.Part02(tester.memory)
//	want := 4
//	if answer != want {
//		t.Errorf("got %d, want %d", answer, want)
//	}
//}
//
//func (tester Day03Tester) TestPart02(t *testing.T) {
//	records := utils.Parse2DInts(tester.input_file)
//	answer := tester.solve.Part02(records)
//	want := 301
//	if answer != want {
//		t.Errorf("got %d, want %d", answer, want)
//	}
//}

func TestDay03(t *testing.T) {
	tester := NewDay03Tester(2024, 03)

	t.Run("Part1 (Mock)", tester.TestPart01Mock)
	t.Run("Part1", tester.TestPart01)
	//t.Run("Part2 (Mock)", tester.TestPart02Mock)
	//t.Run("Part2", tester.TestPart02)
}
