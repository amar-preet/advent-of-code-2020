package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/amar-preet/advent-of-code-2020/reader"
)

func main() {
	fileContent := reader.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(fileContent)), "\n\n")
	notes := getSlice(lines[0])
	//yourTicket := getSlice(lines[1])
	nearbyTicket := getSlice(lines[2])
	validNumbers := getValidNumber(notes)
	errorRate(validNumbers, nearbyTicket)
}

func getSlice(input string) []string {
	return strings.Split(input, "\n")
}

func getValidNumber(notes []string) []int {
	var invalidNumbers []int
	var validNumber []int
	for _, note := range notes {
		row := regexp.MustCompile(strings.TrimSpace("[\\:\\-\\s]+")).Split(note, -1)
		for _, rule := range row {
			if val, err := strconv.Atoi(rule); err == nil {
				validNumber = append(validNumber, val)
			}
		}
	}
	//[1,3,5,7,6,11,33,44,13,40,45,50]
	for i := 0; i < len(validNumber); i += 4 {
		invalid := validNumber[i+1]
		for invalid+1 < validNumber[i+2] {
			invalid = invalid + 1
			invalidNumbers = append(invalidNumbers, invalid)
		}
	}
	return validNumber
}

//[1,3,5,7,6,11,33,44,13,40,45,50]
func errorRate(validNumbers []int, nearbyTicket []string) {
	var count = 0
out:
	for _, ticket := range nearbyTicket {
		row := regexp.MustCompile(strings.TrimSpace(",")).Split(ticket, -1)
	in:
		for _, rule := range row {
			if val, err := strconv.Atoi(rule); err == nil {
				var found = false
				for i := 0; i < len(validNumbers); i += 4 {
					if (validNumbers[i] <= val && val <= validNumbers[i+1]) ||
						(validNumbers[i+2] <= val && val <= validNumbers[i+3]) {
						found = true
						continue in
					}
				}
				if !found {
					count += val
					continue out
				}

			}
		}
	}
	fmt.Println("ticket scanning error rate: ", count)
}

// part 2 wip
func sumOfTickets() {

}
