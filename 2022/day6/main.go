package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func dupCheck(value string, list []string) bool {
	var letterCheck = 0
	for _, v := range list {
		if v == value {
			letterCheck++
			if letterCheck >= 2 {
				fmt.Println("Letter:", v, "Found more than once in:", list)
				return true
			}
		}
	}
	return false
}

func main() {

	puzzle, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	signal := string(puzzle)

	var chk bool

	for i, _ := range signal {
		if i >= 3 {

			first := string(signal[i])
			second := string(signal[(i - 1)])
			third := string(signal[(i - 2)])
			fourth := string(signal[(i - 3)])

			tmpSlice := []string{fourth, third, second, first}
			chk = dupCheck(first, tmpSlice)
		} else {
			continue
		}

		if chk == false {
			return
		}
	}
}
