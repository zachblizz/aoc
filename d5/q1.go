package main

import (
	"fmt"
	"strconv"

)

func one(a, b, pos int, input []int) {
	input[pos] = a + b
}

func two(a, b, pos int, input []int) {
	input[pos] = a * b
}

func three(a, b, pos int, input []int) {
	input[a] = b
}

func four(a, b, pos int, input []int) int {
	fmt.Println(input[b])
	return input[b]
}

type opMode struct {
	code      string
	modeOne   int
	modeTwo   int
	modeThree int
	jump      int
}

// modes
// 0 - position
// 1 - imediate
func basicIntCodeComp(input []int, opCodes map[string]interface{}) []int {
	for insPtr := 0; insPtr < len(input); {
		op := input[insPtr]

		if op != 99 {
			opModes := getOpModes(op)

			if _, ok := opCodes[opModes.code]; ok {
				a := input[insPtr+1]
				b := input[insPtr+2]
				pos := input[insPtr+3]

				if opModes.modeOne == 0 {
					a = input[a]
				}

				if opModes.modeTwo == 0 {
					b = input[b]
				}

				opCodes[opModes.code].(func(int, int, int, []int))(a, b, pos, input)
			}

			insPtr += opModes.jump
		}
	}

	return input
}

func getOpModes(op int) opMode {
	var code opMode
	strOp := strconv.Itoa(op)

	if op < 1000 {
		tmpOp := ""
		for i := len(strOp); i < 4; i++ {
			tmpOp += "0"
		}
		tmpOp += strOp

		strOp = tmpOp
	}

	c := []rune(strOp)

	code.code = string(c[2:4])
	code.modeOne, _ = strconv.Atoi(string(c[1:2]))
	code.modeTwo, _ = strconv.Atoi(string(c[0:1]))
	code.modeThree = 0
	code.jump = len(strOp)

	return code
}

func doInstructions(input []int, stopCode int) int {
	opCodes := map[string]interface{}{
		"01": one,
		"02": two,
		"03": three,
		"04": four,
	}

	tmpInput := make([]int, len(input))
	copy(tmpInput, input)

	fmt.Println(basicIntCodeComp(tmpInput, opCodes)[0])

	return 0
}

func main() {
	input := []int{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1102, 91, 92, 225, 1102, 85, 13, 225, 1, 47, 17, 224, 101, -176, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 7, 224, 1, 223, 224, 223, 1102, 79, 43, 225, 1102, 91, 79, 225, 1101, 94, 61, 225, 1002, 99, 42, 224, 1001, 224, -1890, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 6, 224, 1, 224, 223, 223, 102, 77, 52, 224, 1001, 224, -4697, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 7, 224, 1, 224, 223, 223, 1101, 45, 47, 225, 1001, 43, 93, 224, 1001, 224, -172, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 1, 224, 1, 224, 223, 223, 1102, 53, 88, 225, 1101, 64, 75, 225, 2, 14, 129, 224, 101, -5888, 224, 224, 4, 224, 102, 8, 223, 223, 101, 6, 224, 224, 1, 223, 224, 223, 101, 60, 126, 224, 101, -148, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 2, 224, 1, 224, 223, 223, 1102, 82, 56, 224, 1001, 224, -4592, 224, 4, 224, 1002, 223, 8, 223, 101, 4, 224, 224, 1, 224, 223, 223, 1101, 22, 82, 224, 1001, 224, -104, 224, 4, 224, 1002, 223, 8, 223, 101, 4, 224, 224, 1, 223, 224, 223, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 8, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 329, 1001, 223, 1, 223, 1007, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 344, 101, 1, 223, 223, 108, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 359, 1001, 223, 1, 223, 107, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 374, 101, 1, 223, 223, 8, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 389, 1001, 223, 1, 223, 1008, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 404, 101, 1, 223, 223, 7, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 419, 101, 1, 223, 223, 1108, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 434, 101, 1, 223, 223, 1108, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 449, 1001, 223, 1, 223, 107, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 464, 101, 1, 223, 223, 1007, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 479, 101, 1, 223, 223, 1007, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 494, 1001, 223, 1, 223, 1008, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 509, 1001, 223, 1, 223, 1108, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 524, 1001, 223, 1, 223, 108, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 539, 101, 1, 223, 223, 108, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 554, 101, 1, 223, 223, 1008, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 569, 1001, 223, 1, 223, 1107, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 584, 1001, 223, 1, 223, 7, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 599, 1001, 223, 1, 223, 8, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 614, 1001, 223, 1, 223, 7, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 629, 101, 1, 223, 223, 1107, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 644, 1001, 223, 1, 223, 1107, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 659, 1001, 223, 1, 223, 107, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 674, 101, 1, 223, 223, 4, 223, 99, 226}
	// input := []int{1002, 4, 3, 4, 33}

	fmt.Println(doInstructions(input, 19690720))
}
