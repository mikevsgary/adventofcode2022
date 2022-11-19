package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
)

func main() {
	fmt.Println("Advent of Code 2022: Day 3")
	// Open the diagnostic report file
	fileName := path.Join("day-3", "diagnostic_report")
	input, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer input.Close()

	// Create a [][]rune and populate it with the diagnostic report
	diagnosticReport := make([][]rune, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		diagnosticReport = append(diagnosticReport, line)
	}

	// Find the most common and least common rune in each column
	lineLength := len(diagnosticReport[0])
	mostCommon := make([]rune, lineLength)
	leastCommon := make([]rune, lineLength)
	for i := 0; i < lineLength; i++ {
		x, y := 0, 0
		for _, line := range diagnosticReport {
			if (line[i] == '1') {
				y++
			} else {
				x++
			}
		}
		if x < y {
			mostCommon[i] = '1'
			leastCommon[i] = '0'
		} else {
			mostCommon[i] = '0'
			leastCommon[i] = '1'
		}
	}

	// Multiply the gamma (mcr) by the epsilon (lcr) to find the consumption
	gamma, err := strconv.ParseInt(string(mostCommon), 2, 64)
	if err != nil {
		panic(err)
	}
	epsilon, err := strconv.ParseInt(string(leastCommon), 2, 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Gamma: %v Epsilon: %v Consumption: %v\n", string(mostCommon), epsilon, gamma*epsilon)

	// TODO - calculate the life support rating
	// TODO - calculate the oxygenRating
	lineLength = len(diagnosticReport[0])
	for i := 0; i < lineLength; i++ {

	}
	oxygenRating := getRating(diagnosticReport, '1')
	scrubberRating := getRating(diagnosticReport, '0')
	o, err := strconv.ParseInt(string(oxygenRating), 2, 64)
	if err != nil {
		panic(err)
	}
	c, err := strconv.ParseInt(string(scrubberRating), 2, 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Oxygen rating: %v Scrubber rating: %v LifeSupport: %v\n", string(oxygenRating), string(scrubberRating), o*c)
}

func getRating(diagnosticReport [][]rune, bitCriteria rune) []rune {
	lineLength := len(diagnosticReport[0])
	for i := 0; i < lineLength; i++ {
		filteredReport := make([][]rune, 0)
		for _, line := range diagnosticReport {
			if line[i] == bitfilter(diagnosticReport, i, bitCriteria) {
				filteredReport = append(filteredReport, line)
			}
		}
		if len(filteredReport) == 1 {
			return filteredReport[0]
		}
		diagnosticReport = filteredReport
	}
	e := fmt.Errorf("more then one line left in report: %v", diagnosticReport)
	panic(e)
}

func bitfilter(diagnosticReport [][]rune, i int, bitCriteria rune) rune {
	x, y := 0, 0
	for _, line := range diagnosticReport {
		if line[i] == '1' {
			y++
		} else {
			x++
		}
	}
	if (x > y && bitCriteria == '1') {
		return '0'
	} else if (y >= x && bitCriteria == '0') {
		return '0'
	} else {
		return '1'
	}
}
