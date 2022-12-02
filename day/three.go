package day

import (
	"bufio"
	"fmt"
	"os"
)

func Three(filepath string) {
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
