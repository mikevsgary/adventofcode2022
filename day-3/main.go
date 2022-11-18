package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	fmt.Println("Hello World!")
	gamma, epsilon, consumption := getConsumption()
	fmt.Printf("gamma: %v epsilon: %v consumption: %v\n", gamma, epsilon, consumption)
	getOxygenGeneratorRating()
}

func getOxygenGeneratorRating() {
	filename := filepath.Join("day-3", "example")
	input, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	binaryReport := make([][]rune, 0)
	for scanner.Scan() {
		b := scanner.Text()
		binaryLine := make([]rune,0)
		for _, j := range b {
			binaryLine = append(binaryLine, j)
		}
		binaryReport = append(binaryReport, binaryLine)
	}

	
	oxygenRating,c02Rating := getRatings(binaryReport)
	a, err := strconv.ParseInt(string(oxygenRating), 2, 64) 
	if err != nil {
		panic(err)
	}
	b, err := strconv.ParseInt(string(c02Rating), 2, 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("result: %v\n", a *b)
}

func getRatings(binaryReport [][]rune) ([]rune, []rune) { 
	for i := 0; i < len(binaryReport[0]); i++ {
		key := mostCommon(binaryReport, i)
		filteredReport := make([][]rune, 0)
		for _, binaryLine := range binaryReport {
			if binaryLine[i] == key {
				filteredReport = append(filteredReport, binaryLine)
			}
		}
		if len(filteredReport) < 1 {
			break
		}
		binaryReport = filteredReport
	}
	return binaryReport[0], binaryReport[0]
}

func mostCommon(binaryReport [][]rune, i int) rune {
	x := 0
	for _, line := range binaryReport {
		x = x + int(line[i] - '0')
	}
	if x / len(binaryReport) > 0 {
		return '1'
	}
	return '0'
}
func getConsumption() (string, string, int64) {
	filename := filepath.Join("day-3", "example")
	input, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	binaryReport := make([]float64, 5)

	reportCount := 0
	for scanner.Scan() {
		reportCount++
		b := scanner.Text()
		for i, j := range b {
			f := int(j - '0')
			binaryReport[i] = binaryReport[i] + float64(f)
		}
	}

	gamma, epsilon := "", ""
	for _, x := range binaryReport {
		// fmt.Printf("%v : %v\n", x, reportCount)
		if x/float64(reportCount) >= 0.5 {
			gamma = gamma + "1"
			epsilon = epsilon + "0"
		} else {
			gamma = gamma + "0"
			epsilon = epsilon + "1"
		}
	}
	g, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		panic(err)
	}
	e, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		panic(err)
	}
	consumption := g*e
	return gamma, epsilon, consumption
}
