package day8

import (
	"lcatania/aoc-2024-go/utils"
	"strings"
)

func Day8() int {
	fileContent := utils.ReadFile("./day8/input.txt")
	lines := strings.Split(fileContent, "\n")
	grid := make([][]rune, len(lines))
	for y, line := range lines {
		grid[y] = make([]rune, len(line))
		for x, char := range line {
			grid[y][x] = char
		}
	}
	antinodes := calcAntinodes(grid)
	return calcAntinodesCount(antinodes)
}

func Day8Part2() int {
	fileContent := utils.ReadFile("./day8/input.txt")
	lines := strings.Split(fileContent, "\n")
	grid := make([][]rune, len(lines))
	for y, line := range lines {
		grid[y] = make([]rune, len(line))
		for x, char := range line {
			grid[y][x] = char
		}
	}
	antinodes := calcAntinodesP2(grid)
	return calcAntinodesCount(antinodes)
}

func calcAntinodesCount(antinodes [][]rune) int {
	var sum int
	for _, row := range antinodes {
		for _, v := range row {
			if v == '#' {
				sum++
			}
		}
	}
	return sum
}

func calcAntinodes(amap [][]rune) [][]rune {
	antennas := buildAntennasMap(amap)

	antinodes := buildAntinodes(amap)

	for _, positions := range antennas {
		for i := 0; i < len(positions)-1; i++ {
			p1 := positions[i]
			for j := i + 1; j < len(positions); j++ {
				p2 := positions[j]
				vec := subVec(p1, p2)
				a := [][]int{addVec(p1, vec), subVec(p2, vec)}

				for _, antifrequency := range a {
					if isVecIn2DBounds(antinodes, antifrequency) {
						antinodes[antifrequency[0]][antifrequency[1]] = '#'
					}
				}
			}
		}
	}

	return antinodes
}

func calcAntinodesP2(amap [][]rune) [][]rune {
	antennas := buildAntennasMap(amap)

	antinodes := buildAntinodes(amap)

	for _, positions := range antennas {
		for i := 0; i < len(positions)-1; i++ {
			p1 := positions[i]
			for j := i + 1; j < len(positions); j++ {
				p2 := positions[j]
				vec := subVec(p1, p2)

				for antifrequency := addVec(p1, vec); isVecIn2DBounds(antinodes, antifrequency); antifrequency = addVec(antifrequency, vec) {
					antinodes[antifrequency[0]][antifrequency[1]] = '#'
				}

				for antifrequency := subVec(p2, vec); isVecIn2DBounds(antinodes, antifrequency); antifrequency = subVec(antifrequency, vec) {
					antinodes[antifrequency[0]][antifrequency[1]] = '#'
				}

				antinodes[p1[0]][p1[1]] = '#'
				antinodes[p2[0]][p2[1]] = '#'
			}
		}
	}

	return antinodes
}

func buildAntennasMap(amap [][]rune) map[rune][][]int {
	antennas := make(map[rune][][]int)
	for i, row := range amap {
		for j, v := range row {
			if v != '.' {
				pos := []int{i, j}
				if _, exist := antennas[v]; !exist {
					antennas[v] = [][]int{pos}
				} else {
					antennas[v] = append(antennas[v], pos)
				}
			}
		}
	}
	return antennas
}

func buildAntinodes(amap [][]rune) [][]rune {
	antinodes := make([][]rune, len(amap))
	for i, _ := range amap {
		antinodes[i] = make([]rune, len(amap[i]))
		for j := range antinodes[i] {
			antinodes[i][j] = '.'
		}
	}
	return antinodes
}

func isVecIn2DBounds[V int](board [][]rune, vec []V) bool {
	if len(vec) != 2 {
		panic("Vector should be 2D.")
	}

	y := int(vec[0])
	x := int(vec[1])
	if y < 0 || y >= len(board) || x < 0 || x >= len(board[y]) {
		return false
	}

	return true
}

func addVec[V int](a, b []V) []V {
	checkVectorsMatch(a, b)

	r := make([]V, len(a))
	for i := range a {
		r[i] = a[i] + b[i]
	}

	return r
}

func subVec[V int](a, b []V) []V {
	checkVectorsMatch(a, b)

	r := make([]V, len(a))
	for i := range a {
		r[i] = a[i] - b[i]
	}

	return r
}

func checkVectorsMatch[V int](a, b []V) {
	if len(a) != len(b) {
		panic("Vectors not match")
	}
}
