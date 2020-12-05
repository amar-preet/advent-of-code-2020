package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/amar-preet/advent-of-code-2020/reader"
)

func main() {
	fileContent := reader.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(fileContent)), "\n\n")
	validPassport(lines)
	validatePassportFields(lines)
}

func getPassportFields() []string {
	return []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
}

func getPassportFieldsByValidation() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(`byr:(19[2-9]\d|200[0-2])(\s+|$)`),
		regexp.MustCompile(`iyr:(201\d|2020)(\s+|$)`),
		regexp.MustCompile(`eyr:(202\d|2030)(\s+|$)`),
		regexp.MustCompile(`hgt:((1[5-8]\d|19[0-3])cm|(59|6\d|7[0-6])in)(\s+|$)`),
		regexp.MustCompile(`hcl:#[\da-f]{6}(\s+|$)`),
		regexp.MustCompile(`ecl:(amb|blu|brn|gry|grn|hzl|oth)(\s+|$)`),
		regexp.MustCompile(`pid:\d{9}(\s+|$)`),
	}
}

func validPassport(lines []string) {
	validCount := 0
line:
	for _, line := range lines {
		for _, passportItem := range getPassportFields() {
			if ok := strings.Contains(line, passportItem); !ok {
				continue line
			}
		}
		validCount++

	}
	fmt.Println(validCount)
}

func validatePassportFields(lines []string) {
	validCount := 0
line:
	for _, line := range lines {
		for _, passportItem := range getPassportFieldsByValidation() {
			if !passportItem.MatchString(line) {
				continue line
			}
		}
		validCount++

	}
	fmt.Println(validCount)
}
