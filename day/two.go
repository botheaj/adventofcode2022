package day

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Two(filepath string) {
	f, err := os.Open(filepath)
	check(err)
	defer f.Close()
	sc := bufio.NewScanner(f)
	totalScore := 0
	totalScore2 := 0
	for sc.Scan() {
		if sc.Text() != "" {
			g := strings.Split(sc.Text(), " ")
			totalScore += eval(g)
			totalScore2 += eval2(g)
		}
	}
	fmt.Println("My Total score:", totalScore)
	fmt.Println("My Total score part 2:", totalScore2)
}

func eval(g []string) int {
	gameMap := map[string]int{"A": 1, "B": 2, "C": 3, "X": 1, "Y": 2, "Z": 3}
	score := 0
	theirs := gameMap[g[0]]
	mine := gameMap[g[1]]
	switch {
	case theirs == mine:
		score = mine + 3
	case (theirs == 1 && mine == 2) || (theirs == 2 && mine == 3) || (theirs == 3 && mine == 1):
		score = mine + 6
	default:
		score = mine
	}
	return score
}

// x = lose, y = tie, z = win
func eval2(g []string) int {
	gameMap := map[string]int{"A": 1, "B": 2, "C": 3, "X": 0, "Y": 3, "Z": 6}
	winMap := map[string]string{"C": "A", "A": "B", "B": "C"}
	loseMap := map[string]string{"A": "C", "B": "A", "C": "B"}
	score := 0
	theirs := gameMap[g[0]]
	play := gameMap[g[1]]
	switch {
	case play == 6:
		score = gameMap[winMap[g[0]]] + play
	case play == 0:
		score = gameMap[loseMap[g[0]]]
	default:
		score = theirs + play
	}
	return score
}
