package main

import (
	utils "aoc-2020/utils"
	"fmt"
	"strings"
)

func main() {
	input := utils.ReadFile("input.txt")
	partOne(input)
}

/*
MINE
	r1-d1 - 78
	r3-d1 - 178
	r5-d1 - 75	
	r7-d1 - 86
	r1-d2 - 39

	3492520200
*/
/*
SAMPLE
	r1-d1 - 2
	r3-d1 - 7
	r5-d1 - 3
	r7-d1 - 4
	r1-d2 - 2
*/
func partOne(input *[]string) {
	good := twoDInput(input)
	treeCount := 0
	keepGoing := true
	tick := 0
	tickAmt := 1
	rowAmt := 2

	for row := 0; row < len(good) && keepGoing; row += rowAmt {
		if tick >= len(good[0]) {
			keepGoing = false
			break
		}

		spot := good[row][tick]

		if row > 0 && spot == "#" {
			treeCount++
		}

		tick += tickAmt
	}

	fmt.Println(treeCount)
}

func twoDInput(input *[]string) [][]string {
	var ret [][]string
	var tmp []string

	for _, row := range *input {
		var rowStr strings.Builder

		for i := 0; i < 1590; i++ {
			rowStr.WriteString(row)
		}

		tmp = append(tmp, rowStr.String())
	}

	for _, row := range tmp {
		var inner []string

		for _, str := range strings.Split(row, "") {
			inner = append(inner, str)
		}

		ret = append(ret, inner)
	}

	return ret
}
