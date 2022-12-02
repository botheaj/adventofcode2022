// Copy and paste this each day to get started
package day

import (
	"bufio"
	"fmt"
	"os"
)

func Zero(filepath string) {
	f, err := os.Open(filepath)
	check(err)
	defer f.Close()
	sc := bufio.NewScanner(f)

	for sc.Scan() {
		if sc.Text() != "" {
			fmt.Println(sc.Text)
			check(err)
		}
	}
}
