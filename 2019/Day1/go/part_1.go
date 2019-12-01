package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"

	color "github.com/logrusorgru/aurora"
)

func part1() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var totalFuel int

	// Get modules from input and convert bytes into strings in for range
	for _, module := range strings.Split(string(input), "\n") {
		if module == "" {
			continue
		}

		modflt64, _ := strconv.ParseFloat(module, 10)

		var mass float64 = math.Floor(modflt64/3) - 2

		massInt := int(mass)

		// fmt.Println("Module:", color.Magenta(module), "|", "Mass:", color.Magenta(mass))

		totalFuel = totalFuel + massInt
	}
	fmt.Println("Part 1 - Total fuel needed:", color.Green(totalFuel))
}
