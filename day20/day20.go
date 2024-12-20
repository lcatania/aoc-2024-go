package day20

import (
	_ "embed"
	"lcatania/aoc-2024-go/utils"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Position struct {
	Line, Column int
}

var N = Position{Line: -1, Column: 0}
var S = Position{Line: 1, Column: 0}
var E = Position{Line: 0, Column: 1}
var W = Position{Line: 0, Column: -1}
var DirectionsSlice = []Position{N, S, W, E}

func isValidSize(grid []string, pos Position) bool {
	return pos.Line >= 0 && pos.Column >= 0 && pos.Line < len(grid) && pos.Column < len(grid[0])
}

func isValidFair(grid []string, pos Position) bool {
	return isValidSize(grid, pos) && eval(grid, pos) != '#'
}

func getTimeFair(grid []string, start Position, lastDirection Position, end Position) int {
	if start == end {
		return 1
	}
	for _, direction := range DirectionsSlice {
		if direction != opposedDirection(lastDirection) {
			var newPos = addPositions(start, direction)
			if isValidFair(grid, newPos) {
				return 1 + getTimeFair(grid, newPos, direction, end)
			}
		}
	}
	return 0
}

func getPositionsFair(grid []string, start Position, lastDirection Position, end Position) []Position {
	var length int = getTimeFair(grid, start, lastDirection, end)
	var res []Position = make([]Position, length)
	for i := range length - 1 {
		for _, direction := range DirectionsSlice {
			if direction != opposedDirection(lastDirection) {
				var newPos = addPositions(start, direction)
				if isValidFair(grid, newPos) {
					res[i] = start
					start = newPos
					lastDirection = direction
					break
				}
			}
		}
	}
	res[length-1] = end
	return res
}

func Day20() string {
	fileContent := utils.ReadFile("./day20/input.txt")
	var lines = strings.Split(strings.TrimSuffix(fileContent, "\n"), "\n")
	var start Position = searchStartLines(lines, 'S')
	var end Position = searchStartLines(lines, 'E')

	var positions []Position = getPositionsFair(lines, start, S, end)
	slices.Reverse(positions)
	var res int
	for i, p1 := range positions {
		for j := i + 1; j < len(positions); j++ {
			var p2 Position = positions[j]
			if distance(p1, p2) == 2 && (p1.Line == p2.Line || p1.Column == p2.Column) {
				if j-i-2 >= 100 {
					res++
				}
			}
		}
	}
	return strconv.Itoa(res)
}

func getFirstDirection(grid []string, start Position) Position {
	for _, direction := range DirectionsSlice {
		if eval(grid, addPositions(start, direction)) == '.' {
			return direction
		}
	}
	return defPosition(-1, -1)
}

func Day20Part2() string {
	fileContent := utils.ReadFile("./day20/input.txt")
	var lines = strings.Split(strings.TrimSuffix(fileContent, "\n"), "\n")
	var start Position = searchStartLines(lines, 'S')
	var end Position = searchStartLines(lines, 'E')
	var direction Position = getFirstDirection(lines, start)
	var positions []Position = getPositionsFair(lines, start, direction, end)
	slices.Reverse(positions)
	var res int
	for i, p1 := range positions {
		for j := i + 2; j < len(positions); j++ {
			var p2 Position = positions[j]
			if distance(p1, p2) <= 20 {
				if j-i-distance(p1, p2) >= 100 {
					res++
				}
			}
		}
	}
	return strconv.Itoa(res)
}

func eval(grid []string, pos Position) byte {
	return grid[pos.Line][pos.Column]
}

func opposedDirection(pos Position) Position {
	return defPosition(-pos.Line, -pos.Column)
}

func defPosition(line int, column int) Position {
	return Position{Line: line, Column: column}
}

func addPositions(pos1 Position, pos2 Position) Position {
	return defPosition(pos1.Line+pos2.Line, pos1.Column+pos2.Column)
}

func searchStartLines(grid []string, start rune) Position {
	for i, line := range grid {
		for j, char := range line {
			if char == start {
				return Position{Line: i, Column: j}
			}
		}
	}
	return Position{Line: -1, Column: -1}
}

func distance(p1 Position, p2 Position) int {
	return abs(p1.Line-p2.Line) + abs(p1.Column-p2.Column)
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}
