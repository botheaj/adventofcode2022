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

	rowsOfTrees := 0
	columnsOfTrees := 0
	lineCount := 0
	rowA := []string{}
	rowB := []string{}
	rowC := []string{}
	for sc.Scan() {
		lineCount++
		columnsOfTrees = (len(sc.Text()) * 2) - 4
	}

	for sc.Scan() {
		rowsOfTrees += 2

	}
	totalVisibleTrees := rowsOfTrees + columnsOfTrees
	fmt.Println("total visible tress:", totalVisibleTrees)
}
