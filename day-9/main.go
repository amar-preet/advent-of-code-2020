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
	invalidNumber := invalidXmas(intLines, 25)
	fmt.Println(invalidNumber)
	encryptionWeakness(intLines, invalidNumber)
}

func invalidXmas(intLines []int, preamble int) int {
	var numberToSum = 0
	for i := 0; i < len(intLines); i++ {
		numberToSum = intLines[preamble]
		nextFive := intLines[i:preamble]
		if sumofPreamble(nextFive, numberToSum) {
			preamble++
		} else {
			break
		}
	}
	return numberToSum
}

func sumofPreamble(nextFive []int, numberToSum int) bool {
	var found = false
loop:
	for j := 0; j < len(nextFive); j++ {
		for k := 1; k < len(nextFive); k++ {
			if numberToSum == nextFive[j]+nextFive[k] {
				found = true
				break loop
			}
		}

	}
	return found
}

func encryptionWeakness(intLines []int, invalidNumber int) {
	for i := 0; i < len(intLines); i++ {
		for j := i + 1; j < len(intLines); j++ {
			sum := 0
			for _, v := range intLines[i : j+1] {
				sum += v
			}
			if sum == invalidNumber {
				sort.Ints(intLines[i : j+1])
				fmt.Println(intLines[i] + intLines[j])
				return
			}
		}
	}
}
