package main

import (
	utils "aoc-2020/utils"
	"fmt"
	"strconv"
	"strings"
)

type lineInfo struct {
	high int
	low int
	char string
	pwd string
}

func main() {
	input := utils.ReadFile("input.txt")
	partOne(input)
	partTwo(input)
}

func partOne(input *[]string) {
	correct := 0

	for _, line := range *input {
		info := getLineInfo(line)
		numChar := 0

		for _, r := range info.pwd {
			if string(r) == info.char {
				numChar++
			}
		}

		if numChar <= info.high && numChar >= info.low {
			correct++
		}
	}

	fmt.Println(correct)
}

func partTwo(input *[]string) {
	correct := 0

	for _, line := range *input {
		info := getLineInfo(line)

		start := 0
		end := info.high

		if info.low > 0 {
			start = info.low - 1
		}

		if info.high > 0 {
			end = info.high - 1
		}

		if isPasswordCorrect(info.pwd, info.char, start, end) {
			correct++
		}
	}

	fmt.Println(correct)
}

func isPasswordCorrect(pwd, char string, start, end int) bool {
	correctWayOne := string(pwd[start]) == char && string(pwd[end]) != char
	correctWayTwo := string(pwd[start]) != char && string(pwd[end]) == char

	return correctWayOne || correctWayTwo
}

func getLineInfo(line string) lineInfo {
	var info lineInfo

	s := strings.Split(line, ":")
	highLowChar := strings.Split(strings.Replace(s[0], " ", "-", 1), "-")
	info.pwd = strings.TrimLeft(s[1], " ")
	info.low, _ = strconv.Atoi(string(highLowChar[0]))
	info.high, _ = strconv.Atoi(string(highLowChar[1]))
	info.char = string(highLowChar[2])

	return info
}
