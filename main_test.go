package main

import (
	"fmt"
	"testing"

	"github.com/botheaj/adventofcode2022/day"
)

func TestMain(t *testing.T) {
	if 2+2 != 4 {
		t.Error("placeholder for day four")
	}
}

func TestDay5(t *testing.T) {
	var expected []string
	var results []string
	expected = append(expected, "C")
	expected = append(expected, "M")
	expected = append(expected, "Z")
	results = day.Five("inputs/5_test.txt")
	fmt.Println("5 expected", expected)
	if !compare(expected, results) {
		t.Error("Does not match expected output")
	}
}

func compare(a []string, b []string) bool {
	match := true
	for i, v := range a {
		if v != b[i] {
			match = false
		}
	}
	return match
}

func TestDay6(t *testing.T) {
	expected := 5
	results = day.Six("inputs/5_test.txt")
	fmt.Println("6 expected", expected)
	if !compare(expected, results) {
		t.Error("Does not match expected output")
	}
}
