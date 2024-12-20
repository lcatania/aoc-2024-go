package day14

import (
	"fmt"
	"lcatania/aoc-2024-go/utils"
	"strconv"
	"strings"
)

type Bot struct {
	x  int
	y  int
	dx int
	dy int
}

const (
	maxX = 101
	maxY = 103
)

func Day14() string {
	var bots []Bot

	fileContent := utils.ReadFile("./day14/input.txt")
	lines := strings.Split(fileContent, "\n")

	for _, line := range lines {
		var b Bot

		_, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &b.x, &b.y, &b.dx, &b.dy)
		if err != nil {
			fmt.Println("Error parsing line:", err)
			continue
		}
		bots = append(bots, b)
	}

	for i := 0; i < 100; i++ {
		for j := 0; j < len(bots); j++ {
			bots[j].x += bots[j].dx
			bots[j].y += bots[j].dy

			// Wrap around the edges
			if bots[j].x < 0 {
				bots[j].x += maxX
			} else if bots[j].x >= maxX {
				bots[j].x -= maxX
			}

			if bots[j].y < 0 {
				bots[j].y += maxY
			} else if bots[j].y >= maxY {
				bots[j].y -= maxY
			}
		}
	}

	q1, q2, q3, q4 := 0, 0, 0, 0
	for _, bot := range bots {
		if bot.x < maxX/2 && bot.y < maxY/2 {
			q1++
		}
		if bot.x > maxX/2 && bot.y < maxY/2 {
			q2++
		}
		if bot.x < maxX/2 && bot.y > maxY/2 {
			q3++
		}
		if bot.x > maxX/2 && bot.y > maxY/2 {
			q4++
		}
	}

	out := q1 * q2 * q3 * q4
	return strconv.Itoa(out)
}

func Day14Part2() string {
	var bots []Bot

	fileContent := utils.ReadFile("./day14/input.txt")
	lines := strings.Split(fileContent, "\n")

	for _, line := range lines {
		var b Bot

		_, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &b.x, &b.y, &b.dx, &b.dy)
		if err != nil {
			fmt.Println("Error parsing line:", err)
			continue
		}
		bots = append(bots, b)
	}

	for i := 0; i < 1e5; i++ {

		var grid [maxY][maxX]rune
		for y := range grid {
			for x := range grid[y] {
				grid[y][x] = '.'
			}
		}

		distinct := true

		// Update bot positions and check for collisions
		for j := 0; j < len(bots); j++ {
			bots[j].x += bots[j].dx
			bots[j].y += bots[j].dy

			// Wrap around the edges
			if bots[j].x < 0 {
				bots[j].x += maxX
			} else if bots[j].x >= maxX {
				bots[j].x -= maxX
			}

			if bots[j].y < 0 {
				bots[j].y += maxY
			} else if bots[j].y >= maxY {
				bots[j].y -= maxY
			}

			// Mark the grid cell
			if grid[bots[j].y][bots[j].x] == '.' {
				grid[bots[j].y][bots[j].x] = '#'
			} else {
				distinct = false
			}
		}

		if distinct {
			return strconv.Itoa(i + 1)
		}
	}
	return "0"
}
