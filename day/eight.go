// Copy and paste this each day to get started
package day

import (
	"bufio"
	"fmt"
	"os"
)

func Eight(filepath string) {
	f, err := os.Open(filepath)
	check(err)
	defer f.Close()
	sc := bufio.NewScanner(f)
	trees := []string{}
	rowsOfTrees := 0
	columnsOfTrees := 0
	lineCount := 0

	for sc.Scan() {
		lineCount++
		rowsOfTrees += 2
		trees = append(trees, sc.Text())
	}

	columnsOfTrees = (len(trees[0]) * 2) - 4
	totalVisibleTrees := rowsOfTrees + columnsOfTrees

	for _, t := range trees {
		fmt.Println(t)
	}
	fmt.Println("total visible tress:", totalVisibleTrees)
}
