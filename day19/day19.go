package day19

import (
	"lcatania/aoc-2024-go/utils"
	"strconv"
	"strings"
)

func Day19() string {
	fileContent := utils.ReadFile("./day19/input.txt")
	lines := strings.Split(fileContent, "\n")

	towels, designs := parseInput(lines)

	possibleDesigns, _ := findPossibleDesigns(towels, designs)

	return strconv.Itoa(possibleDesigns)
}

func Day19Part2() string {
	fileContent := utils.ReadFile("./day19/input.txt")
	lines := strings.Split(fileContent, "\n")

	towels, designs := parseInput(lines)

	_, allPatterns := findPossibleDesigns(towels, designs)

	return strconv.Itoa(allPatterns)
}

func parseInput(input []string) ([]string, []string) {
	towels := make([]string, 0)
	towelsLine := input[0]
	towelsStrings := strings.Split(towelsLine, ", ")
	towels = append(towels, towelsStrings...)

	// parse designs
	designs := make([]string, 0)

	// parse designs
	designs = append(designs, input[2:]...)

	return towels, designs
}

func findPossibleDesigns(towels []string, designs []string) (int, int) {
	possibleDesigns := make([]string, 0)
	allPatterns := 0
	for _, design := range designs {
		if isPossible, patterns := isPossibleDesign(towels, design); isPossible {
			possibleDesigns = append(possibleDesigns, design)
			allPatterns += patterns
		}
	}

	return len(possibleDesigns), allPatterns
}

func isPossibleDesign(towels []string, design string) (bool, int) {
	dp := make([]int, len(design)+1)

	dp[0] = 1

	for i := 1; i <= len(design); i++ {

		for _, towel := range towels {
			if i >= len(towel) &&
				dp[i-len(towel)] > 0 &&
				design[i-len(towel):i] == towel {

				dp[i] += dp[i-len(towel)]
			}
		}
	}

	return dp[len(design)] > 0, dp[len(design)]
}
