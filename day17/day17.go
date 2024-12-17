package day17

import (
	"fmt"
	"lcatania/aoc-2024-go/utils"
	"strconv"
	"strings"
)

func Day17() int {
	fileContent := utils.ReadFile("./day17/sample_input.txt")
	registers, instructions := parseInput(strings.Split(fileContent, "\n"))

	fmt.Println(registers, instructions)
	return 0
}

func Day17Part2() int {
	fileContent := utils.ReadFile("./day17/sample_input.txt")
	registers, instructions := parseInput(strings.Split(fileContent, "\n"))

	fmt.Println(registers, instructions)
	return 0
}

func parseInput(contents []string) (map[byte]int64, []int64) {
	registers := make(map[byte]int64)
	instructions := make([]int64, 0)

	for _, line := range contents {
		if len(line) == 0 {
			continue
		}

		if line[:8] == "Register" {
			register := line[10]
			val := line[12:]

			intVal, _ := strconv.ParseInt(val, 10, 64)

			registers[register] = intVal
		}

		if line[:7] == "Program" {
			stringInstructions := strings.Split(line, ",")
			for _, val := range stringInstructions {
				intVal, _ := strconv.ParseInt(val, 10, 64)
				instructions = append(instructions, intVal)
			}
		}
	}

	return registers, instructions
}
