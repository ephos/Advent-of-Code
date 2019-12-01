package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"

	color "github.com/logrusorgru/aurora"
)

func main() {
	// Read in contents of input, this will return an array of bytes
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Convert byte array into big string
	inputString := string(input)

	// Read in all the strings from the file and split them on newlines (/n)
	presents := strings.Split(string(inputString), "\n")

	// Define var for total square feet.
	var totalSqFt int
	var totalFtRibbon int

	for index, present := range presents {
		if present == "" {
			continue
		}

		// Read in the present dimensions
		presentDimensions := strings.Split(present, "x")
		length, _ := strconv.ParseInt(presentDimensions[0], 10, 0)
		width, _ := strconv.ParseInt(presentDimensions[1], 10, 0)
		height, _ := strconv.ParseInt(presentDimensions[2], 10, 0)

		// Get smallest area for additional paper needed.
		presentPaperSlice := []int{int(length * width), int(width * height), int(height * length)}
		sort.Ints(presentPaperSlice)
		smallestSide := int(presentPaperSlice[0])

		presentRibbonSlice := []int{int(length), int(width), int(height)}
		sort.Ints(presentRibbonSlice)
		// Ribbon Side
		ftRibbon := int(presentRibbonSlice[0] + presentRibbonSlice[0] + presentRibbonSlice[1] + presentRibbonSlice[1])
		// Ribbon Bow
		ftBow := int(length * width * height)
		// Total ribbon
		presentRibbonFt := int(ftRibbon + ftBow)

		// fmt.Println("Present No.:", color.Cyan(index), ",", "length:", color.Cyan(length), ",", "width:", color.Cyan(width), ",", "height:", color.Cyan(height), ",", "smallest side:", color.Cyan(smallestSide))

		presentSqFt := int(2*(length*width)+2*(width*height)+2*(height*length)) + smallestSide

		fmt.Println("Present No.:", color.Cyan(index), "Total Sq Ft:", color.Magenta(presentSqFt))
		fmt.Println("Present No.:", color.Cyan(index), "Total Ribion Ft:", color.Magenta(presentRibbonFt))

		totalSqFt = totalSqFt + presentSqFt
		totalFtRibbon = totalFtRibbon + presentRibbonFt
		// fmt.Println("Total Sq Ft.:", color.Bold(color.Green(totalSqFt)), "")
	}

	fmt.Println("Elves need:", color.Green(totalSqFt), "sq feet of wrapping paper.")
	fmt.Println("Elves need:", color.Green(totalFtRibbon), "feet of ribbon.")
}
