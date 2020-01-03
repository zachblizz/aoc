package utils

import "fmt"

// ProgramState - the state of the intcode program
type ProgramState struct {
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

// One - does the needful for opcode one
func (state *ProgramState) One(input []int) {
	input[state.p3] = state.p1 + state.p2
	state.ip += state.jump
}

// Two - does the needful for opcode two
func (state *ProgramState) Two(input []int) {
	input[state.p3] = state.p1 * state.p2
	state.ip += state.jump
}

// Three - does the needful for opcode three
func (state *ProgramState) Three(input []int) {
	input[state.p1] = state.sysID[state.sysIDptr]
	state.sysIDptr++
	state.ip += state.jump
}

// Four - does the needful for opcode Four
func (state *ProgramState) Four(input []int) {
	if state.p1 < len(input) {
		state.output = input[state.p1]
	} else {
		fmt.Printf("oob error... @ pos %v\n", state.p1)
		state.output = state.p1
	}
	state.ip += state.jump
}

// Five - does the needful for opcode Five
func (state *ProgramState) Five(_ []int) {
	if state.p1 != 0 {
		state.jump = state.p2 - state.ip
	}
	state.ip += state.jump
}

// Six - does the needful for opcode Six
func (state *ProgramState) Six(_ []int) {
	if state.p1 == 0 {
		state.jump = state.p2 - state.ip
	}
	state.ip += state.jump
}

// Seven - does the needful for opcode Seven
func (state *ProgramState) Seven(input []int) {
	input[state.p3] = 0
	if state.p1 < state.p2 {
		input[state.p3] = 1
	}
	state.ip += state.jump
}

// Eight - does the needful for opcode Eight
func (state *ProgramState) Eight(input []int) {
	if state.p3 < len(input) {
		input[state.p3] = 0
		if state.p1 == state.p2 {
			input[state.p3] = 1
		}
	}
	state.ip += state.jump
}

// ClearStateModeAndParams - clears the modes and params
func (state *ProgramState) ClearStateModeAndParams() {
	state.p1 = 0
	state.p2 = 0
	state.p3 = 0
	state.modeOne = 0
	state.modeTwo = 0
	state.modeThree = 0
	state.jump = 3
}

// ResetState - resets the instruction pointer and the sysIDPtr
func (state *ProgramState) ResetState() {
	state.ClearStateModeAndParams()
	state.ip = 0
	state.sysIDptr = 0
}

// GetParams - gets the input params (p1,2,3)
func (state *ProgramState) GetParams(input []int) {
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

// NewState - creates a ProgramState pointer
func NewState() *ProgramState {
	state := ProgramState{}
	state.ResetState()

	return &state
}
