package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello World!")

	filename := filepath.Join("day-2", "puzzle_input")
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	commandList := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		commandList = append(commandList, scanner.Text())
	}

	position, aim, depth := 0, 0, 0

	for _, command := range commandList {
		fields := strings.Fields(command)
		x, err := strconv.Atoi(fields[1])
		check(err)
		switch fields[0] {
		case "forward":
			position += x
			depth = depth + (aim * x)
		case "down":
			aim += x
		case "up":
			aim -= x
		}
	}

	fmt.Printf("Position: %v Depth: %v Answer: %v", position, depth, position*depth)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
