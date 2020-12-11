package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/amar-preet/advent-of-code-2020/reader"
)

func main() {
	lines := strings.Split(string(reader.ReadFile("input.txt")), "\n")
	intLines := make([]int, len(lines))
	for i := range intLines {
		intLines[i], _ = strconv.Atoi(lines[i])
	}
	sort.Ints(intLines)
	findOneJoltThreeJolt(intLines)

	lastAdapter := append(intLines, intLines[len(intLines)-1]+3)
	// add 0 and last adapter
	//(0), 1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19, (22)
	intLines = append([]int{0}, lastAdapter...)
	fmt.Println("totalCombinations:", totalCombinations(0, intLines, make(map[int]int)))
}

func findOneJoltThreeJolt(intLines []int) {
	fmt.Println(oneJolt(intLines) * threeJolt(intLines))
}

func oneJolt(intLines []int) int {
	var oneVolt = 0
loop:
	for i := 0; i < len(intLines); i++ {
		for j := i + 1; j < len(intLines); j++ {
			if intLines[i]+1 == intLines[j] {
				oneVolt++
			}
			continue loop
		}
	}
	return oneVolt + 1
}

func threeJolt(intLines []int) int {
	var threeJolt = 0
loop:
	for i := 0; i < len(intLines); i++ {
		for j := i + 1; j < len(intLines); j++ {
			if intLines[i]+3 == intLines[j] {
				threeJolt++
			}
			continue loop
		}
	}
	return threeJolt + 1
}

func totalCombinations(currIndex int, intLines []int, adapterMap map[int]int) int {
	var count int
	currNum := intLines[currIndex]

	if currIndex >= len(intLines)-3 {
		return 1
	}

	if val, ok := adapterMap[currNum]; ok {
		return val
	}

	for i := currIndex + 1; i < currIndex+4; i++ {
		num := intLines[i]
		if currNum+1 == num || currNum+2 == num || currNum+3 == num {
			count += totalCombinations(i, intLines, adapterMap)
		}
	}

	adapterMap[currNum] = count
	return count
}
