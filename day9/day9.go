package day9

import (
	"lcatania/aoc-2024-go/utils"
	"strconv"
)

func Day9() string {
	var nums []int

	line := utils.ReadFile("./day9/input.txt")
	nums = make([]int, len(line))

	for i, r := range line {
		nums[i] = int(r - '0')

	}
	total := 0
	idx := 0

	left, right := 0, len(nums)-1
	if right%2 == 1 {
		right--
	}

	for left <= right {
		if left%2 == 0 {
			n := nums[left]
			blockId := left / 2
			total += (idx + idx + n - 1) * n / 2 * blockId
			idx += n
			nums[left] = 0
		} else {
			fill := nums[left]
			blockId := right / 2
			n := nums[right]

			if fill >= n {
				total += (idx + idx + n - 1) * n / 2 * blockId
				idx += n
				nums[left] -= n
				nums[right] -= n
			} else {
				total += (idx + idx + fill - 1) * fill / 2 * blockId
				idx += fill
				nums[left] -= fill
				nums[right] -= fill
			}
		}

		if nums[left] == 0 {
			left++
		}

		if right >= 0 && nums[right] == 0 {
			right -= 2
		}
	}

	return strconv.Itoa(total)
}

func Day9Part2() string {
	var nums []int

	line := utils.ReadFile("./day9/input.txt")
	nums = make([]int, len(line))

	for i, r := range line {
		nums[i] = int(r - '0')
	}

	total := 0
	right := len(nums) - 1
	if right%2 == 1 {
		right--
	}

	prefix := make([]int, len(nums)+1)

	for i, n := range nums {
		prefix[i+1] = prefix[i] + n
	}

	for right >= 0 {
		n := nums[right]
		moved := false
		for i := 1; i < right; i += 2 {
			if nums[i] >= 0 && nums[i] >= n {
				blockId := right / 2
				idx := prefix[i]
				total += (idx + idx + n - 1) * n / 2 * blockId
				nums[i] -= n
				prefix[i] += n
				moved = true
				break
			}
		}
		if !moved {
			blockId := right / 2

			idx := prefix[right]
			total += (idx + idx + n - 1) * n / 2 * blockId
		}
		right -= 2
	}

	return strconv.Itoa(total)
}
