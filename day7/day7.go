package day7

import (
	"lcatania/aoc-2024-go/utils"
	"log"
	"math"
	"strconv"
	"strings"
)

func Day7() int {
	fileContent := utils.ReadFile("./day7/input.txt")
	lines := strings.Split(fileContent, "\n")
	sum := 0
	for _, line := range lines {
		parts := strings.Split(line, ":")
		result, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal("Int conversation failed")
		}

		equations := utils.ConvertStringArrayToInt(strings.Split(strings.TrimSpace(parts[1]), " "))

		solved := calc(result, equations[0], equations[1:], 1)
		if solved {
			sum += result
		}

	}
	return sum
}

func Day7Part2() int {
	fileContent := utils.ReadFile("./day7/input.txt")
	lines := strings.Split(fileContent, "\n")
	sum := 0
	for _, line := range lines {
		parts := strings.Split(line, ":")
		result, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal("Int conversation failed")
		}

		equations := utils.ConvertStringArrayToInt(strings.Split(strings.TrimSpace(parts[1]), " "))

		solved := calc(result, equations[0], equations[1:], 2)
		if solved {
			sum += result
		}

	}
	return sum
}

func calc(testValue, result int, numbers []int, part int) bool {
	if result > testValue {
		return false
	}

	if len(numbers) == 0 {
		return result == testValue
	}

	if calc(testValue, result+numbers[0], numbers[1:], part) {
		return true
	}

	if calc(testValue, result*numbers[0], numbers[1:], part) {
		return true
	}

	if part == 1 {
		return false
	}

	shift := int(math.Floor(math.Log10(float64(numbers[0])))) + 1
	concat := result*int(math.Pow10(shift)) + numbers[0]
	return calc(testValue, concat, numbers[1:], part)
}
