package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/amar-preet/advent-of-code-2020/reader"
)

func main() {
	lines := strings.Split(string(reader.ReadFile("input.txt")), ",")
	fmt.Println(findNthnumberFast(lines, 2020))
}

// super slow approach O(N^2)
func findNthNumber(lines []string, till int) map[int]int {

	var numberMap = map[int]int{}
	for start := range lines {
		x, _ := strconv.Atoi(lines[start])
		numberMap[start+1] = x
	}

	for i := len(numberMap); i < till; i++ {
		var islastNumberNew = true
		lastNumber := numberMap[i]
		for j := len(numberMap) - 1; j > 0; j-- {
			if numberMap[j] == lastNumber {
				islastNumberNew = false
				break
			}
		}
		if islastNumberNew {
			numberMap[i+1] = 0
		} else {
			var secondOccurence = 0
			for k := len(numberMap) - 1; k > 0; k-- {
				if numberMap[k] == lastNumber {
					secondOccurence = k
					break
				}
			}
			numberMap[i+1] = i - secondOccurence
		}
	}
	return numberMap
}

// bit faster O(N)
func findNthnumberFast(lines []string, till int) int {
	var numberMap = map[int]int{}
	var last = 0
	for i, s := range lines {
		last, _ = strconv.Atoi(s)
		numberMap[last] = i + 1
	}

	for i := len(numberMap); i < till; i++ {
		if v, ok := numberMap[last]; ok {
			numberMap[last], last = i, i-v
		} else {
			numberMap[last], last = i, 0
		}

		if i == till {
			break
		}
	}
	return last
}
