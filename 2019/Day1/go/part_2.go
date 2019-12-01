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

// I found it easier to just work with float64 rather than doing a bunch of casting back and forth, one final cast back to int might be best.
func part2() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// This value will count up to the total fuel needed plus the additional fuel needed.
	moduleTotalFuel := int32(0)

	for _, module := range strings.Split(string(input), "\n") {
		if module == "" {
			continue
		}

		// Define mass here so we can count down to 0 or negative
		mass, _ := strconv.ParseFloat(module, 64)

		// Loop through until the counter is 0 or negative.
		for mass >= 0 {

			mass = math.Floor(mass/3) - 2
			// Might be a better way to do this, but need to eject the loop here so it doesn't subtract extra fuel.
			if mass <= 0 {
				continue
			}

			massInt32 := int32(mass)
			moduleTotalFuel = moduleTotalFuel + massInt32
		}
	}
	fmt.Println("Part 2 - Total fuel needed:", color.Green(moduleTotalFuel))
}
