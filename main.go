package main

import (
	"flag"
	"fmt"
	"lcatania/aoc-2024-go/day1"
	"lcatania/aoc-2024-go/day10"
	"lcatania/aoc-2024-go/day11"
	"lcatania/aoc-2024-go/day12"
	"lcatania/aoc-2024-go/day13"
	"lcatania/aoc-2024-go/day14"
	"lcatania/aoc-2024-go/day15"
	"lcatania/aoc-2024-go/day16"
	"lcatania/aoc-2024-go/day17"
	"lcatania/aoc-2024-go/day18"
	"lcatania/aoc-2024-go/day19"
	"lcatania/aoc-2024-go/day2"
	"lcatania/aoc-2024-go/day20"
	"lcatania/aoc-2024-go/day3"
	"lcatania/aoc-2024-go/day4"
	"lcatania/aoc-2024-go/day5"
	"lcatania/aoc-2024-go/day6"
	"lcatania/aoc-2024-go/day7"
	"lcatania/aoc-2024-go/day8"
	"lcatania/aoc-2024-go/day9"
	"strconv"
)

var dayFuncMapping = map[string](func() string){
	"1":    day1.Day1,
	"1.5":  day1.Day1Part2,
	"2":    day2.Day2,
	"2.5":  day2.Day2Part2,
	"3":    day3.Day3,
	"3.5":  day3.Day3Part2,
	"4":    day4.Day4,
	"4.5":  day4.Day4Part2,
	"5":    day5.Day5,
	"5.5":  day5.Day5Part2,
	"6":    day6.Day6,
	"6.5":  day6.Day6Part2,
	"7":    day7.Day7,
	"7.5":  day7.Day7Part2,
	"8":    day8.Day8,
	"8.5":  day8.Day8Part2,
	"9":    day9.Day9,
	"9.5":  day9.Day9Part2,
	"10":   day10.Day10,
	"10.5": day10.Day10Part2,
	"11":   day11.Day11,
	"11.5": day11.Day11Part2,
	"12":   day12.Day12,
	"12.5": day12.Day12Part2,
	"13":   day13.Day13,
	"13.5": day13.Day13Part2,
	"14":   day14.Day14,
	"14.5": day14.Day14Part2,
	"15":   day15.Day15,
	"15.5": day15.Day15Part2,
	"16":   day16.Day16,
	"16.5": day16.Day16Part2,
	"17":   day17.Day17,
	"17.5": day17.Day17Part2,
	"18":   day18.Day18,
	"18.5": day18.Day18Part2,
	"19":   day19.Day19,
	"19.5": day19.Day19Part2,
	"20":   day20.Day20,
	"20.5": day20.Day20Part2,
}

func main() {
	dayToRun := flag.Int("day", -1, "Day to run")
	flag.Parse()
	maxDayToRun := len(dayFuncMapping) / 2
	if *dayToRun > maxDayToRun {
		dayToRun = &maxDayToRun
	}
	if *dayToRun == -1 {
		fmt.Println("-----------------------------")
		for i := 1; i <= maxDayToRun; i++ {
			day := strconv.Itoa(i)
			dayPart2 := day + ".5"
			fmt.Println(fmt.Sprintf("Day%s Result -", day), dayFuncMapping[day]())
			fmt.Println(fmt.Sprintf("Day%s Part2 Result -", day), dayFuncMapping[dayPart2]())
			fmt.Println("-----------------------------")
		}
	} else {
		fmt.Println("-----------------------------")
		day := strconv.Itoa(*dayToRun)
		dayPart2 := day + ".5"
		fmt.Println(fmt.Sprintf("Day%s Result -", day), dayFuncMapping[day]())
		fmt.Println(fmt.Sprintf("Day%s Part2 Result -", day), dayFuncMapping[dayPart2]())
		fmt.Println("-----------------------------")
	}

}
