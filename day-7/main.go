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
	findShinyBags(lines, "shiny gold")
}

func findShinyBags(lines []string, bagItem string) {
	bagMap := map[string]map[string]int{}
	for _, s := range lines {
		parentBag := regexp.MustCompile(`(.+) bags contain`).FindStringSubmatch(s)[1]

		bagMap[parentBag] = map[string]int{}
		childBags := regexp.MustCompile(`(\d+) (.+?) bag`).FindAllStringSubmatch(s, -1)
		for _, in := range childBags {
			bagMap[parentBag][in[2]], _ = strconv.Atoi(in[1])
		}
	}

	totalShinyGold := 0
	for bag := range bagMap {
		if isInBag(bagMap, bag, bagItem) {
			totalShinyGold++
		}
	}
	fmt.Println("totalShinyGold", totalShinyGold)
	fmt.Println(countBags(bagMap, bagItem))
}

func isInBag(bags map[string]map[string]int, bag string, bagItem string) bool {
	if _, ok := bags[bag][bagItem]; ok {
		return true
	}
	for out := range bags[bag] {
		if isInBag(bags, out, bagItem) {
			return true
		}
	}
	return false
}

func countBags(bags map[string]map[string]int, bag string) (total int) {
	for k, v := range bags[bag] {
		total += v * (countBags(bags, k) + 1)
	}
	return
}
