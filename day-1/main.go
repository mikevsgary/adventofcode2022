package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fileName := filepath.Join("resources", "depths.txt")
	depths := make([]int, 0)
	dat, err := os.ReadFile(fileName)
	check(err)
	input := strings.Fields(string(dat))
	sum := 0
	for j, i := range input {
		x, err := strconv.Atoi(i)
		check(err)
		sum = sum + x
		if j > 2 {
			y, err := strconv.Atoi(input[j-3])
			check(err)
			sum = sum - y
		}
		if j > 1 {
			depths = append(depths, sum)
		}
	}
	CountIncreases(depths)
}

func CountIncreases(depths []int) int {
	count := 0
	prevDepth := depths[0]
	for _, depth := range depths {
		change := "No change"
		if depth > prevDepth {
			count++
			change = "Increased"
		} else if depth < prevDepth {
			change = "Decreased"
		}
		fmt.Printf("%v Prev Depth %v %v Count: %v\n", depth, prevDepth, change, count)
		prevDepth = depth
	}
	return count
}
