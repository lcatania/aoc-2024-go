package day18

import (
	"fmt"
	"lcatania/aoc-2024-go/utils"
	"math"
	"strconv"
	"strings"
)

type Coord struct {
	x, y int
}

func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic("strconv")
	}

	return i
}

func AStar(start, goal Coord, x, y int, grid *[71][71]bool) (int, error) {
	openSet := map[Coord]bool{start: true}
	gScore := map[Coord]int{start: 0}
	fScore := map[Coord]int{start: manhattanDistance(start, goal)}
	distFunc := manhattanDistance

	for len(openSet) > 0 {
		var current Coord
		minK := math.MaxInt
		for k := range openSet {
			if s, ok := fScore[k]; ok && s < minK {
				minK = fScore[k]
				current = k
			}
		}
		if current.x == goal.x && current.y == goal.y {
			return gScore[current], nil
		}
		delete(openSet, current)
		for _, nb := range neighbors(current) {
			if nb.x < 0 || nb.y < 0 || nb.x > x || nb.y > y {
				continue
			}
			if (*grid)[nb.x][nb.y] {
				continue
			}

			gScoreTentative := gScore[current] + distFunc(current, nb)
			if gScoreCurrent, ok := gScore[nb]; !ok || gScoreTentative < gScoreCurrent {
				gScore[nb] = gScoreTentative
				fScore[nb] = gScoreTentative + manhattanDistance(nb, goal)
				if _, ok := openSet[nb]; !ok {
					openSet[nb] = true
				}
			}
		}
	}
	return 0, fmt.Errorf("Couldn't find the path")
}

func neighbors(c Coord) []Coord {
	return []Coord{
		c.plus(Coord{-1, 0}),
		c.plus(Coord{0, -1}),
		c.plus(Coord{0, 1}),
		c.plus(Coord{1, 0}),
	}
}

func (c Coord) plus(o Coord) Coord {
	return Coord{c.x + o.x, c.y + o.y}
}

func manhattanDistance(p, q Coord) int {
	return abs(p.x-q.x) + abs(p.y-q.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Day18() string {
	fileContent := utils.ReadFile("./day18/input.txt")
	lines := strings.Split(fileContent, "\n")

	var positions []Coord
	for _, line := range lines {
		l := strings.SplitN(line, ",", 2)
		positions = append(positions, Coord{MustAtoi(l[1]), MustAtoi(l[0])})
	}

	var grid [71][71]bool
	for i := 0; i < 1024; i++ {
		p := positions[i]
		grid[p.x][p.y] = true
	}

	s, _ := AStar(Coord{0, 0}, Coord{70, 70}, 70, 70, &grid)

	return strconv.Itoa(s)
}

func Day18Part2() string {
	fileContent := utils.ReadFile("./day18/input.txt")
	lines := strings.Split(fileContent, "\n")
	var positions []Coord
	for _, line := range lines {
		l := strings.SplitN(line, ",", 2)
		positions = append(positions, Coord{MustAtoi(l[1]), MustAtoi(l[0])})
	}

	var grid [71][71]bool
	for i := 0; i < 2048; i++ {
		p := positions[i]
		grid[p.x][p.y] = true
	}
	for i := 2048; i < len(positions); i++ {
		p := positions[i]
		grid[p.x][p.y] = true

		_, err := AStar(Coord{0, 0}, Coord{70, 70}, 70, 70, &grid)
		if err != nil {
			return fmt.Sprintf("%d,%d", p.y, p.x)
		}
	}

	return "Not found"
}
