package day3

import (
	"lcatania/aoc-2024-go/utils"
	"regexp"
	"strconv"
	"strings"
)

func Day3() string {
	fileContent := utils.ReadFile("./day3/input.txt")

	return strconv.Itoa(getSum(fileContent))
}

func Day3Part2() string {
	fileContent := utils.ReadFile("./day3/input.txt")

	dontSplit := strings.Split(fileContent, "don't()")

	doSplits := make([]string, 0)
	for idx, val := range dontSplit {
		if idx == 0 {
			doSplits = append(doSplits, val)
			continue
		}
		doSplit := strings.Split(val, "do()")
		doSplits = append(doSplits, doSplit[1:]...)
	}

	sum := getSum(strings.Join(doSplits, ""))
	return strconv.Itoa(sum)
}

func getSum(valStr string) int {
	var regexPattern = regexp.MustCompile("mul\\((\\d{1,3},\\d{1,3})\\)")

	multParses := regexPattern.FindAllStringSubmatch(valStr, -1)
	multList := make([][]int, len(multParses))
	for i := range multParses {
		multLine := strings.Split(multParses[i][1], ",")
		v1, _ := strconv.Atoi(multLine[0])
		v2, _ := strconv.Atoi(multLine[1])
		multList[i] = []int{v1, v2}
	}

	return utils.Reduce(multList, func(acc int, b []int) int { return acc + (b[0] * b[1]) }, 0)
}
