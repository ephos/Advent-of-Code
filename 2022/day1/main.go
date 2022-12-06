package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

type elf struct {
	elf      int
	calories int
}

func main() {

	puzzle, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	elves := strings.Split(string(puzzle), "\n\n")

	elfSlice := []elf{}

	for ei, e := range elves {

		cals := strings.Split(e, "\n")

		totalCals := 0
		for _, c := range cals {
			if c != "" {
				ci, err := strconv.Atoi(c)
				if err != nil {
					log.Fatal(err)
				}
				totalCals += ci
			}
		}

		// Create elf with his little fucking dumb elf number and how much that fat fuck is hoarding
		currentElf := elf{
			elf:      ei,
			calories: totalCals,
		}
		elfSlice = append(elfSlice, currentElf)
	}

	sort.Slice(elfSlice, func(i, j int) bool {
		return elfSlice[i].calories > elfSlice[j].calories

	})

	fmt.Println("Part 1:", elfSlice[0].calories)
	fmt.Println("Part 2:", elfSlice[0].calories+elfSlice[1].calories+elfSlice[2].calories)
}
