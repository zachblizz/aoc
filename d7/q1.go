package main

import (
	"fmt"
	"strconv"

)

type programState struct {
	code string // op code

	sysID int // system check id

	ip int // instruction pointer
	p1 int
	p2 int
	p3 int

	modeOne   int
	modeTwo   int
	modeThree int

	jump int
}

func one(state *programState, input []int) {
	input[state.p3] = state.p1 + state.p2
	state.ip += state.jump
}

func two(state *programState, input []int) {
	input[state.p3] = state.p1 * state.p2
	state.ip += state.jump
}

func three(state *programState, input []int) {
	input[state.p1] = state.sysID
	state.ip += state.jump
}

func four(state *programState, input []int) {
	if state.p1 < len(input) {
		fmt.Printf("val @%v is %v\n", state.p1, input[state.p1])
	} else {
		fmt.Printf("@ pos %v\n", state.p1)
	}
	state.ip += state.jump
}

// modes
// 0 - position
// 1 - imediate
func runInstructions(input []int, opCodes map[string]interface{}, state programState) []int {
	for state.ip < len(input) && input[state.ip] != 99 {
		op := input[state.ip]
		getCurrState(op, &state, input)

		if _, ok := opCodes[state.code]; ok {
			opCodes[state.code].(func(*programState, []int))(&state, input)
		} else {
			state.ip++
		}
	}

	return input
}

func clearState(s *programState) {
	s.p1 = 0
	s.p2 = 0
	s.p3 = 0
	s.modeOne = 0
	s.modeTwo = 0
	s.modeThree = 0
	s.jump = 3
}

func getCurrState(op int, state *programState, input []int) {
	strOp := strconv.Itoa(op)
	clearState(state)

	if len(strOp) < 4 {
		strOp = fmt.Sprintf("%04v", op)
	}

	c := []rune(strOp)
	state.code = string(c[2:4])

	state.modeOne, _ = strconv.Atoi(string(c[1:2]))
	state.modeTwo, _ = strconv.Atoi(string(c[0:1]))

	if state.code == "01" || state.code == "02" || state.code == "07" || state.code == "08" {
		state.jump = 4
	} else if state.code == "03" || state.code == "04" {
		state.jump = 2
	}

	getParams(state, input)
}

func getParams(state *programState, input []int) {
	if state.code != "09" {
		state.p1 = input[state.ip+1] // imidiate mode assumed

		if state.code != "03" {
			state.p2 = input[state.ip+2]
		}

		if state.code == "01" || state.code == "02" || state.code == "07" || state.code == "08" {
			state.p3 = input[state.ip+3]
		}
	}

	// position mode swap
	if state.modeOne == 0 && state.code != "03" && state.code != "04" && state.p1 < len(input) {
		state.p1 = input[state.p1]
	}

	if state.modeTwo == 0 && state.p2 < len(input) {
		state.p2 = input[state.p2]
	}
}

func doInstructions(input []int, inputVal int) {
	opCodes := map[string]interface{}{
		"01": one,
		"02": two,
		"03": three,
		"04": four,
		"05": func(state *programState, _ []int) {
			if state.p1 != 0 {
				state.jump = state.p2 - state.ip
			}
			state.ip += state.jump
		},
		"06": func(state *programState, _ []int) {
			if state.p1 == 0 {
				state.jump = state.p2 - state.ip
			}
			state.ip += state.jump
		},
		"07": func(state *programState, input []int) {
			input[state.p3] = 0
			if state.p1 < state.p2 {
				input[state.p3] = 1
			}
			state.ip += state.jump
		},
		"08": func(state *programState, input []int) {
			if state.p3 < len(input) {
				input[state.p3] = 0
				if state.p1 == state.p2 {
					input[state.p3] = 1
				}
			}
			state.ip += state.jump
		},
	}

	var state programState
	state.sysID = inputVal
	state.jump = 3

	runInstructions(input, opCodes, state)
}

func main() {
	input := []int{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1102, 91, 92, 225, 1102, 85, 13, 225, 1, 47, 17, 224, 101, -176, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 7, 224, 1, 223, 224, 223, 1102, 79, 43, 225, 1102, 91, 79, 225, 1101, 94, 61, 225, 1002, 99, 42, 224, 1001, 224, -1890, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 6, 224, 1, 224, 223, 223, 102, 77, 52, 224, 1001, 224, -4697, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 7, 224, 1, 224, 223, 223, 1101, 45, 47, 225, 1001, 43, 93, 224, 1001, 224, -172, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 1, 224, 1, 224, 223, 223, 1102, 53, 88, 225, 1101, 64, 75, 225, 2, 14, 129, 224, 101, -5888, 224, 224, 4, 224, 102, 8, 223, 223, 101, 6, 224, 224, 1, 223, 224, 223, 101, 60, 126, 224, 101, -148, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 2, 224, 1, 224, 223, 223, 1102, 82, 56, 224, 1001, 224, -4592, 224, 4, 224, 1002, 223, 8, 223, 101, 4, 224, 224, 1, 224, 223, 223, 1101, 22, 82, 224, 1001, 224, -104, 224, 4, 224, 1002, 223, 8, 223, 101, 4, 224, 224, 1, 223, 224, 223, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 8, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 329, 1001, 223, 1, 223, 1007, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 344, 101, 1, 223, 223, 108, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 359, 1001, 223, 1, 223, 107, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 374, 101, 1, 223, 223, 8, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 389, 1001, 223, 1, 223, 1008, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 404, 101, 1, 223, 223, 7, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 419, 101, 1, 223, 223, 1108, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 434, 101, 1, 223, 223, 1108, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 449, 1001, 223, 1, 223, 107, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 464, 101, 1, 223, 223, 1007, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 479, 101, 1, 223, 223, 1007, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 494, 1001, 223, 1, 223, 1008, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 509, 1001, 223, 1, 223, 1108, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 524, 1001, 223, 1, 223, 108, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 539, 101, 1, 223, 223, 108, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 554, 101, 1, 223, 223, 1008, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 569, 1001, 223, 1, 223, 1107, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 584, 1001, 223, 1, 223, 7, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 599, 1001, 223, 1, 223, 8, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 614, 1001, 223, 1, 223, 7, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 629, 101, 1, 223, 223, 1107, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 644, 1001, 223, 1, 223, 1107, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 659, 1001, 223, 1, 223, 107, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 674, 101, 1, 223, 223, 4, 223, 99, 226}

	input = []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}

	// 4 -> 44
	// 44 -> 484
	// 474 -> 5324
	input[1] = 4
	input[2] = 0
	doInstructions(input, 4)
}
