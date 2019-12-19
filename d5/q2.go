package main

import (
	"fmt"
	"strconv"

)

func one(a, b, pos, _ int, input []int) {
	input[pos] = a + b
}

func two(a, b, pos, _ int, input []int) {
	input[pos] = a * b
}

func three(pos, _, _, val int, input []int) {
	input[pos] = val
}

func four(pos, _, _, _ int, input []int) {
	if pos < len(input) {
		fmt.Printf("val @%v is %v\n", pos, input[pos])
	} else {
		fmt.Printf("@ pos %v\n", pos)
	}
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
func basicIntCodeComp(input []int, opCodes map[string]interface{}, inputVal int) []int {
	for ip := 0; ip < len(input) && input[ip] != 99; {
		op := input[ip]

		if op != 99 {
			opModes := getOpModes(op)
			var a, b, pos int

			if opModes.code != "09" {
				a = input[ip+1] // imidiate mode assumed

				if opModes.code != "03" {
					b = input[ip+2]
				}

				if opModes.code == "01" || opModes.code == "02" || opModes.code == "07" || opModes.code == "08" {
					pos = input[ip+3]
				}
			}

			// printOpMode(opModes, op, ip, input)
			if opModes.modeOne == 0 && opModes.code != "03" && opModes.code != "04" && a < len(input) { // position mode swap
				a = input[a]
			}

			if opModes.modeTwo == 0 && b < len(input) {
				b = input[b]
			}

			if _, ok := opCodes[opModes.code]; ok {
				opCodes[opModes.code].(func(int, int, int, int, []int))(a, b, pos, inputVal, input)
				ip += opModes.jump
			} else {
				switch opModes.code {
				case "05":
					if a != 0 {
						opModes.jump = b - ip
					}

					ip += opModes.jump
					break
				case "06":
					if a == 0 {
						opModes.jump = b - ip
					}

					ip += opModes.jump
					break
				case "07":
					input[pos] = 0
					if a < b {
						input[pos] = 1
					}

					ip += opModes.jump
					break
				case "08":
					if pos < len(input) {
						input[pos] = 0
						if a == b {
							input[pos] = 1
						}
					}

					ip += opModes.jump
					break
				default:
					ip++
					break
				}
			}
		} else {
			break
		}
	}

	return input
}

func printOpMode(opM opMode, op, ip int, input []int) {
	fmt.Printf("orig code: %v; op: %v; m1: %v; m2: %v; m3: %v; jmp: %v\n",
		op, opM.code, opM.modeOne, opM.modeTwo, opM.modeThree, opM.jump)
	fmt.Println(input[ip], input[ip+1], input[ip+2], input[ip+3])
}

func getOpModes(op int) opMode {
	strOp := strconv.Itoa(op)
	code := opMode{strOp, 0, 0, 0, 3}

	if len(strOp) < 4 {
		strOp = fmt.Sprintf("%04v", op)
	}

	c := []rune(strOp)
	code.code = string(c[2:4])

	code.modeOne, _ = strconv.Atoi(string(c[1:2]))
	code.modeTwo, _ = strconv.Atoi(string(c[0:1]))

	if code.code == "01" || code.code == "02" || code.code == "07" || code.code == "08" {
		code.jump = 4
	} else if code.code == "03" || code.code == "04" {
		code.jump = 2
	}

	return code
}

func doInstructions(input []int, inputVal int) {
	opCodes := map[string]interface{}{
		"01": one,
		"02": two,
		"03": three,
		"04": four,
	}

	basicIntCodeComp(input, opCodes, inputVal)
}

func main() {
	input := []int{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1102, 91, 92, 225, 1102, 85, 13, 225, 1, 47, 17, 224, 101, -176, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 7, 224, 1, 223, 224, 223, 1102, 79, 43, 225, 1102, 91, 79, 225, 1101, 94, 61, 225, 1002, 99, 42, 224, 1001, 224, -1890, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 6, 224, 1, 224, 223, 223, 102, 77, 52, 224, 1001, 224, -4697, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 7, 224, 1, 224, 223, 223, 1101, 45, 47, 225, 1001, 43, 93, 224, 1001, 224, -172, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 1, 224, 1, 224, 223, 223, 1102, 53, 88, 225, 1101, 64, 75, 225, 2, 14, 129, 224, 101, -5888, 224, 224, 4, 224, 102, 8, 223, 223, 101, 6, 224, 224, 1, 223, 224, 223, 101, 60, 126, 224, 101, -148, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 2, 224, 1, 224, 223, 223, 1102, 82, 56, 224, 1001, 224, -4592, 224, 4, 224, 1002, 223, 8, 223, 101, 4, 224, 224, 1, 224, 223, 223, 1101, 22, 82, 224, 1001, 224, -104, 224, 4, 224, 1002, 223, 8, 223, 101, 4, 224, 224, 1, 223, 224, 223, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 8, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 329, 1001, 223, 1, 223, 1007, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 344, 101, 1, 223, 223, 108, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 359, 1001, 223, 1, 223, 107, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 374, 101, 1, 223, 223, 8, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 389, 1001, 223, 1, 223, 1008, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 404, 101, 1, 223, 223, 7, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 419, 101, 1, 223, 223, 1108, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 434, 101, 1, 223, 223, 1108, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 449, 1001, 223, 1, 223, 107, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 464, 101, 1, 223, 223, 1007, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 479, 101, 1, 223, 223, 1007, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 494, 1001, 223, 1, 223, 1008, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 509, 1001, 223, 1, 223, 1108, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 524, 1001, 223, 1, 223, 108, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 539, 101, 1, 223, 223, 108, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 554, 101, 1, 223, 223, 1008, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 569, 1001, 223, 1, 223, 1107, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 584, 1001, 223, 1, 223, 7, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 599, 1001, 223, 1, 223, 8, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 614, 1001, 223, 1, 223, 7, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 629, 101, 1, 223, 223, 1107, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 644, 1001, 223, 1, 223, 1107, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 659, 1001, 223, 1, 223, 107, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 674, 101, 1, 223, 223, 4, 223, 99, 226}

	// < 8 -> 999
	// 8 -> 1000
	// > 8 -> 1001
	// input = []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
	// 	1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
	// 	999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}

	// 0 -> 0
	// !0 -> 1
	// input = []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	// input = []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}

	// input = []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8} // if input == 8 -> 1 else -> 0
	// input = []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8} // if input < 8 -> 1 else -> 0

	// input = []int{3, 3, 1108, -1, 8, 3, 4, 3, 99} // if input == 8 -> 1 else -> 0
	// input = []int{3, 3, 1107, -1, 8, 3, 4, 3, 99} // if input < 8 -> 1; else -> 0

	doInstructions(input, 5)
}
