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
	// key - ans, val - map -> key - idx, val - doesn't matter
	m := make(map[rune]map[int]int)
	ret := 0

	for _, r := range *input {
		if r == "" {
			fmt.Print("ques:")
			for v := range m {
				fmt.Printf(" %v", string(v))
			}
			fmt.Println()

			m = make(map[rune]map[int]int)
		}

		for idx, ans := range r {
			idxM, ok := m[ans]
			_, idxOk := idxM[idx]

			if !ok || !idxOk {
				ret++

				if !ok {
					m[ans] = map[int]int{idx: 1}
				} else {
					m[ans][idx] = 1
				}
			}
		}
	}

	fmt.Print("ques:")
	for v := range m {
		fmt.Printf(" %v", string(v))
	}
	fmt.Println()

	fmt.Println(ret)
}
