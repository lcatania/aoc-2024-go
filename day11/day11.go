package day11

import (
	"fmt"
	"lcatania/aoc-2024-go/utils"
	"log"
	"slices"
	"strconv"
	"strings"
)

func Day11() int {
	input := utils.ReadFile("./day11/input.txt")
	list := strings.Split(input, " ")
	for i := 0; i < 25; i++ {
		for j := 0; j < len(list); j++ {
			if list[j] == "0" {
				list[j] = "1"
			} else if len(list[j])%2 == 0 {
				leftNumber := list[j][0 : len(list[j])/2]
				rightNumber := list[j][len(list[j])/2:]
				convertedNumber, err := strconv.Atoi(rightNumber)
				if err != nil {
					log.Fatal("Could not convert to number")
				}
				list[j] = leftNumber
				list = slices.Insert(list, j+1, strconv.Itoa(convertedNumber))
				j++
			} else {
				number, err := strconv.Atoi(list[j])
				if err != nil {
					log.Fatal("Could not convert to number")
				}
				list[j] = strconv.Itoa(number * 2024)
			}
		}
	}
	return len(list)
}

func Day11Part2() int {
	input := utils.ReadFile("./day11/input.txt")
	list := utils.ConvertStringArrayToInt(strings.Split(input, " "))
	fmt.Println(list)
	sum := 0
	cache := make(map[int]map[int]int)
	for i := 0; i <= 75; i++ {
		cache[i] = make(map[int]int)
	}

	for _, stone := range list {
		sum += countStones(stone, cache, 75)
	}
	return sum
}

func countStones(stone int, cache map[int]map[int]int, reps int) int {
	if reps == 0 {
		return 1
	} else {
		if v, ok := cache[reps][stone]; ok {
			return v
		} else {
			if stone == 0 {
				count := countStones(1, cache, reps-1)
				cache[reps][stone] = count
				return count
			} else if len(strconv.Itoa(stone))%2 == 0 {
				stoneString := strconv.Itoa(stone)
				newStone, _ := strconv.Atoi(stoneString[len(stoneString)/2:])
				newStone2, _ := strconv.Atoi(stoneString[:len(stoneString)/2])
				count := countStones(newStone, cache, reps-1) + countStones(newStone2, cache, reps-1)
				cache[reps][stone] = count
				return count
			} else {
				count := countStones(stone*2024, cache, reps-1)
				cache[reps][stone] = count
				return count
			}
		}
	}
}
