package main

import (
	utils "aoc-2020/utils"
	"fmt"
	"strconv"
	"strings"
)

type bag struct {
	color string

	amt int
	// key - bag color, value - how many
	contains map[string]bag
}

func main() {
	input := utils.ReadFile("sample.txt")

	partOne(parseInput(input))
}

func partOne(input *[]bag) {
	for _, bag := range *input {
		fmt.Println(bag)
	}
}

func parseInput(input *[]string) *[]bag {
	var ret []bag

	for _, row := range *input {
		parentBag := bag{
			contains: make(map[string]bag),
		}

		for i, desc := range strings.Split(row, ",") {
			bags := strings.Split(desc, " ")

			if i == 0 {
				parentBag.color = getColor(bags, 0, 1)
				color := getColor(bags, 5, 6)
				if color !=	"other bags." {
					amt, _ := strconv.Atoi(bags[4])
					parentBag.contains[color] = bag{color: color, amt: amt}
				}
			} else {
				color := getColor(bags, 2, 3)
				if color !=	"other bags." {
					amt, _ := strconv.Atoi(bags[1])
					parentBag.contains[color] = bag{color: color, amt: amt}
				}
			}
		}

		ret = append(ret, parentBag)
	}

	return &ret
}

func getColor(desc []string, start, end int) string {
	var color strings.Builder

	color.WriteString(desc[start])
	color.WriteString(" ")
	color.WriteString(desc[end])

	return color.String()
}
