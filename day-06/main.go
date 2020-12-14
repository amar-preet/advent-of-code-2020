package main

import (
	"fmt"
	"strings"

	"github.com/amar-preet/advent-of-code-2020/reader"
)

func main() {
	fileContent := reader.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(fileContent)), "\n\n")
	fmt.Println("Sum: ", sumOfCount(lines))
}

func sumOfCount(lines []string) int {
	var part1 = 0
	var part2 = 0
	for i := 0; i < len(lines); i++ {
		inputArray := strings.Split(strings.TrimSpace(lines[i]), "")
		part1 = part1 + len(numberOfQuestionsAnyoneAnswered(inputArray))
		if numberOfPeopleAnswered(inputArray) == 1 {
			part2 = part2 + len(inputArray)
		} else {
			part2 = part2 + len(numberOfQuestionsEveryoneAnswered(inputArray))
		}
	}
	fmt.Println("part2:", part2)
	return part1
}

func numberOfPeopleAnswered(inputArray []string) int {
	var count = 1
	for i := 0; i < len(inputArray); i++ {
		if strings.TrimSpace(inputArray[i]) == "" {
			count++
		}
	}
	return count
}

func numberOfQuestionsAnyoneAnswered(inputArray []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for v := range inputArray {
		if encountered[inputArray[v]] == true {
		} else {
			if strings.TrimSpace(inputArray[v]) != "" {
				encountered[inputArray[v]] = true
				result = append(result, inputArray[v])
			}
		}
	}
	return result
}

func numberOfQuestionsEveryoneAnswered(inputArray []string) []string {
	encountered := map[string]bool{}
	result := []string{}
	for v := range inputArray {
		if encountered[inputArray[v]] == true {
			result = append(result, inputArray[v])
		} else {
			if strings.TrimSpace(inputArray[v]) != "" {
				encountered[inputArray[v]] = true

			}
		}
	}
	return result
}
