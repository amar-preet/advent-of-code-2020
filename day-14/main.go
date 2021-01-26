package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/amar-preet/advent-of-code-2020/reader"
)

// wip
func main() {
	lines := strings.Split(string(reader.ReadFile("input.txt")), "\n")
	var mask = ""
	for _, v := range lines {
		line := strings.Split(v, " = ")
		if strings.Contains(line[1], "X") {
			mask = line[1]
			fmt.Println(mask)
		} else {
			memory, _ := strconv.Atoi(line[1])
			allBinaryNumbers(memory)
		}
	}
}

func allBinaryNumbers(length int) []string {
	binaries := make([]string, 0)

	to := ""
	for i := 0; i < length; i++ {
		to += "1"
	}

	var binary string
	for i := 0; binary != to; i++ {
		binary = strconv.FormatInt(int64(i), 2)
		binaries = append(binaries, binary)
	}

	return binaries
}
