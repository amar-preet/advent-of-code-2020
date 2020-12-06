package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/amar-preet/advent-of-code-2020/reader"
)

func main() {

	lines := strings.Split(string(reader.ReadFile("input.txt")), "\n")
	var seatIDs []int
	for i := 0; i < len(lines); i++ {
		row, column := binaryBoarding(lines[i])
		seatID := row*8 + column
		seatIDs = append(seatIDs, seatID)
	}
	fmt.Println("highest seat ID on a boarding pass", findMinAndMax(seatIDs))
	partTwo(seatIDs)
}

func findMinAndMax(seatIDs []int) int {
	max := seatIDs[0]
	min := seatIDs[0]
	for _, value := range seatIDs {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return max
}

func binaryBoarding(lines string) (int, int) {
	var row = 0
	var column = 0
	inputArray := strings.Split(lines, "")
	rowsArray := inputArray[0:7]
	columnArray := inputArray[7:10]
	row = parseRows(rowsArray)
	column = parseColumns(columnArray)
	return row, column
}

func parseColumns(rowsArray []string) int {
	var min = 0
	var max = 0
	for i := 0; i < len(rowsArray); i++ {
		if i == 0 {
			if rowsArray[i] == "L" {
				min = 0
				max = 3
			}
			if rowsArray[i] == "R" {
				min = 4
				max = 7
			}
		} else if i == 2 {
			if rowsArray[i] == "L" {
				return min
			}
			if rowsArray[i] == "R" {
				return max
			}
		} else {
			min, max = giveMinMax(rowsArray[i], min, max)
		}

	}
	return 0
}

func parseRows(rowsArray []string) int {
	var min = 0
	var max = 0
	for i := 0; i < len(rowsArray); i++ {
		if i == 0 {
			if rowsArray[i] == "F" {
				min = 0
				max = 63
			}
			if rowsArray[i] == "B" {
				min = 64
				max = 127
			}
		} else if i == 6 {
			if rowsArray[i] == "F" {
				return min
			}
			if rowsArray[i] == "B" {
				return max
			}
		} else {
			min, max = giveMinMax(rowsArray[i], min, max)
		}

	}
	return 0
}

func giveMinMax(code string, lower int, upper int) (int, int) {
	var min int
	var max int
	if code == "B" || code == "R" {
		min = ((upper-lower)+1)/2 + lower
		max = upper
	} else if code == "F" || code == "L" {
		min = lower
		max = (upper-lower)/2 + lower
	}
	return min, max
}

func partTwo(seatIDs []int) {
	sort.Ints(seatIDs)
	for i := range seatIDs {
		if seatIDs[i+1] != seatIDs[i]+1 {
			fmt.Println(seatIDs[i] + 1)
			break
		}
	}
}
