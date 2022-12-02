package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"sort"
)

func main() {
	var elfWeights []int
	elfCal := 0

	f, err := os.Open("input.txt")
    check(err)
    defer f.Close()
	sc := bufio.NewScanner(f)


	for sc.Scan() {
		if sc.Text() != "" {
			elfCalLine, err := strconv.Atoi(sc.Text())
			check(err)
			elfCal += elfCalLine
		} else {
			elfWeights = append(elfWeights, elfCal)
			elfCal = 0
		}
    }
	
	sort.Slice(elfWeights, func(i, j int) bool {
		return elfWeights[i] < elfWeights[j]
	})
	fmt.Println("part 1:", elfWeights[len(elfWeights)-1])

	fmt.Println("part 2:", (elfWeights[len(elfWeights)-1] + elfWeights[len(elfWeights)-2] + elfWeights[len(elfWeights)-3]))
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}