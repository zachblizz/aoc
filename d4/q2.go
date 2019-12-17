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
	fmt.Println(min, max)
	count := 0

	for i := min; i < max; i++ {
		if neverDecrease(i) && hasTwoAdjacent(i) {
			count++
		}
	}

	return count
}

func hasTwoAdjacent(num int) bool {
	str := strings.Split(strconv.Itoa(num), "")
	pm := make(map[string]int)
	i := 0

	for ; i < len(str)-1; i++ {
		pm[str[i]]++
	}

	pm[str[i]]++

	for _, v := range pm {
		if v == 2 {
			fmt.Println(pm)
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
