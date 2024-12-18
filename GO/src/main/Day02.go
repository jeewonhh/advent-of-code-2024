package solution

import (
	"advent-of-code-2024/src/utils"
)

type Day02 struct{}

func isValid(a int, b int, increasing bool) bool {
	if increasing {
		a, b = b, a
	}
	return 0 < a-b && a-b < 4
}

func isSafe(record []int) bool {
	if len(record) < 2 {
		return true
	}
	increasing := true
	if record[0] > record[1] {
		increasing = false
	}
	for i := 0; i < len(record)-1; i++ {
		if !isValid(record[i], record[i+1], increasing) {
			return false
		}
	}
	return true
}

func isCloseEnough(record []int) bool {
	if len(record) < 2 {
		return true
	}
	increasing := true
	if record[0] > record[1] {
		increasing = false
	}
	for i := 0; i < len(record)-1; i++ {
		if isValid(record[i], record[i+1], increasing) {
			continue
		}
		//. . . (i-1) (i) (i+1) (i+2) . . .
		a := isSafe(utils.RemoveAtIndex(record, i)) ||
			isSafe(utils.RemoveAtIndex(record, i+1)) ||
			(i > 0 && isSafe(utils.RemoveAtIndex(record, i-1)))
		return a
	}
	return true
}

func (d Day02) Part01(records [][]int) int {
	saves := 0
	for _, record := range records {
		saves += utils.BoolToInt(isSafe(record))
	}
	return saves
}

func (d Day02) Part02(records [][]int) int {
	saves := 0
	for _, record := range records {
		saves += utils.BoolToInt(isCloseEnough(record))
	}
	return saves
}
