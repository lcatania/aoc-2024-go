package day24

import (
	"fmt"
	"lcatania/aoc-2024-go/utils"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Gate struct {
	in1, in2 string
	out      string
	operator string
	executed bool
}

func Day24() string {
	fileContent := utils.ReadFile("./day24/input.txt")
	content := strings.Split(fileContent, "\n\n")
	setup := strings.Split(content[0], "\n")
	commands := strings.Split(content[1], "\n")

	register := make(map[string]bool)
	var zs []string

	i := 0
	line := ""
	for _, line = range setup {

		var name string
		var value int
		fmt.Sscanf(line, "%s %d", &name, &value)
		name = name[:len(name)-1]
		register[name] = value == 1
	}

	gates := make([]Gate, 0)

	for _, line := range commands {
		var in1, in2, out, operator string
		fmt.Sscanf(line, "%s %s %s -> %s", &in1, &operator, &in2, &out)
		gates = append(gates, Gate{in1, in2, out, operator, false})

		if out[0] == 'z' && !slices.Contains(zs, out) {
			zs = append(zs, out)
		}
	}

	sort.Strings(zs)

	done := false
	for !done {
		done = true

		for _, gate := range gates {
			if gate.executed {
				continue
			}

			v1, found1 := register[gate.in1]
			v2, found2 := register[gate.in2]

			if !found1 || !found2 {
				//done = false
				continue
			}

			var value bool
			switch gate.operator {
			case "AND":
				value = v1 && v2
			case "OR":
				value = v1 || v2
			case "XOR":
				value = v1 != v2
			}

			register[gate.out] = value
			gate.executed = true
		}

		for _, z := range zs {
			_, found := register[z]
			if !found {
				done = false
				break
			}
		}
	}

	result := 0
	for i = len(zs) - 1; i >= 0; i-- {
		result = result << 1
		if register[zs[i]] {
			result++
		}
	}

	return strconv.Itoa(result)
}

func Day24Part2() string {
	return ""
}
