package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	color "github.com/logrusorgru/aurora"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Creates a slice of bytes
	// steps := []byte(input)
	// Create a var in for floor count
	var floor int
	// Create a slice of ints with no initial values
	basementVisits := []int{}

	// This is basically a "foreach" in Go.  Convert the byte array into strings on the for range line.
	for index, step := range strings.Split(string(input), "") {
		// Convert bytes back to string
		fmt.Println("Current Step:", color.Cyan(string(step)), "Running for floor move:", color.Cyan(index))

		if step == "(" {
			fmt.Println("Up")
			floor++
		} else if step == ")" {
			fmt.Println("Down")
			floor--
		}

		if floor == -1 {
			basementVisits = append(basementVisits, index)
			fmt.Println("Santa has entered the basement at floor: ", color.Green(index))
		}
	}
	// Add one since index starts at 0 but the puzzle starts at 1
	firstBasementVisit := basementVisits[0] + 1
	fmt.Println("Current floor:", color.Magenta(floor))
	fmt.Println("First basement visit:", color.Magenta(firstBasementVisit))
}
