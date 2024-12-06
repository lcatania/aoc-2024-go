package day5

import (
	"lcatania/aoc-2024-go/utils"
	"slices"
	"strconv"
	"strings"
)

func Day5() int {
	fileContent := utils.ReadFile("./day5/input.txt")
	splittedInput := strings.Split(fileContent, "\n\n")

	rules := getRules(splittedInput[0])
	updates := getUpdates(splittedInput[1])

	sum := 0
	for _, u := range updates {
		if slices.IsSortedFunc(u, compareFn(rules)) {
			sum += u[len(u)/2]
		}
	}

	return sum
}

func Day5Part2() int {
	fileContent := utils.ReadFile("./day5/input.txt")
	splittedInput := strings.Split(fileContent, "\n\n")

	rules := getRules(splittedInput[0])
	updates := getUpdates(splittedInput[1])

	sum := 0
	for _, u := range updates {
		before := slices.Clone(u)
		slices.SortStableFunc(u, compareFn(rules))

		if slices.Compare(before, u) != 0 {
			sum += u[len(u)/2]
		}
	}

	return sum
}

func compareFn(rules map[int][]int) func(a, b int) int {
	return func(a, b int) int {
		aFirst := slices.Contains(rules[a], b)
		bFirst := slices.Contains(rules[b], a)

		if !aFirst && !bFirst {
			return 0
		}

		if aFirst {
			return -1
		}

		return 1
	}
}

func getUpdates(rawUpdates string) [][]int {
	lines := strings.Split(rawUpdates, "\n")
	result := [][]int{}
	for _, line := range lines {
		rule := utils.ConvertStringArrayToInt(strings.Split(line, ","))
		result = append(result, rule)
	}
	return result
}

func getRules(rawRules string) map[int][]int {
	lines := strings.Split(rawRules, "\n")

	rules := map[int][]int{}
	for _, l := range lines {
		v := strings.Split(l, "|")
		p1, _ := strconv.Atoi(v[0])
		p2, _ := strconv.Atoi(v[1])
		rules[p1] = append(rules[p1], p2)
	}

	return rules
}
