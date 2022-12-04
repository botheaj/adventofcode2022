// Copy and paste this each day to get started
package day

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Four(filepath string) {
	// as always, read the file
	f, err := os.Open(filepath)
	check(err)
	defer f.Close()
	sc := bufio.NewScanner(f)
	part1Solution := 0
	part2Solution := 0
	for sc.Scan() {
		if sc.Text() != "" {
			part1Solution += part1(sc.Text())
			part2Solution += part2(sc.Text())
		}
	}
	fmt.Println("Part 1:", part1Solution)
	fmt.Println("Part 2:", part2Solution)
}

func part1(s string) int {
	// slice line into values of ranges
	ranges := strings.Split(s, ",")
	// slice ranges into integer values for each side
	range1 := strings.Split(ranges[0], "-")
	range2 := strings.Split(ranges[1], "-")

	// convert to integer and compare
	range1Lower, _ := strconv.Atoi(range1[0])
	range1Upper, _ := strconv.Atoi(range1[1])
	range2Lower, _ := strconv.Atoi(range2[0])
	range2Upper, _ := strconv.Atoi(range2[1])
	if range2Lower >= range1Lower && range2Upper <= range1Upper {
		return 1
	} else if range1Lower >= range2Lower && range1Upper <= range2Upper {
		return 1
	}
	return 0
}

func part2(s string) int {
	// slice line into values of ranges
	ranges := strings.Split(s, ",")
	// slice ranges into integer values for each side
	range1 := strings.Split(ranges[0], "-")
	range2 := strings.Split(ranges[1], "-")

	// 	2-4,6-8
	// 2-3,4-5
	// 5-7,7-9 *
	// 2-8,3-7 *
	// 6-6,4-6 *
	// 2-6,4-8 *

	// convert to integer and compare
	range1Lower, _ := strconv.Atoi(range1[0])
	range1Upper, _ := strconv.Atoi(range1[1])
	range2Lower, _ := strconv.Atoi(range2[0])
	range2Upper, _ := strconv.Atoi(range2[1])
	var rangeSlice1 []int
	var rangeSlice2 []int

	for i := range1Lower; i <= range1Upper; i++ {
		rangeSlice1 = append(rangeSlice1, i)
	}
	for i := range2Lower; i <= range2Upper; i++ {
		rangeSlice2 = append(rangeSlice2, i)
	}
	return compareIntSlices(rangeSlice1, rangeSlice2)
}

func compareIntSlices(a []int, b []int) int {
	match := 0
	for _, itema := range a {
		for _, itemb := range b {
			if itema == itemb {
				match = 1
			}
		}
	}
	return match
}
