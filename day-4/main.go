package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Advent of Code: Day 4")

	callsFileName := "calls"
	cardsFileName := "cards"
	callsFilePath := path.Join("day-4", callsFileName)
	cardsFilePath := path.Join("day-4", cardsFileName)

	cardsFile, err := os.Open(cardsFilePath)
	if err != nil {
		panic(err)
	}
	defer cardsFile.Close()

	callsFile, err := os.ReadFile(callsFilePath)
	if err != nil {
		panic(err)
	}

	calls := strings.Fields(string(callsFile))
	fmt.Printf("calls: %v\n", calls)

	cards := make([][][]int, 0)
	card := make([][]int, 0)

	scanner := bufio.NewScanner(cardsFile)

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, 0)

		if len(line) < 1 {
			cards = append(cards, card)
			card = make([][]int, 0)
			continue
		}

		f := strings.Fields(line)
		for _, c := range f {
			n, err := strconv.Atoi(c)
			if err != nil {
				panic(err)
			}
			row = append(row, n)
		}
		card = append(card, []int(row))
	}
		cards = append(cards, card)

	for i, card := range cards {
		fmt.Printf("Card: %v\n", i)
		for _, row := range card {
			fmt.Printf("  row: %v\n", row)
		}
	}

	scoreCards := cards
	for i, scoreCard := range scoreCards {
		fmt.Printf("scoreCard: %v\n", i)
		for _, row := range scoreCard {
			fmt.Printf("  row: %v\n", row)
		}
	}

}
