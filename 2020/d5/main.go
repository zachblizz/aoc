package main

import (
	utils "aoc-2020/utils"
	"fmt"
)

const (
	totalRows  = 127
	totalSeats = 7
)

func main() {
	input := utils.ReadFile("input.txt")
	partOne(input)
	partTwo(input)
}

func partOne(input *[]string) {
	lowestID := 10000
	highestID := 0

	for _, i := range *input {
		getPassID(&i, &highestID, &lowestID, make(map[int]int))
	}

	fmt.Println("lowest id: ", lowestID)
	fmt.Println("highest id: ", highestID)
}

func partTwo(input *[]string) {
	lowestID := 10000
	highestID := 0
	m := make(map[int]int)

	for _, i := range *input {
		getPassID(&i, &highestID, &lowestID, m)
	}

	run := lowestID + 1
	for ; lowestID < highestID; lowestID++ {
		if _, ok := m[run]; !ok {
			fmt.Println("my passid is: ", run)
			break
		}

		run++
	}
}

func getPassID(position *string, high, low *int, m map[int]int) {
	sr := 0
	rtr := totalRows
	er := totalRows

	ss := 0
	rts := totalSeats
	es := totalSeats

	for _, c := range *position {
		pos := string(c)

		switch pos {
		case "F":
			er = rtr / 2
		case "B":
			sr = (rtr / 2) + 1
		case "L":
			es = rts / 2
		case "R":
			ss = (rts / 2) + 1
		}

		rtr = er + sr
		rts = es + ss
	}

	id := sr*8 + ss
	// fmt.Printf("seat id: %d\n", id)
	m[id] = 1

	if id < *low {
		*low = id
	}

	if id > *high {
		*high = id
	}
}
