package main

import (
	"fmt"
	"strings"

	"github.com/amar-preet/advent-of-code-2020/reader"
)

func main() {
	lines := strings.Split(string(reader.ReadFile("input.txt")), "\n")
	part1(lines)
	part2(lines)
}

func part1(input []string) {
	eastWest := 0
	northSouth := 0

	dirx := []int{1, 0, -1, 0}
	diry := []int{0, -1, 0, 1}

	currentDir := 0

	for _, l := range input {
		var direction rune
		var value int
		fmt.Sscanf(l, "%c%d", &direction, &value)
		switch direction {
		case 'N':
			northSouth += value
		case 'S':
			northSouth -= value
		case 'E':
			eastWest += value
		case 'W':
			eastWest -= value
		case 'L':
			turn := value / 90
			currentDir -= turn
			currentDir %= 4

			if currentDir < 0 {
				currentDir += 4
			}

		case 'R':
			turn := value / 90
			currentDir += turn
			currentDir %= 4
			if currentDir < 0 {
				currentDir += 4
			}

		case 'F':
			eastWest += (dirx[currentDir] * value)
			northSouth += (diry[currentDir] * value)

		}
	}

	if northSouth < 0 {
		northSouth = -northSouth
	}
	if eastWest < 0 {
		eastWest = -eastWest
	}
	fmt.Println("Manhattan distance:", eastWest+northSouth)
}

func part2(input []string) {
	eastWest := 0
	northSouth := 0

	wx := 10
	wy := 1

	for _, l := range input {
		var direction rune
		var value int
		fmt.Sscanf(l, "%c%d", &direction, &value)
		switch direction {
		case 'N':
			wy += value
		case 'S':
			wy -= value
		case 'E':
			wx += value
		case 'W':
			wx -= value
		case 'L':
			turn := value / 90

			for i := 0; i < turn; i++ {
				wx, wy = -wy, wx
			}

		case 'R':
			turn := value / 90

			for i := 0; i < turn; i++ {
				wx, wy = wy, -wx
			}

		case 'F':
			eastWest += value * wx
			northSouth += value * wy
		}
	}

	if northSouth < 0 {
		northSouth = -northSouth
	}
	if eastWest < 0 {
		eastWest = -eastWest
	}
	fmt.Println(eastWest + northSouth)
}
