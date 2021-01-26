package main

import (
	"fmt"
	"strings"

	"github.com/amar-preet/advent-of-code-2020/reader"
)

//wip
func main() {
	x := 5 + (8*3 + 9 + 3*4*3)
	fmt.Println(x)
	fileContent := reader.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(fileContent)), "\n")
	fmt.Println(lines)
}
