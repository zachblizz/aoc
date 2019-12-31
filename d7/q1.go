package main

import (
	"bytes"
	"fmt"
	"strconv"

)

type programState struct {
	code string // op code

	sysID    []int // system check id
	sysIDptr int

	ip int // instruction pointer
	p1 int
	p2 int
	p3 int

	modeOne   int
	modeTwo   int
	modeThree int

	jump   int
	output int
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
	input[state.p1] = state.sysID[state.sysIDptr]
	state.sysIDptr++
	state.ip += state.jump
}

func four(state *programState, input []int) {
	if state.p1 < len(input) {
		state.output = input[state.p1]
	} else {
		fmt.Printf("oob error... @ pos %v\n", state.p1)
		state.output = state.p1
	}
	state.ip += state.jump
}

// modes
// 0 - position
// 1 - imediate
func runInstructions(input []int, opCodes map[string]interface{}, state *programState) int {
	var output int
	for state.ip < len(input) && input[state.ip] != 99 {
		op := input[state.ip]
		getCurrState(op, state, input)

		if _, ok := opCodes[state.code]; ok {
			opCodes[state.code].(func(*programState, []int))(state, input)
		} else {
			state.ip++
		}
	}

	return output
}

func clearStateModeAndParams(s *programState) {
	s.p1 = 0
	s.p2 = 0
	s.p3 = 0
	s.modeOne = 0
	s.modeTwo = 0
	s.modeThree = 0
	s.jump = 3
}

func resetState(s *programState) {
	clearStateModeAndParams(s)
	s.ip = 0
	s.sysIDptr = 0
}

func getCurrState(op int, state *programState, input []int) {
	strOp := strconv.Itoa(op)
	clearStateModeAndParams(state)

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

func doInstructions(input []int, state *programState) {
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

	runInstructions(input, opCodes, state)
}

func scrambleSeq(s []int, i int) []int {
	if i+1 >= len(s) {
		return nil
	}

	swap := s[i+1]

	s[i+1] = s[i]
	s[i] = swap

	return s
}

func swapIndecies(s []int, i, j int) []int {
	swap := s[i]

	s[i] = s[j]
	s[j] = swap

	return s
}

func getSeqKey(sequence []int) string {
	var key bytes.Buffer
	for j := 0; j < len(sequence); j++ {
		key.WriteString(strconv.Itoa(sequence[j]))
	}
	return key.String()
}

func getSequences(sequence []int) [][]int {
	var sequences [][]int
	seqMap := make(map[string][]int)

	for _, seq := range seqMap {
		sequences = append(sequences, seq)
	}

	return sequences
}

func main() {
	input := []int{3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 46, 67, 76, 101, 118, 199, 280, 361, 442, 99999, 3, 9, 1002, 9, 4, 9, 1001, 9, 2, 9, 102, 3, 9, 9, 101, 3, 9, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 3, 9, 102, 2, 9, 9, 1001, 9, 2, 9, 1002, 9, 3, 9, 4, 9, 99, 3, 9, 101, 3, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 2, 9, 1002, 9, 5, 9, 101, 5, 9, 9, 1002, 9, 4, 9, 101, 5, 9, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 1001, 9, 5, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99}

	input = []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}

	var state programState
	sequences := getSequences([]int{0, 1, 2, 3, 4})

	for _, seq := range sequences {
		for i := 0; i < len(seq); i++ {
			if i == 0 {
				state.sysID = []int{seq[i], 0}
			} else {
				state.sysID = []int{seq[i], state.output}
			}

			c := make([]int, len(input))
			copy(c, input)

			resetState(&state)
			doInstructions(c, &state)
		}

		fmt.Printf("seq: %v, state.output: %v\n", seq, state.output)
	}
}
