package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseTwoIntColumns(filename string) ([]int, []int) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Print("There has been an error!: ", err)
	}
	defer f.Close()

	var left, right []int
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			continue
		}

		fields := strings.Fields(line)

		// Convert the fields to integers
		num1, err := strconv.Atoi(fields[0])
		if err != nil {
			fmt.Println("Error converting first field to integer:", err)
			continue
		}
		num2, err := strconv.Atoi(fields[1])
		if err != nil {
			fmt.Println("Error converting second field to integer:", err)
			continue
		}

		// Append the integers to the respective slices
		left = append(left, num1)
		right = append(right, num2)
	}
	return left, right
}

func Parse2DInts(filename string) [][]int {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Print("There has been an error!: ", err)
	}
	defer f.Close()

	var records [][]int
	record_index := 0

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			continue
		}

		fields := strings.Fields(line)
		records = append(records, Map(fields, strconv.Atoi))
		record_index++
	}

	return records
}
