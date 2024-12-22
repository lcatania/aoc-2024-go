package day22

import (
	"lcatania/aoc-2024-go/utils"
	"math"
	"strconv"
	"strings"
)

func Day22() string {
	fileContent := utils.ReadFile("./day22/input.txt")
	lines := utils.ConvertStringArrayToInt(strings.Split(fileContent, "\n"))
	var sum int64 = 0
	for _, line := range lines {
		result := int64(line)
		for i := 0; i < 2000; i++ {
			result = step3(step2(step1(result)))
		}
		sum += result
	}
	return strconv.Itoa(int(sum))
}

func Day22Part2() string {
	fileContent := utils.ReadFile("./day22/input.txt")
	lines := utils.ConvertStringArrayToInt(strings.Split(fileContent, "\n"))

	b := make([][]bananaPrice, len(lines))
	for i, line := range lines {
		b[i] = make([]bananaPrice, 2000)
		num := int64(line)
		prev := int(num % 10)
		for j := range 2000 {
			num = int64(step3(step2(step1(num))))
			b[i][j] = getNumAndChange(num, prev)
			prev = int(num % 10)
		}
	}

	return strconv.Itoa(findMaxNumberOfBananas(b))
}

func step1(secret int64) int64 {
	result := secret * 64
	return prune(mix(secret, result))
}

func step2(secret int64) int64 {
	result := math.Floor(float64(secret) / 32)
	return prune(mix(secret, int64(result)))
}

func step3(secret int64) int64 {
	result := secret * 2048
	return prune(mix(secret, result))
}

func mix(secret int64, mixIn int64) int64 {
	return secret ^ mixIn
}

func prune(secret int64) int64 {
	return secret % 16777216
}

type bananaPrice struct {
	num    int
	change int
}

func getNumAndChange(input int64, previous int) bananaPrice {
	return bananaPrice{num: int(input % 10), change: int(input%10) - previous}
}

type seq struct {
	a, b, c, d int
}

func findMaxNumberOfBananas(b [][]bananaPrice) int {
	seqMap := make(map[seq][]int)
	for i, _ := range b {
		s := []int{b[i][0].change, b[i][1].change, b[i][2].change}
		for j := 3; j < len(b[i]); j++ {
			s = append(s, b[i][j].change)

			if _, ok := seqMap[seq{s[0], s[1], s[2], s[3]}]; !ok {
				seqMap[seq{s[0], s[1], s[2], s[3]}] = make([]int, len(b))
			}

			if seqMap[seq{s[0], s[1], s[2], s[3]}][i] == 0 {
				seqMap[seq{s[0], s[1], s[2], s[3]}][i] = b[i][j].num
			}

			s = s[1:]
		}
	}
	max := 0
	for _, n := range seqMap {
		sum := 0
		for _, r := range n {
			sum += r
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
