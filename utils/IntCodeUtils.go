package utils

import (
	"fmt"
	"strconv"

)

// ProgramState - the state of the intcode program
type ProgramState struct {
	Code string // op code

	SysID    []int // system check id
	SysIDptr int

	IP int // instruction pointer
	P1 int
	P2 int
	P3 int

	ModeOne   int
	ModeTwo   int
	ModeThree int

	Jump   int
	Output int

	LoopMode int // indecates if we're in the feedback loop (0 - orig mode, 1 - feedback mode)

	SendZeroSignal bool
}

// One - does the needful for opcode one
func (state *ProgramState) One(input []int) {
	input[state.P3] = state.P1 + state.P2
	state.IP += state.Jump
}

// Two - does the needful for opcode two
func (state *ProgramState) Two(input []int) {
	input[state.P3] = state.P1 * state.P2
	state.IP += state.Jump
}

// Three - does the needful for opcode three
func (state *ProgramState) Three(input []int) {
	input[state.P1] = state.SysID[state.SysIDptr%2]
	state.SysIDptr++
	state.IP += state.Jump
}

// Four - does the needful for opcode Four
func (state *ProgramState) Four(input []int) {
	if state.P1 < len(input) {
		state.Output = input[state.P1]
	} else {
		fmt.Printf("oob error... @ pos %v\n", state.P1)
		state.Output = state.P1
	}
	state.IP += state.Jump
}

// Five - does the needful for opcode Five
func (state *ProgramState) Five(_ []int) {
	if state.P1 != 0 {
		state.Jump = state.P2 - state.IP
	}
	state.IP += state.Jump
}

// Six - does the needful for opcode Six
func (state *ProgramState) Six(_ []int) {
	if state.P1 == 0 {
		state.Jump = state.P2 - state.IP
	}
	state.IP += state.Jump
}

// Seven - does the needful for opcode Seven
func (state *ProgramState) Seven(input []int) {
	input[state.P3] = 0
	if state.P1 < state.P2 {
		input[state.P3] = 1
	}
	state.IP += state.Jump
}

// Eight - does the needful for opcode Eight
func (state *ProgramState) Eight(input []int) {
	if state.P3 < len(input) {
		input[state.P3] = 0
		if state.P1 == state.P2 {
			input[state.P3] = 1
		}
	}
	state.IP += state.Jump
}

// ClearStateModeAndParams - clears the modes and params
func (state *ProgramState) ClearStateModeAndParams() {
	state.P1 = 0
	state.P2 = 0
	state.P3 = 0
	state.ModeOne = 0
	state.ModeTwo = 0
	state.ModeThree = 0
	state.Jump = 3
}

// ResetState - resets the instruction pointer and the SysIDPtr
func (state *ProgramState) ResetState() {
	state.ClearStateModeAndParams()
	state.IP = 0
	state.SysIDptr = 0
}

// GetParams - gets the input params (P1,2,3)
func (state *ProgramState) GetParams(input []int) {
	if state.Code != "09" {
		state.P1 = input[state.IP+1] // imidiate mode assumed

		if state.Code != "03" {
			state.P2 = input[state.IP+2]
		}

		if state.Code == "01" || state.Code == "02" || state.Code == "07" || state.Code == "08" {
			state.P3 = input[state.IP+3]
		}
	}

	// position mode swap
	if state.ModeOne == 0 && state.Code != "03" && state.Code != "04" && state.P1 < len(input) {
		state.P1 = input[state.P1]
	}

	if state.ModeTwo == 0 && state.P2 < len(input) {
		state.P2 = input[state.P2]
	}
}

// GetCurrState - gets the current state of the intcode program
func (state *ProgramState) GetCurrState(op int, input []int) {
	strOp := strconv.Itoa(op)
	state.ClearStateModeAndParams()

	if len(strOp) < 4 {
		strOp = fmt.Sprintf("%04v", op)
	}

	c := []rune(strOp)
	state.Code = string(c[2:4])

	state.ModeOne, _ = strconv.Atoi(string(c[1:2]))
	state.ModeTwo, _ = strconv.Atoi(string(c[0:1]))

	if state.Code == "01" || state.Code == "02" || state.Code == "07" || state.Code == "08" {
		state.Jump = 4
	} else if state.Code == "03" || state.Code == "04" {
		state.Jump = 2
	}

	state.GetParams(input)
}

// NewState - creates a ProgramState pointer
func NewState() *ProgramState {
	state := ProgramState{}
	state.ResetState()
	state.SendZeroSignal = true

	return &state
}
