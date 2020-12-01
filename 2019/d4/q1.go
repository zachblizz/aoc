package main

import (
	"fmt"
	"strconv"
	"strings"

)

func parseInput(input *string) (int, int) {
	s := strings.Split(*input, "-")

	min, _ := strconv.Atoi(s[0])
	max, _ := strconv.Atoi(s[1])
	return min, max
}

func getNumOfDiffs(max, min int) int {
	count := 0

	for i := min; i < max; i++ {
		if hasTwoAdjacent(i) && neverDecrease(i) {
			count++
		}
	}

	return count
}

func hasTwoAdjacent(num int) bool {
	str := strings.Split(strconv.Itoa(num), "")

	for i := 0; i < len(str)-1; i++ {
		if str[i] == str[i+1] {
			return true
		}
	}

	return false
}

func neverDecrease(num int) bool {
	str := strings.Split(strconv.Itoa(num), "")
	prev := str[0]

	for i := 1; i < len(str); i++ {
		if str[i] < prev {
			return false
		}

		prev = str[i]
	}

	return true
}

func main() {
	input := "240920-789857"
	min, max := parseInput(&input)

	count := getNumOfDiffs(max, min)
	fmt.Println(count)
}
