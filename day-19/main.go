package main

import (
	"fmt"
	"strings"

	"github.com/amar-preet/advent-of-code-2020/reader"
)

//wip
func main() {
	lines := strings.Split(string(reader.ReadFile("input.txt")), "\n\n")
	rules := lines[1]
	message := lines[0]
	fmt.Println(message, rules)
}
