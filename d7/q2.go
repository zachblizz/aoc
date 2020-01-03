package main

import (
	"fmt"
	"math"

	utils "github.com/zachblizz/aoc/utils"

)

// modes
// 0 - position
// 1 - imediate
func runInstructions(input []int, opCodes map[string]interface{}, state *utils.ProgramState) int {
	var output int
	for state.IP < len(input) && input[state.IP] != 99 {
		op := input[state.IP]
		state.GetCurrState(op, input)

		if _, ok := opCodes[state.Code]; ok {
			opCodes[state.Code].(func([]int))(input)
		} else {
			state.IP++
		}
	}

	return output
}

func runSequences(opCodes map[string]interface{}, input []int, seq1, seq2 [][]int, state *utils.ProgramState, maxFound *int) {
	ampMap := map[int]string{
		0: "A",
		1: "B",
		2: "C",
		3: "D",
		4: "E",
	}

	for _, seq := range seq1 {
		// this runs through A to E for given sequence
		for i, inputOne := range seq {
			if i == 0 && state.SendZeroSignal {
				// fmt.Println("sending zero signal")
				state.SendZeroSignal = false
				state.SysID = []int{inputOne, 0}
			} else {
				state.SysID = []int{inputOne, state.Output}
			}

			c := make([]int, len(input))
			copy(c, input)

			state.ResetState()
			runInstructions(c, opCodes, state)

			if state.LoopMode == 0 {
				fmt.Printf("amp %v output: %v seq: %v max: %v\n", ampMap[i], state.Output, seq, *maxFound)
			}
		} // state.Output should be from E at this point...

		// here is where we need to switch the loopmode to 1...
		if state.LoopMode == 0 {
			fmt.Printf("\n")
			state.LoopMode = 1
			runSequences(opCodes, input, seq2, seq1, state, maxFound)
		} else if *maxFound < state.Output {
			*maxFound = state.Output
		}
	}

	state.LoopMode = 0 // reset loop mode

	return // go for the next [0,1,2,3,4] sequence
}

func main() {
	state := utils.NewState()
	state.LoopMode = 0

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

	input := []int{3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 46, 67, 76, 101, 118, 199, 280, 361, 442, 99999, 3, 9, 1002, 9, 4, 9, 1001, 9, 2, 9, 102, 3, 9, 9, 101, 3, 9, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 3, 9, 102, 2, 9, 9, 1001, 9, 2, 9, 1002, 9, 3, 9, 4, 9, 99, 3, 9, 101, 3, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 2, 9, 1002, 9, 5, 9, 101, 5, 9, 9, 1002, 9, 4, 9, 101, 5, 9, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 1001, 9, 5, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99}

	// tests
	// should get - 139629729
	input = []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5}

	// input = []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}

	sequences := utils.GetSequences([]int{0, 1, 2, 3, 4})
	sequences2 := utils.GetSequences([]int{5, 6, 7, 8, 9})
	maxFound := math.MinInt64

	runSequences(opCodes, input, sequences, sequences2, state, &maxFound)

	fmt.Printf("max found: %v\n", maxFound)
}
