package day6

import (
	"lcatania/aoc-2024-go/utils"
	"strconv"
	"strings"
)

type vec [2]int

func Day6() string {
	fileContent := utils.ReadFile("./day6/input.txt")
	lines := strings.Split(fileContent, "\n")
	w := len(lines[0])
	h := len(lines)

	grid := make([][]rune, h)
	for y, l := range lines {
		grid[y] = []rune(l)
	}
	start := findGuard(grid)

	visited := map[vec]bool{}
	guard := start
	dir := vec{0, -1}
	for {
		visited[guard] = true
		next := move(guard, dir)
		if !inBounds(next, w, h) {
			break
		}

		for grid[next[1]][next[0]] == '#' {
			dir.turn()
			next = move(guard, dir)
		}
		guard = next
	}
	return strconv.Itoa(len(visited))
}

func Day6Part2() string {
	fileContent := utils.ReadFile("./day6/input.txt")
	lines := strings.Split(fileContent, "\n")
	w := len(lines[0])
	h := len(lines)

	grid := make([][]rune, h)
	for y, l := range lines {
		grid[y] = []rune(l)
	}
	start := findGuard(grid)
	visited := map[vec]bool{}
	guard := start
	dir := vec{0, -1}
	for {
		visited[guard] = true
		next := move(guard, dir)
		if !inBounds(next, w, h) {
			break
		}

		for grid[next[1]][next[0]] == '#' {
			dir.turn()
			next = move(guard, dir)
		}
		guard = next
	}
	// result
	result := 0
	for v := range visited {
		if v == start {
			continue
		}

		grid[v[1]][v[0]] = '#'
		if detectLoop(grid, w, h, start) {
			result++
		}
		grid[v[1]][v[0]] = '.'
	}

	return strconv.Itoa(result)
}

func findGuard(grid [][]rune) vec {
	for y, l := range grid {
		for x, r := range l {
			if r == '^' {
				grid[y][x] = '.'
				return vec{x, y}
			}
		}
	}
	panic("no guard found")
}

func inBounds(pos vec, w, h int) bool {
	if pos[0] < 0 || pos[0] >= w {
		return false
	}
	if pos[1] < 0 || pos[1] >= h {
		return false
	}
	return true
}

func detectLoop(grid [][]rune, w, h int, start vec) bool {
	visited := map[vec]vec{}
	guard := start
	dir := vec{0, -1}
	for {
		if visited[guard] == dir {
			return true
		}

		visited[guard] = dir
		next := move(guard, dir)
		if !inBounds(next, w, h) {
			break
		}

		for grid[next[1]][next[0]] == '#' {
			dir.turn()
			next = move(guard, dir)
		}
		guard = next
	}
	return false
}

func move(pos, dir vec) vec {
	return vec{pos[0] + dir[0], pos[1] + dir[1]}
}

func (dir *vec) turn() {
	dir[0], dir[1] = -dir[1], dir[0]
}
