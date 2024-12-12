package day12

import (
	"lcatania/aoc-2024-go/utils"
	"strings"
)

type Point struct {
	x int
	y int
}

func Day12() int {
	fileContent := utils.ReadFile("./day12/input.txt")
	lines := strings.Split(fileContent, "\n")
	grid := map[Point]rune{}

	maxPoint := Point{x: 0, y: 0}
	for y, line := range lines {
		for x, c := range line {
			grid[Point{x: x, y: y}] = c
		}
	}
	maxPoint.x = len(lines[0])
	maxPoint.y = len(lines)

	dirs := []Point{
		{x: 0, y: 1},
		{x: 0, y: -1},
		{x: 1, y: 0},
		{x: -1, y: 0},
	}
	result := 0
	complete := map[Point]bool{}
	for i := 0; i < maxPoint.x; i++ {
		for j := 0; j < maxPoint.y; j++ {
			basePoint := Point{x: i, y: j}
			baseRune := grid[basePoint]
			if complete[basePoint] {
				continue
			}

			blob := map[Point]bool{basePoint: true}
			touching := map[Point]map[Point]bool{}
			anyFound := true
			for anyFound {
				anyFound = false
				for p := range blob {
					for _, d := range dirs {
						pNew := Point{p.x + d.x, p.y + d.y}
						if grid[pNew] != baseRune {
							if touching[pNew] == nil {
								touching[pNew] = map[Point]bool{}
							}
							touching[pNew][p] = true
							continue
						}
						if blob[pNew] {
							continue
						}
						anyFound = true
						blob[pNew] = true
					}
				}
			}
			area := len(blob)
			perimiter := 0
			for _, v := range touching {
				perimiter += len(v)
			}
			result += perimiter * area

			for p := range blob {
				complete[p] = true
			}
		}
	}
	return result
}

func Day12Part2() int {
	fileContent := utils.ReadFile("./day12/input.txt")
	lines := strings.Split(fileContent, "\n")
	grid := map[Point]rune{}

	maxPoint := Point{x: 0, y: 0}
	for y, line := range lines {
		for x, c := range line {
			grid[Point{x: x, y: y}] = c
		}
	}
	maxPoint.x = len(lines[0])
	maxPoint.y = len(lines)

	dirs := []Point{
		{x: 0, y: 1},
		{x: 0, y: -1},
		{x: 1, y: 0},
		{x: -1, y: 0},
	}
	result := 0
	complete := map[Point]bool{}
	for i := 0; i < maxPoint.x; i++ {
		for j := 0; j < maxPoint.y; j++ {
			basePoint := Point{x: i, y: j}
			baseRune := grid[basePoint]
			if complete[basePoint] {
				continue
			}

			blob := map[Point]bool{basePoint: true}
			touching := map[Point]map[Point]bool{}
			anyFound := true
			for anyFound {
				anyFound = false
				for p := range blob {
					for _, d := range dirs {
						pNew := Point{p.x + d.x, p.y + d.y}
						if grid[pNew] != baseRune {
							if touching[pNew] == nil {
								touching[pNew] = map[Point]bool{}
							}
							touching[pNew][p] = true
							continue
						}
						if blob[pNew] {
							continue
						}
						anyFound = true
						blob[pNew] = true
					}
				}
			}
			area := len(blob)

			for p := range blob {
				complete[p] = true
			}

			sides := 0
			sidesToWalk := touching
			for len(sidesToWalk) > 0 {
				var pOut Point
				var pIn Point
				for po, pi := range sidesToWalk {
					pOut = po
					if len(pi) == 0 {
						panic("Should never happen")
					}
					for p := range pi {
						pIn = p
						break
					}
					break
				}

				walk(pIn, pOut, Point{x: -1, y: 0}, &sidesToWalk)
				walk(pIn, pOut, Point{x: 1, y: 0}, &sidesToWalk)
				walk(pIn, pOut, Point{x: 0, y: 1}, &sidesToWalk)
				walk(pIn, pOut, Point{x: 0, y: -1}, &sidesToWalk)
				delete(sidesToWalk[pOut], pIn)
				if len(sidesToWalk[pOut]) == 0 {
					delete(sidesToWalk, pOut)
				}

				sides++
			}
			result += sides * area
		}
	}
	return result
}

func walk(walkIn Point, walkOut Point, dir Point, sidesToWalk *map[Point]map[Point]bool) {
	for {
		walkOutNew := Point{x: walkOut.x + dir.x, y: walkOut.y + dir.y}
		walkInNew := Point{x: walkIn.x + dir.x, y: walkIn.y + dir.y}
		if len((*sidesToWalk)[walkOutNew]) == 0 {
			break
		}
		if !(*sidesToWalk)[walkOutNew][walkInNew] {
			break
		}
		walkOut = walkOutNew
		walkIn = walkInNew
		delete((*sidesToWalk)[walkOutNew], walkInNew)
		if len((*sidesToWalk)[walkOutNew]) == 0 {
			delete((*sidesToWalk), walkOutNew)
		}
	}
}
