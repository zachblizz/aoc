package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"

	ps "github.com/zachblizz/aoc/IntCodeUtils"
)

// modes
// 0 - position
// 1 - imediate
func runInstructions(input []int, opCodes map[string]interface{}, state *ps.ProgramState) int {
	var output int
	for state.ip < len(input) && input[state.ip] != 99 {
		op := input[state.ip]
		getCurrState(op, state, input)

		if _, ok := opCodes[state.code]; ok {
			opCodes[state.code].(func([]int))(state, input)
		} else {
			state.ip++
		}
	}

	return output
}


func getCurrState(op int, state *ps.ProgramState, input []int) {
	strOp := strconv.Itoa(op)
	state.ClearStateModeAndParams()

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

	state.GetParams(input)
}

func doInstructions(input []int, state *ps.ProgramState) {
	opCodes := map[string]interface{}{
		"01": state.One,
		"02": state.Two,
		"03": state.Three,
		"04": state.Four,
		"05": state.Five,
		"06": state.Six,
		"07": state.Seven,
		"08": state.Eight,
	}

	runInstructions(input, opCodes, state)
}

func swapIndecies(s []int, i, j int) {
	swap := s[i]

	s[i] = s[j]
	s[j] = swap
}

func getSeqKey(sequence []int) string {
	var key bytes.Buffer
	for j := 0; j < len(sequence); j++ {
		key.WriteString(strconv.Itoa(sequence[j]))
	}
	return key.String()
}

// thank you - https://www.codesdope.com/blog/article/generating-permutations-of-all-elements-of-an-arra/
func getPermutations(s []int, m map[string][]int, start, end int) {
	if start == end {
		c := make([]int, len(s))
		copy(c, s)

		k := getSeqKey(c)
		m[k] = c

		return
	}

	for i := start; i <= end; i++ {
		swapIndecies(s, i, start)
		getPermutations(s, m, start+1, end)
		swapIndecies(s, i, start)
	}
}

func getSequences(sequence []int) [][]int {
	var sequences [][]int
	seqMap := make(map[string][]int)

	c := make([]int, len(sequence))
	copy(c, sequence)
	getPermutations(c, seqMap, 0, len(sequence)-1)

	for _, seq := range seqMap {
		sequences = append(sequences, seq)
	}

	return sequences
}

func main() {
	input := []int{3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 46, 67, 76, 101, 118, 199, 280, 361, 442, 99999, 3, 9, 1002, 9, 4, 9, 1001, 9, 2, 9, 102, 3, 9, 9, 101, 3, 9, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 3, 9, 102, 2, 9, 9, 1001, 9, 2, 9, 1002, 9, 3, 9, 4, 9, 99, 3, 9, 101, 3, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 2, 9, 1002, 9, 5, 9, 101, 5, 9, 9, 1002, 9, 4, 9, 101, 5, 9, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 1001, 9, 5, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99}

	// input = []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}

	state := ps.NewState()
	sequences := getSequences([]int{0, 1, 2, 3, 4})

	maxFound := math.MinInt64
	for _, seq := range sequences {
		for i := 0; i < len(seq); i++ {
			if i == 0 {
				state.sysID = []int{seq[i], 0}
			} else {
				state.sysID = []int{seq[i], state.output}
			}

			c := make([]int, len(input))
			copy(c, input)

			state.ResetState()
			doInstructions(c, state)
		}

		if maxFound < state.output {
			maxFound = state.output
		}

		// fmt.Printf("seq: %v, state.output: %v\n", seq, state.output)
	}

	fmt.Printf("max found: %v\n", maxFound)
}
