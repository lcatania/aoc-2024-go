package day25

import (
	"lcatania/aoc-2024-go/utils"
	"strconv"
	"strings"
)

func Day25() string {
	fileContent := utils.ReadFile("./day25/input.txt")
	lines := strings.Split(fileContent, "\n")
	var data [][]string
	for _, line := range lines {
		data = append(data, strings.Split(line, ""))
	}

	var keys, locks [][]int

	for i := 0; i < len(data); i += 8 {
		isLock := data[i][0] == "#"

		combination := make([]int, 5)
		for col := range 5 {
			c := 0
			for row := i + 1; row < i+6; row++ {
				if data[row][col] == "#" {
					c++
				}
			}
			combination[col] = c
		}

		if isLock {
			locks = append(locks, combination)
		} else {
			keys = append(keys, combination)
		}
	}

	result := 0
	for _, key := range keys {
		for _, lock := range locks {
			check := true
			for col := range 5 {
				if key[col]+lock[col] > 5 {
					check = false
					break
				}
			}

			if check {
				result++
			}
		}
	}

	return strconv.Itoa(result)
}

func Day25Part2() string {
	return ""
}
