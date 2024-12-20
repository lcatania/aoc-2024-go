package day16

import (
	"fmt"
	"lcatania/aoc-2024-go/utils"
	"math"
	"slices"
	"strconv"
	"strings"
)

type node struct {
	x         int
	y         int
	direction rune
	cost      int
}

func newNode(x, y int, direction rune, cost int) *node {
	return &node{x: x, y: y, direction: direction, cost: cost}
}

func (n *node) String() string {
	return fmt.Sprintf("(%d,%d,%s)", n.x, n.y, string(n.direction))
}

type priorityQueue struct {
	data []*node
}

func newPriorityQueue() *priorityQueue {
	return &priorityQueue{data: []*node{}}
}

func (pq *priorityQueue) Append(n *node) {
	pq.data = append(pq.data, n)
	pq.Sort()
}

func (pq *priorityQueue) Pop() *node {
	outNode := pq.data[0]
	pq.data = pq.data[1:]
	return outNode
}

func (pq *priorityQueue) Sort() {
	slices.SortFunc(pq.data, func(n1, n2 *node) int {
		if n1.cost == n2.cost {
			return 0
		}
		if n1.cost < n2.cost {
			return -1
		}
		return 1
	})
}

func readInput() [][]rune {
	var data [][]rune
	fileContent := utils.ReadFile("./day16/input.txt")
	lines := strings.Split(fileContent, "\n")
	for _, line := range lines {
		data = append(data, []rune(line))
	}
	return data
}

func findStart(data [][]rune) (int, int) {
	for i := range data {
		for j := range data[i] {
			if data[i][j] == rune('S') {
				return i, j
			}
		}
	}
	return -1, -1
}

func findEnd(data [][]rune) (int, int) {
	for i := range data {
		for j := range data[i] {
			if data[i][j] == rune('E') {
				return i, j
			}
		}
	}
	return -1, -1
}

func findShortest(data [][]rune, startX, startY, endX, endY int, startDir rune) (int, map[rune][][]int) {
	pq := newPriorityQueue()
	pq.Append(newNode(startX, startY, startDir, 0))
	costMatrix := make(map[rune][][]int)
	costMatrix[rune('^')] = make([][]int, len(data))
	costMatrix[rune('v')] = make([][]int, len(data))
	costMatrix[rune('>')] = make([][]int, len(data))
	costMatrix[rune('<')] = make([][]int, len(data))
	for i := range data {
		costMatrix[rune('^')][i] = make([]int, len(data[i]))
		costMatrix[rune('v')][i] = make([]int, len(data[i]))
		costMatrix[rune('>')][i] = make([]int, len(data[i]))
		costMatrix[rune('<')][i] = make([]int, len(data[i]))
		for j := range data[i] {
			costMatrix[rune('^')][i][j] = math.MaxInt32
			costMatrix[rune('v')][i][j] = math.MaxInt32
			costMatrix[rune('>')][i][j] = math.MaxInt32
			costMatrix[rune('<')][i][j] = math.MaxInt32
		}
	}
	costMatrix[rune('>')][startX][startY] = 0
	for len(pq.data) > 0 {
		currNode := pq.Pop()
		if currNode.x == endX && currNode.y == endY {
			return currNode.cost, costMatrix
		}
		// Check straight
		dx, dy := translateDirection(currNode.direction)
		newX := currNode.x + dx
		newY := currNode.y + dy
		if data[newX][newY] != rune('#') {
			newCost := currNode.cost + 1
			if newCost < costMatrix[currNode.direction][newX][newY] {
				costMatrix[currNode.direction][newX][newY] = newCost
				pq.Append(newNode(newX, newY, currNode.direction, newCost))
			}
		}
		// Check right
		newDirection := turnRight(currNode.direction)
		dx, dy = translateDirection(newDirection)
		newX = currNode.x + dx
		newY = currNode.y + dy
		if data[newX][newY] != rune('#') {
			newCost := currNode.cost + 1001
			if newCost < costMatrix[newDirection][newX][newY] {
				costMatrix[newDirection][currNode.x][currNode.y] = newCost - 1
				costMatrix[newDirection][newX][newY] = newCost
				pq.Append(newNode(newX, newY, newDirection, newCost))
			}
		}
		// Check left
		newDirection = turnLeft(currNode.direction)
		dx, dy = translateDirection(newDirection)
		newX = currNode.x + dx
		newY = currNode.y + dy
		if data[newX][newY] != rune('#') {
			newCost := currNode.cost + 1001
			if newCost < costMatrix[newDirection][newX][newY] {
				costMatrix[newDirection][currNode.x][currNode.y] = newCost - 1
				costMatrix[newDirection][newX][newY] = newCost
				pq.Append(newNode(newX, newY, newDirection, newCost))
			}
		}
	}
	return -1, costMatrix
}

func Day16() string {
	data := readInput()
	result := 0
	endX, endY := findEnd(data)
	startX, startY := findStart(data)
	result, _ = findShortest(data, startX, startY, endX, endY, rune('>'))
	return strconv.Itoa(result)
}

func Day16Part2() string {
	data := readInput()
	result := 0
	endX, endY := findEnd(data)
	startX, startY := findStart(data)
	shortestPath, costMatrix1 := findShortest(data, startX, startY, endX, endY, rune('>'))
	_, costMatrix2 := findShortest(data, endX, endY, startX, startY, rune('^'))
	_, costMatrix3 := findShortest(data, endX, endY, startX, startY, rune('>'))
	_, costMatrix4 := findShortest(data, endX, endY, startX, startY, rune('v'))
	_, costMatrix5 := findShortest(data, endX, endY, startX, startY, rune('<'))
	for d := range costMatrix1 {
		flippedDir := turnRight(turnRight(d))
		for i := range costMatrix1[d] {
			for j := range costMatrix1[d][i] {
				for _, v := range []int{costMatrix2[flippedDir][i][j],
					costMatrix3[flippedDir][i][j],
					costMatrix4[flippedDir][i][j],
					costMatrix5[flippedDir][i][j]} {
					flipSum := costMatrix1[d][i][j] + v
					if flipSum == shortestPath {
						data[i][j] = rune('O')
					}
				}
			}
		}
	}
	for i := range data {
		for j := range data[i] {
			if data[i][j] == rune('O') || data[i][j] == rune('S') || data[i][j] == rune('E') {
				result++
			}
		}
	}
	return strconv.Itoa(result)
}

func turnRight(direction rune) rune {
	switch direction {
	case rune('^'):
		return rune('>')
	case rune('>'):
		return rune('v')
	case rune('v'):
		return rune('<')
	case rune('<'):
		return rune('^')
	default:
		fmt.Println("Wrong direction provided")
		return direction
	}
}

func translateDirection(direction rune) (int, int) {
	switch direction {
	case rune('^'):
		return -1, 0
	case rune('<'):
		return 0, -1
	case rune('v'):
		return 1, 0
	case rune('>'):
		return 0, 1
	default:
		fmt.Println("Wrong direction provided")
		return 0, 0
	}
}

func turnLeft(direction rune) rune {
	switch direction {
	case rune('^'):
		return rune('<')
	case rune('<'):
		return rune('v')
	case rune('v'):
		return rune('>')
	case rune('>'):
		return rune('^')
	default:
		fmt.Println("Wrong direction provided")
		return direction
	}
}
