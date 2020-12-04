package main

import (
	utils "aoc-2020/utils"
	"fmt"
	"strconv"
	"strings"
)

const validFields = 7

func main() {
	// cid: country id is not important

	input := utils.ReadFile("input.txt")
	passports := getAllPassports(input)
	// partOne(passports)
	partTwo(passports)
}

// off by one?
func partOne(passports []string) {
	importantFields := map[string]bool{
		"byr": true,
		"iyr": true,
		"eyr": true,
		"hgt": true,
		"hcl": true,
		"ecl": true,
		"pid": true,
	}
	validPP := 0

	for _, pp := range passports {
		items := strings.Split(pp, " ")
		goodField := 0

		for _, item := range items {
			if strings.Trim(item, " ") != "" {
				key, _ := getItemKeyVal(item)

				if _, ok := importantFields[key]; ok {
					goodField++
				}
			}
		}

		if goodField == validFields {
			validPP++
		} else {
			fmt.Printf("bad pp: %v\n", pp)
		}
	}

	fmt.Println(validPP)
}

// off by one again?
func partTwo(passports []string) {
	importantFields := map[string]interface{}{
		"byr": func(byrStr string) bool {
			byr, _ := strconv.Atoi(byrStr)
			return byr >= 1920 && byr <= 2002
		},
		"iyr": func(iyrStr string) bool {
			iyr, _ := strconv.Atoi(iyrStr)
			return iyr >= 2010 && iyr <= 2020
		},
		"eyr": func(eyrStr string) bool {
			eyr, _ := strconv.Atoi(eyrStr)
			return eyr >= 2020 && eyr <= 2030
		},
		"hgt": func(hgtStr string) bool {
			mIdx := len(hgtStr) - 2
			metric := hgtStr[mIdx:]
			length, _ := strconv.Atoi(hgtStr[0: mIdx])

			if metric == "cm" {
				return length >= 150 && length <= 193
			} else if metric == "in" {
				return length >= 59 && length <= 76
			}

			return false
		},
		"hcl": func(hcl string) bool {
			if hcl[0:1] != "#" || len(hcl) != 7 {
				return false
			}

			for _, c := range hcl[1:] {
				isNum := c >= '0' && c <= '9'
				isAtF := c >= 'a' && c <= 'f'
				if (!isNum && !isAtF) {
					return false
				}
			}

			return true
		},
		"ecl": func(ecl string) bool {
			colors := map[string]bool {
				"amb": true,
				"blu": true,
				"brn": true,
				"gry": true,
				"grn": true,
				"hzl": true,
				"oth": true,
			}
			_, ok := colors[ecl]

			return ok
		},
		"pid": func(pid string) bool {
			if len(pid) != 9 {
				return false
			}

			for _, c := range pid {
				if c < '0' || c > '9' {
					return false
				}
			}

			return true
		},
		"cid": func(_ string) bool { return false },
	}

	validPP := 0

	for _, pp := range passports {
		items := strings.Split(pp, " ")
		goodField := 0
		reason := ""

		for _, item := range items {
			if strings.Trim(item, " ") != "" {
				key, val := getItemKeyVal(item)

				if verify, ok := importantFields[key]; ok && verify.(func(string)bool)(val) {
					goodField++
				} else if key != "cid" {
					reason = fmt.Sprintf("%v is bad with %v", key, val)
				}
			}
		}

		if reason == "" && goodField == validFields {
			validPP++
		} else {
			if reason == "" && goodField < validFields {
				reason = "not enough fields"
			}

			fmt.Printf("bad pp: %v\nreason: %v\n\n", pp, reason)
		}
	}

	fmt.Println(validPP)
}

func getAllPassports(input *[]string) []string {
	var passports []string
	var passport strings.Builder

	for _, row := range *input {
		if row == "" {
			passports = append(passports, passport.String())
			passport.Reset()
		} else {
			passport.WriteString(row)
			passport.WriteString(" ")
		}
	}

	return passports
}

func getItemKeyVal(item string) (string, string) {
	itemSplit := strings.Split(item, ":")
	return itemSplit[0], itemSplit[1]
}
