// Copy and paste this each day to get started
package day

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Six(filepath string) int {
	f, err := os.Open(filepath)
	check(err)
	defer f.Close()
	sc := bufio.NewScanner(f)
	result := 0
	start := ""
	for sc.Scan() {
		message := sc.Text()
		fmt.Println(message)
		for i, c := range message {
			if i > len(message)-4 {
				break
			} else {
				start = string(c) + string(message[i+1]) + string(message[i+2]) + string(message[i+3])
				if !checkDups(start, string(c)) && !checkDups(start, string(message[i+1])) && !checkDups(start, string(message[i+2])) {
					result = i + 4
					break
				}
			}
		}
		fmt.Println(start, result)
	}
	return result
}

func Six2(filepath string) int {
	f, err := os.Open(filepath)
	check(err)
	defer f.Close()
	sc := bufio.NewScanner(f)
	result := 0
	start := ""
	for sc.Scan() {
		message := sc.Text()
		fmt.Println(message)
		for i, c := range message {
			if i > len(message)-4 {
				break
			} else {
				start = string(c) + string(message[i+1]) + string(message[i+2]) + string(message[i+3]) + string(message[i+4]) + string(message[i+5]) + string(message[i+6]) + string(message[i+7]) + string(message[i+8]) + string(message[i+9]) + string(message[i+10]) + string(message[i+11]) + string(message[i+12]) + string(message[i+13])
				if !checkDups(start, string(c)) && !checkDups(start, string(message[i+1])) && !checkDups(start, string(message[i+2])) && !checkDups(start, string(message[i+3])) && !checkDups(start, string(message[i+4])) && !checkDups(start, string(message[i+5])) && !checkDups(start, string(message[i+6])) && !checkDups(start, string(message[i+7])) && !checkDups(start, string(message[i+8])) && !checkDups(start, string(message[i+9])) && !checkDups(start, string(message[i+10])) && !checkDups(start, string(message[i+11])) && !checkDups(start, string(message[i+12])) && !checkDups(start, string(message[i+13])) {
					result = i + 14
					break
				}
			}
		}
		fmt.Println(start, result)
	}
	return result
}

func SixPart2(filepath string) int {
	f, err := os.Open(filepath)
	check(err)
	defer f.Close()
	sc := bufio.NewScanner(f)
	result := 0
	unique := 14
	start := ""
	for sc.Scan() {
		message := sc.Text()
		fmt.Println(message)
		for i, c := range message {
			start = string(c)
			if i > len(message)-unique {
				break
			} else {
				for f := 0; f < unique; f++ {
					start = start + string(message[f])
				}
				fmt.Println(start)
				for g := 0; g < len(start); g++ {
					if checkDups(start, string(start[g])) {
						continue
					} else {
						result = i
						break
					}
				}
			}
		}
		fmt.Println(start, result)
	}
	return result
}

func checkDups(s string, c string) bool {
	count := strings.Count(s, c)
	return count > 1
}
