package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Advent of Code - Day 2")

	filename := "puzzle_input"
	filepath := filepath.Join("day-2", filename)
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	m := make(map[byte]int)
	m['A'] = 1
	m['B'] = 2
	m['C'] = 3
	m['X'] = 1
	m['Y'] = 2
	m['Z'] = 3

	score := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rps := scanner.Bytes()
		fmt.Printf("%v %v ", string(rps[0]), string(rps[2]))
		if rps[2] == 'Y' {
			score += 3
			score += m[rps[0]]
			 fmt.Printf("result %v new score %v\n", "3", score)
		} else if rps[2] == 'Z' {
			score += 6
			result := (m[rps[0]] + 1)
			if result > 3 {
				result = result %3
			}
			score += result
			 fmt.Printf("result %v new score %v\n", result, score)
		} else {
			result := (m[rps[0]] + 2)
			if result > 3 {
				result = result %3
			}
			score += result
			 fmt.Printf("result %v new score %v\n", result, score)
		}
	}

	fmt.Printf("%v\n", score)
}
