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
		// fmt.Printf("%v %v\n", string(rps[0]), string(rps[2]))
		score += m[rps[2]]
		// fmt.Printf("throw %v new score %v\n", m[rps[2]], score)
		if m[rps[2]] == m[rps[0]] {
			score += 3
			// fmt.Printf("result %v new score %v\n", "3", score)
			continue
		}
		if rps[2] == 'X' && rps[0] == 'C' || rps[2] == 'Y' && rps[0] == 'A' {
			score += 6
			// fmt.Printf("result %v new score %v\n", "6", score)
			continue
		}
		if rps[2] == 'Z' && rps[0] == 'B' {
			score += 6
			// fmt.Printf("result %v new score %v\n", "6", score)
			continue
		}
	}

	fmt.Printf("%v\n", score)
}
