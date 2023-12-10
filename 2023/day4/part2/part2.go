package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	winningNumbers  []int
	guessNumbers    []int
	extraIterations int
}

var allCards []Card

func main() {
	file, err := os.Open("./2023/input/day4_part2_input")
	if err != nil {
		println("Error reading file")
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			println("Had trouble closing the file")
		}
	}(file)
	scan := bufio.NewScanner(file)

	totalCards := 0

	for scan.Scan() {
		currentCard := Card{}
		snipBeginning := strings.Split(scan.Text(), ":")
		snipMiddle := strings.Split(snipBeginning[1], "|")

		winningNumbers := splitAndConvert(snipMiddle[0])
		guessNumbers := splitAndConvert(snipMiddle[1])

		currentCard.winningNumbers = winningNumbers
		currentCard.guessNumbers = guessNumbers

		allCards = append(allCards, currentCard)
	}
	println("Length of allCards: ", len(allCards))

	for i := 0; i < len(allCards); i++ {
		currentCard := allCards[i]
		if currentCard.extraIterations > 0 {
			for j := -1; j < currentCard.extraIterations; j++ {
				checkForHits(currentCard.guessNumbers, currentCard.winningNumbers, i)
				totalCards++
			}
		} else {
			checkForHits(currentCard.guessNumbers, currentCard.winningNumbers, i)
			totalCards++
		}
	}

	println("Total worth of all cards: ", totalCards)
}

func checkForHits(guessNumbers, winningNumbers []int, idx int) {
	wins := 0
	for _, v := range guessNumbers {
		if slices.Contains(winningNumbers, v) {
			wins += 1
		}
	}
	for cardCopy := idx + 1; cardCopy < idx+wins+1; cardCopy++ {
		allCards[cardCopy].extraIterations += 1
	}
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
