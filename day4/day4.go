package day4

import (
	"lcatania/aoc-2024-go/utils"
	"strings"
)

func Day4() int {
	fileContent := utils.ReadFile("./day4/input.txt")
	lines := strings.Split(fileContent, "\n")
	result := 0
	for i, line := range lines {
		for j, c := range line {
			if c == 'X' {
				// forward
				if j < (len(line)-3) && line[j+1] == 'M' && line[j+2] == 'A' && line[j+3] == 'S' {
					result++
				}

				// reverse
				if j > 2 && line[j-1] == 'M' && line[j-2] == 'A' && line[j-3] == 'S' {

					result++
				}

				// down
				if i < (len(lines)-3) && lines[i+1][j] == 'M' && lines[i+2][j] == 'A' && lines[i+3][j] == 'S' {

					result++
				}
				// up
				if i > 2 && lines[i-1][j] == 'M' && lines[i-2][j] == 'A' && lines[i-3][j] == 'S' {
					result++
				}

				// diagonal - down
				if i < (len(lines)-3) && j < (len(lines)-3) && lines[i+1][j+1] == 'M' && lines[i+2][j+2] == 'A' && lines[i+3][j+3] == 'S' {
					result++
				}

				// diagonal - up
				if i > 2 && j < (len(line)-3) && lines[i-1][j+1] == 'M' && lines[i-2][j+2] == 'A' && lines[i-3][j+3] == 'S' {
					result++
				}

				// reverse - diagonal - down
				if i < (len(lines)-3) && j > 2 && lines[i+1][j-1] == 'M' && lines[i+2][j-2] == 'A' && lines[i+3][j-3] == 'S' {
					result++
				}

				// reverse - diagonal - up
				if i > 2 && j > 2 && lines[i-1][j-1] == 'M' && lines[i-2][j-2] == 'A' && lines[i-3][j-3] == 'S' {
					result++
				}

			}
		}
	}
	return result
}

func Day4Part2() int {
	fileContent := utils.ReadFile("./day4/input.txt")
	lines := strings.Split(fileContent, "\n")
	result := 0
	for i := 1; i < len(lines)-1; i++ {
		for j := 1; j < len(lines[i])-1; j++ {
			if lines[i][j] == 'A' {
				diag0 := lines[i-1][j-1]|lines[i+1][j+1] == 'S'|'M'
				diag1 := lines[i+1][j-1]|lines[i-1][j+1] == 'S'|'M'
				if diag0 && diag1 {
					result++
				}
			}
		}
	}
	return result
}
