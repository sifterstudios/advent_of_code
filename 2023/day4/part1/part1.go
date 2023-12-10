package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./2023/input/day4_part1_input")
	if err != nil {
		println("Error reading file")
	}

	defer file.Close()
	scan := bufio.NewScanner(file)

	totalWorthOfPoints := 0

	for scan.Scan() {
		snipBeginning := strings.Split(scan.Text(), ":")
		snipMiddle := strings.Split(snipBeginning[1], "|")

		winningNumbers := splitAndConvert(snipMiddle[0])
		guessNumbers := splitAndConvert(snipMiddle[1])

		count := checkForHits(guessNumbers, winningNumbers)
		totalWorthOfPoints += count
	}
	println("Total worth of all cards: ", totalWorthOfPoints)
}

func checkForHits(guessNumbers, winningNumbers []int) int {
	worth := 0
	for _, v := range guessNumbers {
		if slices.Contains(winningNumbers, v) {
			if worth == 0 {
				worth = 1
			} else {
				worth *= 2
			}
		}
	}
	return worth
}

func splitAndConvert(s string) []int {
	var numbers []int
	for _, v := range strings.Split(s, " ") {
		v, _ := strconv.Atoi(v)
		if v != 0 {
			numbers = append(numbers, v)
		}
	}
	return numbers
}
