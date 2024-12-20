package day1

import (
	"lcatania/aoc-2024-go/utils"
	"log"
	"slices"
	"strconv"
	"strings"
)

func Day1() string {
	fileContent := utils.ReadFile("./day1/input.txt")
	lines := strings.Split(fileContent, "\n")
	firstList := make([]int, len(lines))
	secondList := make([]int, len(lines))

	for i := 0; i < len(lines); i++ {
		values := strings.Split(lines[i], "   ")
		firstValue, err := strconv.Atoi(values[0])
		if err != nil {
			log.Fatal(err)
		}
		secondValue, err := strconv.Atoi(values[1])
		if err != nil {
			log.Fatal(err)
		}
		firstList[i] = firstValue
		secondList[i] = secondValue
	}
	slices.Sort(firstList)
	slices.Sort(secondList)

	distance := 0

	for i := 0; i < len(lines); i++ {
		if firstList[i] > secondList[i] {
			distance += firstList[i] - secondList[i]
		} else {
			distance += secondList[i] - firstList[i]
		}
	}
	return strconv.Itoa(distance)
}

func Day1Part2() string {
	fileContent := utils.ReadFile("./day1/input.txt")
	lines := strings.Split(fileContent, "\n")
	firstList := make([]int, len(lines))
	secondList := make([]int, len(lines))

	for i := 0; i < len(lines); i++ {
		values := strings.Split(lines[i], "   ")
		firstValue, err := strconv.Atoi(values[0])
		if err != nil {
			log.Fatal(err)
		}
		secondValue, err := strconv.Atoi(values[1])
		if err != nil {
			log.Fatal(err)
		}
		firstList[i] = firstValue
		secondList[i] = secondValue
	}

	distance := 0

	for _, s := range firstList {
		distance += s * utils.Count(secondList, func(x int) bool { return x == s })
	}
	return strconv.Itoa(distance)
}
