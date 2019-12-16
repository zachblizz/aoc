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

// TODO: account for 244444... and the like
// func getNumOfDiffs(max, count int, curr []string) int {
// 	currNum, _ := strconv.Atoi(strings.Join(curr, ""))

// 	if currNum >= max {
// 		return count
// 	}

// 	prev, _ := strconv.Atoi(curr[0])
// 	for i, n := range curr[1:len(curr)] {
// 		c, _ := strconv.Atoi(string(n))

// 		if c < prev {
// 			// if this happens mark the rest as prev...
// 			curr[i+1] = strconv.Itoa(prev)
// 		} else {
// 			prev = c
// 		}
// 	}

// 	count++
// 	newNum, _ := strconv.Atoi(strings.Join(curr, ""))
// 	fmt.Println(newNum)
// 	newNum++

// 	return getNumOfDiffs(max, count, strings.Split(strconv.Itoa(newNum), ""))
// }

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
