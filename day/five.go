// Copy and paste this each day to get started
package day

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func Five(filepath string) []string {
	f, err := os.Open(filepath)
	check(err)
	defer f.Close()
	sc := bufio.NewScanner(f)

	//regex to confirm we're working with letters later
	var IsLetter = regexp.MustCompile(`^[A-Z]+$`).MatchString

	// slices we'll be using for our output
	var set [][]string
	var output []string
	var digits []int

	// calculate column values

	// column1 := []string{"Z", "N"}
	// column2 := []string{"M", "C", "D"}
	// column3 := []string{"P"}

	// set = append(set, column1)
	// set = append(set, column2)
	// set = append(set, column3)

	for sc.Scan() {

		// blank out the digits slice
		digits = nil
		crateCount := 9
		for i := 0; i < crateCount; i++ {
			set = append(set, []string{})
		}

		// read the not-blank line
		if sc.Text() != "" {
			//			[D] [S] [R] [D] [G] [F] [S] [L] [Q]
			if strings.Contains(sc.Text(), "[") {
				crateContent := strings.Split(sc.Text(), "")
				for i := 0; i < crateCount; i++ {
					crateColumnCounter := (i * 4) + 1
					if crateColumnCounter >= len(crateContent) {
						break
					}
					if IsLetter(crateContent[crateColumnCounter]) {
						set[i] = append(set[i], strings.TrimSpace(crateContent[crateColumnCounter]))
					}
				}
			}

			//read each move and execute it
			if strings.HasPrefix(sc.Text(), "move") {
				move := strings.Split(sc.Text(), " ")
				for _, v := range move {
					if isInt(v) {
						value, _ := strconv.Atoi(v)
						digits = append(digits, value)
					}
				} // do the moves
				moveCount := digits[0]
				moveFrom := digits[1] - 1
				moveTo := digits[2] - 1
				for i := 0; i < moveCount; i++ {
					set[moveTo] = append(set[moveTo], set[moveFrom][len(set[moveFrom])-1])
					set[moveFrom] = set[moveFrom][:len(set[moveFrom])-1]
				}
			}
		} else {
			for i := 0; i < crateCount; i++ {
				set[i] = reverse(set[i])
			}
		}
	}

	for _, s := range set {
		if len(s) > 0 {
			output = append(output, s[len(s)-1])
		}
	}

	return output
}

func FivePart2(filepath string) []string {
	f, err := os.Open(filepath)
	check(err)
	defer f.Close()
	sc := bufio.NewScanner(f)

	//regex to confirm we're working with letters later
	var IsLetter = regexp.MustCompile(`^[A-Z]+$`).MatchString

	// slices we'll be using for our output
	var set [][]string
	var output []string
	var digits []int

	// calculate column values

	// column1 := []string{"Z", "N"}
	// column2 := []string{"M", "C", "D"}
	// column3 := []string{"P"}

	// set = append(set, column1)
	// set = append(set, column2)
	// set = append(set, column3)

	for sc.Scan() {

		// blank out the digits slice
		digits = nil
		crateCount := 9
		for i := 0; i < crateCount; i++ {
			set = append(set, []string{})
		}

		// read the not-blank line
		if sc.Text() != "" {
			//			[D] [S] [R] [D] [G] [F] [S] [L] [Q]
			if strings.Contains(sc.Text(), "[") {
				crateContent := strings.Split(sc.Text(), "")
				for i := 0; i < crateCount; i++ {
					crateColumnCounter := (i * 4) + 1
					if crateColumnCounter >= len(crateContent) {
						break
					}
					if IsLetter(crateContent[crateColumnCounter]) {
						set[i] = append(set[i], strings.TrimSpace(crateContent[crateColumnCounter]))
					}
				}
			}

			//read each move and execute it
			if strings.HasPrefix(sc.Text(), "move") {
				move := strings.Split(sc.Text(), " ")
				for _, v := range move {
					if isInt(v) {
						value, _ := strconv.Atoi(v)
						digits = append(digits, value)
					}
				} // do the moves
				moveCount := digits[0]
				moveFrom := digits[1] - 1
				moveTo := digits[2] - 1
				fmt.Printf("Move %v from %v to %v", moveCount, moveFrom, moveTo)
				fmt.Println()
				for i := 0; i < moveCount; i++ {
					set[moveTo] = append(set[moveTo], set[moveFrom][len(set[moveFrom])-(moveCount-i)])
				}
				for i := 0; i < moveCount; i++ {
					set[moveFrom] = set[moveFrom][:len(set[moveFrom])-1]
				}
			}
		} else {
			for i := 0; i < crateCount; i++ {
				set[i] = reverse(set[i])
			}
		}
	}

	for _, s := range set {
		if len(s) > 0 {
			output = append(output, s[len(s)-1])
		}
	}

	return output
}

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func reverse(input []string) []string {
	inputLen := len(input)
	output := make([]string, inputLen)

	for i, n := range input {
		j := inputLen - i - 1

		output[j] = n
	}

	return output
}
