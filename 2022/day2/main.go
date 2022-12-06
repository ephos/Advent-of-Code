package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var rpsDict = func() map[string]string {
	return map[string]string{
		"A": "Rock",
		"B": "Paper",
		"C": "Scissor",
		"X": "Rock",
		"Y": "Paper",
		"Z": "Scissor",
	}
}

func rpsMove(key string) string {
	return rpsDict()[key]
}

func RockPaperScissor(opponent string, player string) int {

	var scoreMap = map[string]int{
		"Rock":    1,
		"Paper":   2,
		"Scissor": 3,
	}

	if player == opponent {
		return (3 + scoreMap[player])
	} else if player == "Rock" {
		if opponent == "Scissor" {
			return (6 + scoreMap[player])
		}
	} else if player == "Scissor" {
		if opponent == "Paper" {
			return (6 + scoreMap[player])
		}
	} else if player == "Paper" {
		if opponent == "Rock" {
			return (6 + scoreMap[player])
		}
	}
	return (0 + scoreMap[player])
}

var rpsDictOutcome = func() map[string]string {
	return map[string]string{
		"A": "Rock",
		"B": "Paper",
		"C": "Scissor",
		"X": "Lose",
		"Y": "Draw",
		"Z": "Win",
	}
}

func rpsMatch(key string) string {
	return rpsDictOutcome()[key]
}

func RockPaperScissorRigged(opponent string, outcome string) int {

	var scoreMap = map[string]int{
		"Rock":    1,
		"Paper":   2,
		"Scissor": 3,
	}

	if outcome == "Lose" {
		if opponent == "Rock" {
			return (scoreMap["Scissor"])
		} else if opponent == "Paper" {
			return (scoreMap["Rock"])
		} else {
			return (scoreMap["Paper"])
		}
	} else if outcome == "Win" {
		if opponent == "Rock" {
			return (6 + scoreMap["Paper"])
		} else if opponent == "Paper" {
			return (6 + scoreMap["Scissor"])
		} else {
			return (6 + scoreMap["Rock"])
		}
	} else {
		if opponent == "Rock" {
			return (3 + scoreMap["Rock"])
		} else if opponent == "Paper" {
			return (3 + scoreMap["Paper"])
		} else {
			return (3 + scoreMap["Scissor"])
		}
	}
}

func sum(array []int) int {
	ret := 0
	for _, v := range array {
		ret += v
	}
	return ret
}

func main() {

	puzzle, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Function helps split strings on \n
	sf := func(c rune) bool {
		return c == '\n'
	}
	gameRounds := strings.FieldsFunc(string(puzzle), sf)

	scoreSlice := []int{}
	scoreSlice2 := []int{}

	for _, g := range gameRounds {
		r := strings.Split(g, " ")

		oppt := rpsMove(r[0])
		plyr := rpsMove(r[1])

		scoreSlice = append(scoreSlice, RockPaperScissor(oppt, plyr))

		//----------------Part 2

		opptShape := rpsMatch(r[0])
		gameOutcome := rpsMatch(r[1])

		scoreSlice2 = append(scoreSlice2, RockPaperScissorRigged(opptShape, gameOutcome))
	}

	fmt.Println("Part 1: ", sum(scoreSlice))
	fmt.Println("Part 2: ", sum(scoreSlice2))
}
