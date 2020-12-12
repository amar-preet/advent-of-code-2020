package main

import (
	"fmt"
	"math"
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

	fmt.Println("Manhattan distance:", math.Abs(float64(eastWest))+math.Abs(float64(northSouth)))
}

func part2(input []string) {
	eastWest := 0
	northSouth := 0

	x := 10
	y := 1

	for _, l := range input {
		var direction rune
		var value int
		fmt.Sscanf(l, "%c%d", &direction, &value)
		switch direction {
		case 'N':
			y += value
		case 'S':
			y -= value
		case 'E':
			x += value
		case 'W':
			x -= value
		case 'L':
			turn := value / 90

			for i := 0; i < turn; i++ {
				x, y = -y, x
			}

		case 'R':
			turn := value / 90

			for i := 0; i < turn; i++ {
				x, y = y, -x
			}

		case 'F':
			eastWest += value * x
			northSouth += value * y
		}
	}

	fmt.Println("Manhattan distance (part2): ", math.Abs(float64(eastWest))+math.Abs(float64(northSouth)))
}
