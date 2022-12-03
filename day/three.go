package day

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Three(filepath string) {

	// Build a map (alphanum) of letters to values eg. a=1, b=2, A=27 Z=52
	alphanum := make(map[string]int)

	var ch byte
	var i = 1
	for ch = 'a'; ch <= 'z'; ch++ {
		alphanum[string(ch)] = i
		i++
	}
	for ch = 'A'; ch <= 'Z'; ch++ {
		alphanum[string(ch)] = i
		i++
	}

	//Open file and read each line
	f, err := os.Open(filepath)
	check(err)
	defer f.Close()
	sc := bufio.NewScanner(f)
	priority1 := 0
	priority2 := 0
	grouping := 1
	group := make([][]string, 3)
	group = nil

	for sc.Scan() {
		if sc.Text() != "" {
			//fmt.Println(sc.Text())
			rucksack := strings.Split(sc.Text(), "")
			group = append(group, rucksack)
			var rucksacka []string
			var rucksackb []string

			for i := 0; i < len(rucksack); i++ {
				if i < len(rucksack)/2 {
					rucksacka = append(rucksacka, rucksack[i])
				} else {
					rucksackb = append(rucksackb, rucksack[i])
				}
			}
			priority1 += alphanum[compareSlices(rucksacka, rucksackb)]
		}

		if grouping == 3 {
			priority2 += alphanum[compare3Slices(group[0], group[1], group[2])]
			group = nil
			grouping = 1
		} else {
			grouping++
		}
	}
	fmt.Println("Part 1:", priority1)
	fmt.Println("Part 2:", priority2)
}

func compareSlices(a []string, b []string) string {
	match := ""
	for _, itema := range a {
		for _, itemb := range b {
			if itema == itemb {
				match = itema
			}
		}
	}
	return match
}

func compare3Slices(a []string, b []string, c []string) string {
	match := ""
	for _, itema := range a {
		for _, itemb := range b {
			if itema == itemb {
				for _, itemc := range c {
					if itema == itemc {
						match = itemc
					}
				}
			}
		}
	}
	return match
}
