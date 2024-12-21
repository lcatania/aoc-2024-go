package day21

import (
	"fmt"
	"lcatania/aoc-2024-go/utils"
	"log"
	"sort"
	"strconv"
	"strings"
)

func Day21() string {
	fileContent := utils.ReadFile("./day21/input.txt")
	lines := strings.Split(fileContent, "\n")
	if len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}
	sum := getSum(lines, 2)
	return strconv.Itoa(sum)
}

func Day21Part2() string {
	fileContent := utils.ReadFile("./day21/input.txt")
	lines := strings.Split(fileContent, "\n")
	if len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}
	sum := getSum(lines, 25)
	return strconv.Itoa(sum)
}

type Point struct {
	X int
	Y int
}

type Keypad map[rune]Point

func (k Keypad) String() string {
	var runes []rune
	for c := range k {
		runes = append(runes, c)
	}
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

/*
+---+---+---+
| 7 | 8 | 9 | 0
+---+---+---+
| 4 | 5 | 6 | 1
+---+---+---+
| 1 | 2 | 3 | 2
+---+---+---+
. X | 0 | A | 3
.   +---+---+
*/
var numKeypad = Keypad{
	'7': {0, 0},
	'8': {1, 0},
	'9': {2, 0},
	'4': {0, 1},
	'5': {1, 1},
	'6': {2, 1},
	'1': {0, 2},
	'2': {1, 2},
	'3': {2, 2},
	'X': {0, 3},
	'0': {1, 3},
	'A': {2, 3},
}

/*
.   +---+---+
. X | ^ | A | 0
+---+---+---+
| < | v | > | 1
+---+---+---+
*/
var dirKeypad = Keypad{
	'X': {0, 0},
	'^': {1, 0},
	'A': {2, 0},
	'<': {0, 1},
	'v': {1, 1},
	'>': {2, 1},
}

func verticalFirst(dx, dy int) string {
	var s strings.Builder
	for dy < 0 {
		s.WriteRune('^')
		dy++
	}
	for dy > 0 {
		s.WriteRune('v')
		dy--
	}
	for dx < 0 {
		s.WriteRune('<')
		dx++
	}
	for dx > 0 {
		s.WriteRune('>')
		dx--
	}
	s.WriteRune('A')
	return s.String()
}

func horizontalFirst(dx, dy int) string {
	var s strings.Builder
	for dx < 0 {
		s.WriteRune('<')
		dx++
	}
	for dx > 0 {
		s.WriteRune('>')
		dx--
	}
	for dy < 0 {
		s.WriteRune('^')
		dy++
	}
	for dy > 0 {
		s.WriteRune('v')
		dy--
	}
	s.WriteRune('A')
	return s.String()
}

func getMoves(dx, dy int) (moves []string) {
	moves = []string{verticalFirst(dx, dy)}
	if dx != 0 && dy != 0 {
		moves = append(moves, horizontalFirst(dx, dy))
	}
	return moves
}

func movesOverEmpty(p Point, move string, keypad Keypad) bool {
	for _, c := range move {
		switch c {
		case 'v':
			p.Y++
		case '^':
			p.Y--
		case '<':
			p.X--
		case '>':
			p.X++
		}
		if keypad['X'] == p {
			return true
		}
	}
	return false
}

type MemoKey struct {
	input   string
	keypads int
}
type MemoValue struct {
	length int
	ok     bool
}

var mem = make(map[MemoKey]MemoValue)

func dfs(input string, keypads []Keypad) (int, bool) {
	key := MemoKey{input: input, keypads: len(keypads)}
	if cached, ok := mem[key]; ok {
		return cached.length, cached.ok
	}
	p := keypads[0]['A']
	var length int
	for _, c := range input {
		np := keypads[0][c]
		dx := np.X - p.X
		dy := np.Y - p.Y
		// we have 2 options:
		// 1. vertical -> horizontal
		// 2. horizontal -> vertical
		// also we need to avoid empty space
		moves := getMoves(dx, dy)
		var shortest int
		var found bool
		for _, move := range moves {
			if movesOverEmpty(p, move, keypads[0]) {
				continue
			}
			if len(keypads) == 1 {
				shortest = len(move)
				found = true
				break
			}
			if candidate, ok := dfs(move, keypads[1:]); ok {
				if !found || candidate < shortest {
					shortest = candidate
					found = true
				}
			}
		}
		if !found {
			mem[key] = MemoValue{length: 0, ok: false}
			return 0, false
		}
		length += shortest
		p = np
	}
	mem[key] = MemoValue{length: length, ok: true}
	return length, true
}

func getSum(lines []string, n int) int {
	clear(mem) // this is needed because number of keypads is changing, and so does their meaning
	keypads := []Keypad{numKeypad}
	for i := 0; i < n; i++ {
		keypads = append(keypads, dirKeypad)
	}
	var sum int
	for _, line := range lines {
		ops, ok := dfs(line, keypads)
		if !ok {
			fmt.Printf("no moves found for %s\n", line)
			return 0
		}
		num, err := strconv.Atoi(line[:len(line)-1])
		if err != nil {
			log.Fatal(err.Error())
		}
		sum += num * ops
	}
	return sum
}
