package solution

import (
	"advent-of-code-2024/src/utils"
	"fmt"
	"slices"
)

type Day01 struct {
}

func (d Day01) Part01(left []int, right []int) (int, error) {
	if len(left) != len(right) {
		return 0, fmt.Errorf("zip: arguments must be of same length")
	}

	slices.Sort(left)
	slices.Sort(right)

	diff_sum := 0
	for i, left_num := range left {
		diff_sum += utils.Abs(left_num - right[i])
	}

	return diff_sum, nil
}

func (d Day01) Part02(left []int, right []int) (int, error) {
	frequency_map := make(map[int]int)

	for _, right_num := range right {
		frequency_map[right_num] += 1
	}

	result := utils.MapSum(left, func(x int) int {
		return x * frequency_map[x]
	})

	return result, nil
}
