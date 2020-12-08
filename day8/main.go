package main

import (
	"fmt"
	"strings"

	"github.com/amar-preet/advent-of-code-2020/reader"
)

func main() {
	lines := strings.Split(string(reader.ReadFile("input.txt")), "\n")
	fmt.Println(halting(lines))
	part2(lines)
}

func halting(lines []string) (int, error) {
	var currentIndex = 0
	var globalCounter = 0
	encountered := map[int]bool{}

	for currentIndex < len(lines) {
		if encountered[currentIndex] == true {
			return globalCounter, fmt.Errorf("loop")
		}
		encountered[currentIndex] = true

		var instruction string
		var argument int
		fmt.Sscanf(lines[currentIndex], "%s %d", &instruction, &argument)

		if instruction == "acc" {
			globalCounter += argument
		}
		if instruction == "jmp" {
			currentIndex += argument - 1
		}

		currentIndex++
	}

	return globalCounter, nil
}

func part2(lines []string) {
	for i, line := range lines {
		tmpLines := make([]string, len(lines))
		copy(tmpLines, lines)
		tmpLines[i] = strings.NewReplacer("jmp", "nop", "nop", "jmp").Replace(line)

		if acc, err := halting(tmpLines); err == nil {
			fmt.Println(acc)
			break
		}
	}
}
