package utils

import (
	"bytes"
	"strconv"

)

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

// GetSequences - gets all sequences for given slice
func GetSequences(sequence []int) [][]int {
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
