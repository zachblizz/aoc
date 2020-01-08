package main

import (
	"fmt"
	"math"

	utils "github.com/zachblizz/aoc/utils"
)

// modes
// 0 - position
// 1 - imediate
func runInstructions(input []int, opCodes map[string]interface{}, state *utils.ProgramState) {
	for state.IP < len(input) {
		op := input[state.IP]
		state.GetCurrState(op, input)

		if state.Code == "99" {
			state.Halted = true
			return
		}

		if _, ok := opCodes[state.Code]; ok {
			opCodes[state.Code].(func([]int))(input)
		} else {
			state.IP++
		}
	}
}

func runSequences(opCodes map[string]interface{}, input []int,
	sequence [][]int, state *utils.ProgramState, maxFound *int) {

	ampNames := map[int]string{
		0: "A",
		1: "B",
		2: "C",
		3: "D",
		4: "E",
	}

	for _, seq := range sequence {
		// this runs through A to E for given sequence
		for amp, inputOne := range seq {
			if amp == 0 && state.SendZeroSignal {
				fmt.Printf("initial amp, passing input signals [%v, 0]\n", inputOne)
				state.InputIns = []int{inputOne, 0}
				state.SendZeroSignal = false
			} else {
				ampName := ampNames[amp-1]
				if ampName == "" {
					ampName = "A"
				}

				fmt.Printf("amp %v's output: %v, which is amp %v's input\n", ampName, state.Output, ampNames[amp])
				fmt.Printf("input signals: [%v, %v]\n", inputOne, state.Output)
				state.InputIns = []int{inputOne, state.Output}
			}

			c := make([]int, len(input))
			copy(c, input)

			state.ResetState()
			runInstructions(c, opCodes, state)
		} // state.Output should be from E at this point...

		if *maxFound < state.Output {
			fmt.Printf("prevmax: %v;  new max: %v\n", *maxFound, state.Output)
			*maxFound = state.Output
		}
	}
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
	// should get - 139629729, from [9,8,7,6,5]
	input = []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5}

	// input = []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}

	sequences := utils.GetSequences([]int{5, 6, 7, 8, 9})
	maxFound := math.MinInt64

	runSequences(opCodes, input, sequences, state, &maxFound)
	fmt.Printf("max found: %v\n", maxFound)
}
