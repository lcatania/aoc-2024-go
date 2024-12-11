package day10

import (
	"lcatania/aoc-2024-go/utils"
	"log"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func (p Point) Add(q Point) Point {
	return Point{X: p.X + q.X, Y: p.Y + q.Y}
}

func (p Point) GetAdjacent() []Point {
	adjacent := make([]Point, 0)
	for _, direction := range []Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
		adjacent = append(adjacent, p.Add(direction))
	}
	return adjacent
}

func Day10() int {
	fileContent := utils.ReadFile("./day10/input.txt")
	lines := strings.Split(fileContent, "\n")
	grid := make(map[Point]int)
	for y, line := range lines {
		for x, c := range strings.Split(line, "") {
			number, err := strconv.Atoi(c)
			if err != nil {
				log.Fatal("Could not convert to number")
			}
			grid[Point{x, y}] = number
		}
	}
	result := 0
	for p, v := range grid {
		if v == 0 {
			heads, _ := findPaths(grid, p, make(map[Point]bool), 0)
			result += len(heads)
		}
	}

	return result
}

func Day10Part2() int {
	fileContent := utils.ReadFile("./day10/input.txt")
	lines := strings.Split(fileContent, "\n")
	grid := make(map[Point]int)
	for y, line := range lines {
		for x, c := range strings.Split(line, "") {
			number, err := strconv.Atoi(c)
			if err != nil {
				log.Fatal("Could not convert to number")
			}
			grid[Point{x, y}] = number
		}
	}
	result := 0
	for p, v := range grid {
		if v == 0 {
			_, numberOfPaths := findPaths(grid, p, make(map[Point]bool), 0)

			result += numberOfPaths
		}
	}
	return result
}

func getValidNeighbours(grid map[Point]int, point Point) []Point {
	neighbours := make([]Point, 0)
	for _, n := range point.GetAdjacent() {
		if v, ok := grid[n]; ok && v == grid[point]+1 {
			neighbours = append(neighbours, n)
		}
	}
	return neighbours
}

func findPaths(grid map[Point]int, start Point, heads map[Point]bool, numberOfPaths int) (map[Point]bool, int) {
	if grid[start] == 9 {
		heads[start] = true
		return heads, numberOfPaths + 1
	}
	neighbours := getValidNeighbours(grid, start)
	if len(neighbours) == 0 {
		return heads, numberOfPaths
	}

	for _, n := range neighbours {
		heads, numberOfPaths = findPaths(grid, n, heads, numberOfPaths)
	}
	return heads, numberOfPaths
}
