package main

import (
	utils "aoc-2020/utils"
	"fmt"
	"strings"
)

func main() {
	input := utils.ReadFile("input.txt")
	all := [][]int{
		[]int{1, 1},
		[]int{3, 1},
		[]int{5, 1},
		[]int{7, 1},
		[]int{1, 2},
	}

	num := 1
	for _, amt := range all {
		tickAmt := amt[0]
		rowAmt := amt[1]

		num = num * partOne(input, tickAmt, rowAmt)
	}

	fmt.Println(num)
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
func partOne(input *[]string, tickAmt, rowAmt int) int {
	good := twoDInput(input)
	treeCount := 0
	tick := 0

	for row := 0; row < len(good) && tick < len(good[0]); row += rowAmt {
		spot := good[row][tick]

		if row > 0 && spot == "#" {
			treeCount++
		}

		tick += tickAmt
	}

	fmt.Println(treeCount)

	return treeCount
}

func twoDInput(input *[]string) [][]string {
	var ret [][]string
	var tmp []string

	for _, row := range *input {
		var rowStr strings.Builder

		for i := 0; i < 150; i++ {
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
