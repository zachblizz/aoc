package utils

import (
	"bufio"
	"log"
	"os"
)

// ReadFile - reads a file, and returns a slice of strings
func ReadFile(filename string) *[]string {
	var ret []string

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}

	return &ret
}
