package day13

import (
	"fmt"
	"lcatania/aoc-2024-go/utils"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func Day13() string {
	fileContent := utils.ReadFile("./day13/input.txt")
	equations := parseInput(strings.Split(fileContent, "\n"))

	totalTickets := int64(0)

	// Now we have equations, use cramer's rule to determine X and Y (aka A btn presses and B btn presses)
	for _, eqn := range equations {
		countA := float64(((eqn.cOne * eqn.bTwo) - (eqn.bOne * eqn.cTwo))) / float64(((eqn.aOne * eqn.bTwo) - (eqn.bOne * eqn.aTwo)))

		if _, frac := math.Modf(countA); frac != 0 {
			continue
		}
		countB := float64((eqn.cOne - eqn.aOne*int64(countA)) / eqn.bOne)

		if _, frac := math.Modf(countB); frac != 0 {
			continue
		}

		if countA > 0 && countA < 100 && countB > 0 && countB < 100 {
			totalTickets += int64(countA*3 + countB)
		}
	}

	return strconv.FormatInt(totalTickets, 10)
}

func Day13Part2() string {
	fileContent := utils.ReadFile("./day13/input.txt")
	equations := parseInput(strings.Split(fileContent, "\n"))

	totalTickets := int64(0)

	// Now we have equations, use cramer's rule to determine X and Y (aka A btn presses and B btn presses)
	for _, eqn := range equations {
		eqn.cOne += 10000000000000
		eqn.cTwo += 10000000000000
		countA := float64(((eqn.cOne * eqn.bTwo) - (eqn.bOne * eqn.cTwo))) / float64(((eqn.aOne * eqn.bTwo) - (eqn.bOne * eqn.aTwo)))

		if _, frac := math.Modf(countA); frac != 0 {
			continue
		}
		countB := float64((eqn.cOne - eqn.aOne*int64(countA))) / float64(eqn.bOne)

		if _, frac := math.Modf(countB); frac != 0 {
			continue
		}

		if countA > 0 && countB > 0 {
			totalTickets += int64(countA*3 + countB)
		}
	}

	return strconv.FormatInt(totalTickets, 10)
}

type equationsStruct struct {
	aOne int64
	bOne int64
	cOne int64
	aTwo int64
	bTwo int64
	cTwo int64
}

func parseInput(contents []string) []equationsStruct {
	equations := make([]equationsStruct, 0)
	currentIndex := 0

	for _, line := range contents {
		if len(line) == 0 {
			currentIndex++
		} else if line[:6] == "Button" {
			if len(equations) == currentIndex {
				// First time at this index so increment slice
				equations = append(equations, equationsStruct{})
			}

			entries := regexp.MustCompile(`^Button (\w): X\+(\d+), Y\+(\d+)`)
			match := entries.FindStringSubmatch(line)

			a, _ := strconv.ParseInt(match[2], 10, 64)
			b, _ := strconv.ParseInt(match[3], 10, 64)

			if match[1] == "A" {
				equations[currentIndex].aOne = a
				equations[currentIndex].aTwo = b
			} else {
				equations[currentIndex].bOne = a
				equations[currentIndex].bTwo = b
			}

		} else if line[:5] == "Prize" {
			var prizeX, prizeY int64
			fmt.Sscanf(line, "Prize: X=%d, Y=%d", &prizeX, &prizeY)

			equations[currentIndex].cOne = prizeX
			equations[currentIndex].cTwo = prizeY
		}
	}

	return equations
}
