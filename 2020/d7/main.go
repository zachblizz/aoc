package main

import (
	utils "aoc-2020/utils"
	"fmt"
	"strconv"
	"strings"
)

type bag struct {
	color     string
	amt       int // amount of bags
	otherBags int
	// key - bag color, value - how many
	contains   map[string]bag
	canContain bool
}

func main() {
	input := utils.ReadFile("sample2.txt")
	// partOne(input)
	partTwo(input)
}

func partTwo(input *[]string) {
	seent := 0
	count := 1
	bags := parseInput(input)

	p2Helper(bags, &seent, &count)
}

func p2Helper(bags []*bag, seent, count *int) {
	if seent == len(bags) {
		return
	}

	*seent++

	for _, bag := range bags {

	}
}

func partOne(input *[]string) {
	seen := 0
	bags := parseInput(input)

	p1Helper(bags, "shiny gold", &seen)

	count := 0
	for _, bag := range bags {
		if bag.canContain == true {
			count++
		}
	}

	fmt.Println(count)
}

func p1Helper(bags []*bag, lookingFor string, seen *int) {
	if *seen == len(bags) {
		return
	}

	*seen++

	for _, bag := range bags {
		if _, ok := bag.contains[lookingFor]; ok {
			bag.canContain = true
			p1Helper(bags, bag.color, seen)
		}
	}
}

func parseInput(input *[]string) []*bag {
	var ret []*bag

	for _, row := range *input {
		parentBag := bag{
			contains: make(map[string]bag),
		}

		for i, desc := range strings.Split(row, ",") {
			bags := strings.Split(desc, " ")

			if i == 0 {
				parentBag.color = getColor(bags, 0, 1)
				color := getColor(bags, 5, 6)
				if color != "other bags." {
					amt, _ := strconv.Atoi(bags[4])
					parentBag.contains[color] = bag{color: color, amt: amt}
				}
			} else {
				color := getColor(bags, 2, 3)
				if color != "other bags." {
					amt, _ := strconv.Atoi(bags[1])
					parentBag.contains[color] = bag{color: color, amt: amt}
				}
			}
		}

		ret = append(ret, &parentBag)
	}

	return ret
}

func getColor(desc []string, start, end int) string {
	var color strings.Builder

	color.WriteString(desc[start])
	color.WriteString(" ")
	color.WriteString(desc[end])

	return color.String()
}
