package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/amar-preet/advent-of-code-2020/reader"
)

func main() {
	lines := strings.Split(string(reader.ReadFile("input.txt")), "\n")
	partOne(lines)
	partTwo(lines)
}

func partOne(lines []string) {
	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			a, _ := strconv.Atoi(lines[i])
			b, _ := strconv.Atoi(lines[j])

			if a+b == 2020 {
				fmt.Println("i:", lines[i], "j", lines[j])
				fmt.Println("i*j:", a*b)
				break
			}
		}
	}
}

func partTwo(lines []string) {
	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			for k := j + 1; k < len(lines); k++ {
				a, _ := strconv.Atoi(lines[i])
				b, _ := strconv.Atoi(lines[j])
				c, _ := strconv.Atoi(lines[k])

				if a+b+c == 2020 {
					fmt.Println("i:", lines[i], "j", lines[j], "k", lines[k])
					fmt.Println("i*j*k:", a*b*c)
					break
				}
			}
		}
	}
}
