package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("Advent of Code - Day 3 taking a break at 47:51 + 15:15")

	// TODO - add a flag for "example" or "puzzle"
	// TODO - check against an output instead of printing to console
	filename := "puzzle_input"
	filepath := filepath.Join("day-3", filename)
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	groups := make([][]byte, 0)
	groupsum := 0
	for scanner.Scan() {
		rucksack := scanner.Bytes()
		groups = append(groups, rucksack)
		// fmt.Printf("%v ", string(rucksack))
		c1, c2 := rucksack[:len(rucksack)/2], rucksack[len(rucksack)/2:]
		// fmt.Printf("%v %v\n", string(c1), string(c2))
		for _, item := range c1 {
			if strings.ContainsRune(string(c2), rune(item)) {
				priority := 0
				if item > 96 {
					priority = int(item) - 96
				} else {
					priority = int(item) - 38
				}
				// fmt.Printf("%v %v\n", string(item), priority)
				sum += int(priority)
				break
			}
		}
		if len(groups) == 3 {
			for _, item := range groups[0] {
				if strings.ContainsRune(string(groups[1]), rune(item)) && strings.ContainsRune(string(groups[2]), rune(item)) {
					priority := 0
					if item > 96 {
						priority = int(item) - 96
					} else {
						priority = int(item) - 38
					}
					// fmt.Printf("%v %v\n", string(item), priority)
					groupsum += int(priority)
					groups = make([][]byte,0)
					break
				}
			}
		}
	}

	fmt.Printf("sum: %v groups: %v\n", sum, groupsum)
}
