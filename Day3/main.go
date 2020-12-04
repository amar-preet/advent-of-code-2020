package main

import (
	"fmt"
	"strings"

	"github.com/amar-preet/advent-of-code-2020/reader"
)

func main() {
	lines := strings.Split(string(reader.ReadFile("input.txt")), "\n")
	righ1Down1 := countTreesByRightDown(lines, 1, 1)
	righ3Down1 := countTreesByRightDown(lines, 3, 1)
	fmt.Println("total tree encountered right 3 down 1", righ3Down1)
	righ5Down1 := countTreesByRightDown(lines, 5, 1)
	righ7Down1 := countTreesByRightDown(lines, 7, 1)
	righ1Down2 := countTreesByRightDown(lines, 1, 2)
	fmt.Println("total tree encountered after muliply", righ1Down1*righ3Down1*righ5Down1*righ7Down1*righ1Down2)
}

/*
row: 0, column: 0 0%11
row: 1, column: 3 1*3%11
row: 2, column: 6 2*3%11
row: 3, column: 9 3*3%11
row: 4, column: 12 (going outside of 11)
row: 5, column: 15
and so on..
..##.......
#..0#...#..
.#....X..#.
..#.#...#0#
.X...##..#.
..#.X#.....
.#.#.#.0..#
.#........X
#.X#...#...
#...#X....#
.#..#...X#
*/

func countTreesByRightDown(lines []string, right int, down int) int {
	var treeCount = 0
	for row := 0; row < len(lines); row = row + down {
		var column = (row * right) % len(lines[row])
		slopeArray := strings.Split(lines[row], "")

		if len(slopeArray) >= column {
			if slopeArray[column] == "#" {
				treeCount++
			}
		}
	}
	return treeCount
}
