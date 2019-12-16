package main

import (
	"fmt"
	"strconv"
	"strings"

)

func parseInput(input *string) ([]string, int) {
	s := strings.Split(*input, "-")

	max, _ := strconv.Atoi(s[1])
	return strings.Split(s[0], ""), max
}

// TODO: account for 244444... and the like
func getNumOfDiffs(max, count int, curr []string) int {
	currNum, _ := strconv.Atoi(strings.Join(curr, ""))

	if currNum >= max {
		return count
	}

	prev, _ := strconv.Atoi(curr[0])
	for i, n := range curr[1:len(curr)] {
		c, _ := strconv.Atoi(string(n))

		if c < prev {
			curr[i+1] = strconv.Itoa(prev)
		} else {
			getPrev = false
			prev = c
		}
	}

	count++
	newNum, _ := strconv.Atoi(strings.Join(curr, ""))
	fmt.Println(newNum)
	newNum++

	return getNumOfDiffs(max, count, strings.Split(strconv.Itoa(newNum), ""))
}

func main() {
	input := "240920-789857"
	min, max := parseInput(&input)

	count := getNumOfDiffs(max, 0, min)
	fmt.Println(count)
}
