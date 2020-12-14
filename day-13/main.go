package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/amar-preet/advent-of-code-2020/reader"
)

func main() {
	lines := strings.Split(string(reader.ReadFile("input.txt")), "\n")
	earliestDepartingtime, _ := strconv.Atoi(lines[0])
	busIds := lines[1]
	earliestBusTime(earliestDepartingtime, busIds)
}

func earliestBusTime(earliestDepartingtime int, busIds string) {
	busArray := strings.Split(busIds, ",")
	max, id := math.MaxInt64, 0

	for _, element := range busArray {
		if element != "x" {
			count := 0
			bus, _ := strconv.Atoi(element)
			for count <= earliestDepartingtime {
				count += bus
			}

			if count < max {
				id = bus
				max = count
			}
		}
	}

	fmt.Println((max - earliestDepartingtime) * id)
}
