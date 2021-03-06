package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/amar-preet/advent-of-code-2020/reader"
)

func main() {
	lines := strings.Split(string(reader.ReadFile("input.txt")), "\n")
	validPassword(lines)
	validPasswordByPosition(lines)
}

func validPassword(lines []string) {
	var validPassword = 0
	for i := 0; i < len(lines); i++ {
		passwordPolicy := regexp.MustCompile("[\\-\\,\\:\\s]+").Split(lines[i], -1)
		if len(passwordPolicy) == 4 {
			atleast, _ := strconv.Atoi(passwordPolicy[0])
			atMost, _ := strconv.Atoi(passwordPolicy[1])
			character := passwordPolicy[2]
			sequence := passwordPolicy[3]

			var characterCount = 0
			for _, c := range sequence {
				if string(c) == character {
					characterCount++
				}
			}
			if atleast <= characterCount && characterCount <= atMost {
				validPassword++
			}
		}
	}
	fmt.Println("Total Valid Passwords:", validPassword)
}

func validPasswordByPosition(lines []string) {
	var validPassword = 0
	for i := 0; i < len(lines); i++ {
		passwordPolicy := regexp.MustCompile("[\\-\\,\\:\\s]+").Split(lines[i], -1)
		if len(passwordPolicy) == 4 {
			firstCharacter, _ := strconv.Atoi(passwordPolicy[0])
			secondCharacter, _ := strconv.Atoi(passwordPolicy[1])
			character := passwordPolicy[2]
			sequence := passwordPolicy[3]
			sequenceArray := strings.Split(sequence, "")

			if sequenceArray[firstCharacter-1] == character && sequenceArray[secondCharacter-1] == character {
				// invalid
			} else if sequenceArray[firstCharacter-1] == character && sequenceArray[secondCharacter-1] != character {
				validPassword++
			} else if sequenceArray[secondCharacter-1] == character && sequenceArray[firstCharacter-1] != character {
				validPassword++
			} else {
				// invalid
			}
		}
	}
	fmt.Println("Total Valid Passwords By Position:", validPassword)
}
