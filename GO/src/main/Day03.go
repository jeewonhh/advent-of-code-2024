package solution

import (
	"strconv"
	"unicode"
)

type Day03 struct {
}

func (d Day03) Part01(memory string) int {
	i, j := 0, 0
	answer := 0
	for i < len(memory)-4 {
		if memory[i:i+4] == "mul(" {
			start := i + 4
			j = start
			for j < start+3 && unicode.IsDigit(rune(memory[j])) {
				j++
			}
			if (j-start) < 1 || memory[j] != ',' {
				i = j
				continue
			}
			a, _ := strconv.Atoi(memory[start:j])
			start = j + 1
			j = start
			for j < start+3 && unicode.IsDigit(rune(memory[j])) {
				j++
			}
			if (j-start) < 1 || memory[j] != ')' {
				i = j
				continue
			}
			b, _ := strconv.Atoi(memory[start:j])
			answer += a * b
			i = j
		} else {
			i++
		}
	}
	return answer
}
