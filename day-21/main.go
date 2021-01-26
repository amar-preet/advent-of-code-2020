package main

import (
	"fmt"
	"strings"

	"github.com/amar-preet/advent-of-code-2020/reader"
)

//wip
func main() {
	lines := strings.Split(strings.TrimSpace(string(reader.ReadFile("input.txt"))), "\n")
	count, allergenMap := countIngredients(lines)
	fmt.Println(count)
	canonicalList(allergenMap)
}

func countIngredients(lines []string) (int, map[string]map[string]string) {
	count := map[string]int{}
	allergenMap := map[string]map[string]string{}
	var total = 0

	for _, s := range lines {
		ingredient := strings.Split(strings.TrimSpace(s), " (contains")
		food := ingredient[0]
		allergen := strings.TrimRight(strings.TrimSpace(ingredient[1]), ")")

		ingredientMap := map[string]string{}
		for _, s := range strings.Fields(food) {
			count[s]++
			ingredientMap[s] = s
		}

		for _, aller := range strings.Split(allergen, ", ") {
			if _, ok := allergenMap[aller]; !ok {
				allergenMap[aller] = map[string]string{}
				for i := range ingredientMap {
					allergenMap[aller][i] = aller
				}
				continue
			}

			for i := range allergenMap[aller] {
				if _, ok := ingredientMap[i]; !ok {
					delete(allergenMap[aller], i)
				}
			}
		}
	}

out:
	for i := range count {
		for _, is := range allergenMap {
			if _, ok := is[i]; ok {
				continue out
			}
		}
		total += count[i]
	}
	return total, allergenMap
}

func canonicalList(allergenMap map[string]map[string]string) {
	//var canonical = ""
	var cannonicalMap = map[string]string{}
	for i, v := range allergenMap {
		for _, v2 := range v {
			cannonicalMap[i] = v2
		}
	}
}
