package main

import (
	utils "aoc-2020/utils"
	"fmt"
)

func main() {
	input := utils.ReadFile("sample.txt")
	// partOne(input)
	partTwo(input)
}

func partOne(input *[]string) {
	m := make(map[rune]int)
	ret := 0

	for _, r := range *input {
		if r == "" {
			m = make(map[rune]int)
		}

		for _, ans := range r {
			if _, ok := m[ans]; !ok {
				ret++
				m[ans] = 1
			}
		}
	}

	fmt.Println(ret)
}

func partTwo(input *[]string) {
	m := make(map[rune]int)
	ret := 0
	groupCount := 0

	for _, r := range *input {
		if r == "" {
			for _, v := range m {
				if v == groupCount {
					ret++
				}
			}

			m = make(map[rune]int)
			groupCount = 0
		}

		groupCount++

		for _, ans := range r {
			if _, ok := m[ans]; !ok {
				m[ans] = 1
			}
		}
	}

	fmt.Println(ret)
}
