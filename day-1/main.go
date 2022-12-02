package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

func main() {
	fileName := "elf-inventory"
	filePath := filepath.Join("day-1", fileName)
	inventory, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer inventory.Close()

	foods := make([]string, 0)
	scanner := bufio.NewScanner(inventory)
	for scanner.Scan() {
		food := scanner.Text()
		foods = append(foods, food)
	}

	pockets := make([]int, 0)
	subTotal := 0
	for _, food := range foods {
		if food == "" {
			pockets = append(pockets, subTotal)
			subTotal = 0
			continue
		}
		calories, err := strconv.Atoi(food)
		if err != nil {
			panic(err)
		}

		subTotal = subTotal + calories
	}
	sort.Ints(pockets)
	i := len(pockets) - 1
	fmt.Printf("%v\n", pockets[i] + pockets[i-1] + pockets[i-2])
}
