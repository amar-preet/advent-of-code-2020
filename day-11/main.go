package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/amar-preet/advent-of-code-2020/reader"
)

type seat struct {
	row    int
	column int
}

var seatMap = []seat{
	{row: 0, column: 1},
	{row: 0, column: -1},
	{row: 1, column: 0},
	{row: 1, column: 1},
	{row: 1, column: -1},
	{row: -1, column: 0},
	{row: -1, column: 1},
	{row: -1, column: -1},
}

func main() {
	lines := strings.Split(string(reader.ReadFile("input.txt")), "\n")
	seatsOccupied(lines, 4)
	seatsOccupied(lines, 5)
}

func seatsOccupied(input []string, seatSpace int) {
	curr := addAboveAndBelowRow(input)
	iteration, firstRound := make([]string, len(curr)), true
	copy(iteration, curr)

	for !equals(curr, iteration) || firstRound {
		firstRound = false
		copy(curr, iteration)

		for row := 1; row < len(curr)-1; row++ {
			if seatSpace == 4 {
				curr = fourOrMore(curr, row, iteration)
			} else {
				curr = fiveOrMore(curr, row, iteration)
			}
		}
	}

	fmt.Println(countCharInSlice(curr, "#"))
}

func fourOrMore(curr []string, row int, iteration []string) []string {
	for col := 1; col < len(curr[row])-1; col++ {
		next := make([]string, 8)

		for i, element := range seatMap {
			next[i] = string(curr[row+element.row][col+element.column])
		}

		count := countCharInSlice(next, "#")

		if string(curr[row][col]) == "L" && count == 0 {
			iteration[row] = replaceCharAt(iteration[row], "#", col)
		} else if string(curr[row][col]) == "#" && count >= 4 {
			iteration[row] = replaceCharAt(iteration[row], "L", col)
		}
	}
	return curr
}

func fiveOrMore(curr []string, row int, iteration []string) []string {
	for col := 1; col < len(curr[row])-1; col++ {
		next := make([]string, 8)

		for i, l := range seatMap {
			nextRow, nextCol := 0, 0
			currentRow := row + l.row + nextRow
			currentCol := col + l.column + nextCol

			c := "."
			for c == "." && currentRow >= 0 && currentRow < len(curr)-1 && currentCol >= 0 && currentCol < len(curr[i])-1 {
				c = string(curr[currentRow][currentCol])

				nextRow += l.row
				nextCol += l.column
				currentRow = row + l.row + nextRow
				currentCol = col + l.column + nextCol
			}

			next[i] = c
		}

		count := countCharInSlice(next, "#")

		if string(curr[row][col]) == "L" && count == 0 {
			iteration[row] = replaceCharAt(iteration[row], "#", col)
		} else if string(curr[row][col]) == "#" && count >= 5 {
			iteration[row] = replaceCharAt(iteration[row], "L", col)
		}
	}
	return curr
}

func addAboveAndBelowRow(lines []string) []string {
	linedLines := make([]string, len(lines)+1)
	dots := strings.Repeat(".", len(lines[0])+2)

	linedLines = append(linedLines, dots)
	linedLines[0] = dots

	for i := 1; i < len(lines)+1; i++ {
		linedLines[i] = fmt.Sprintf(".%s.", lines[i-1])
	}

	return linedLines
}

func equals(first []string, second []string) bool {
	return reflect.DeepEqual(first, second)
}

func countCharInSlice(slice []string, char string) int {
	numChars := 0
	for _, element := range slice {
		numChars += strings.Count(element, char)
	}

	return numChars
}

func replaceCharAt(str string, replacement string, index int) string {
	ret := ""

	for i := 0; i < len(str); i++ {
		if i != index {
			ret += string(str[i])
		} else {
			ret += replacement
		}
	}

	return ret
}
