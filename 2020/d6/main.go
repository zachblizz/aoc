package main

import (
	utils "aoc-2020/utils"
	"fmt"
)

func main() {
	input := utils.ReadFile("input.txt")
	partOne(input)
}

func partOne(input *[]string) {
	m := make(map[rune]int)
	final := 0
	sum := 0

	for _, r := range *input {
		if r == "" {
			m = make(map[rune]int)
			final += sum
			sum = 0
		}

		for _, ans := range r {
			if _, ok := m[ans]; !ok && isAnswerValid(ans) {
				sum++
				m[ans] = 1
			}
		}
	}

	fmt.Println(final + sum)
}

func isAnswerValid(ans rune) bool {
	aToC := ans >= 'a' && ans <= 'c'
	xToZ := ans >= 'x' && ans <= 'z'

	return aToC || xToZ
}
