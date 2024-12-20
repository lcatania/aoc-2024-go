package day2

import (
	"lcatania/aoc-2024-go/utils"
	"math"
	"strconv"
	"strings"
)

func Day2() string {
	fileContent := utils.ReadFile("./day2/input.txt")
	lines := strings.Split(fileContent, "\n")
	numSafe := 0
	for _, line := range lines {

		values := utils.ConvertStringArrayToInt(strings.Split(line, " "))
		if isReportSafe(values) {
			numSafe += 1
		}
	}
	return strconv.Itoa(numSafe)
}

func Day2Part2() string {
	fileContent := utils.ReadFile("./day2/input.txt")
	lines := strings.Split(fileContent, "\n")
	numSafe := 0
	for _, line := range lines {

		values := utils.ConvertStringArrayToInt(strings.Split(line, " "))
		isSafe := isReportSafe(values)
		if !isSafe {
			for i := 0; i < len(values); i++ {
				newReport := []int{}
				newReport = append(newReport, values[:i]...)
				newReport = append(newReport, values[i+1:]...)

				if isReportSafe(newReport) {
					isSafe = true
					break
				}
			}
		}
		if isSafe {
			numSafe += 1
		}
	}
	return strconv.Itoa(numSafe)
}

func isReportSafe(values []int) bool {
	if len(values) < 2 {
		return true // Only one value is always safe
	}
	increasing := values[1] > values[0]
	for i := 1; i < len(values); i++ {
		val := values[i]
		if val == values[i-1] {
			// Value doesn't increase or decrease, automatically unsafe
			return false
		}
		if increasing && val < values[i-1] {
			// Increasing, but we found a value that doesn't increase
			return false
		}
		if !increasing && val > values[i-1] {
			// Decreasing, but we found a value that doesn't decrease
			return false
		}

		safeChange := float64(3)
		diff := math.Abs(float64(val) - float64(values[i-1]))
		if diff > safeChange {
			// Change is too large to be considered safe
			return false
		}
	}

	return true
}
