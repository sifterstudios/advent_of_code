package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type handBet struct {
	allCards []int
	bet      int
	rank     int
}

var (
	fiveOfAKindIndexes  []int
	fourOfAKindIndexes  []int
	fullHouseIndexes    []int
	threeOfAKindIndexes []int
	twoPairIndexes      []int
	onePairIndexes      []int
	highCardIndexes     []int
	totalSumFinito      int
	allHandBets         []handBet
	finalRanks          []handBet
)

func main() {
	file, err := os.Open("./2023/input/day7_part2_input")
	if err != nil {
		println("Could not open file.")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			println("Could not close file.")
		}
	}(file)

	scanner := bufio.NewScanner(file)
	getInput(scanner)

	for i := 0; i < len(allHandBets); i++ {
		determineHandRank(allHandBets[i], i)
	}

	println("Five of a kind:", len(fiveOfAKindIndexes))
	println("Four of a kind:", len(fourOfAKindIndexes))
	println("Full house:", len(fullHouseIndexes))
	println("Three of a kind:", len(threeOfAKindIndexes))
	println("Two pair:", len(twoPairIndexes))
	println("One pair:", len(onePairIndexes))
	println("High card:", len(highCardIndexes))

	determineStrengthOfHandsAndSort(allHandBets, fiveOfAKindIndexes, false)
	determineStrengthOfHandsAndSort(allHandBets, fourOfAKindIndexes, false)
	determineStrengthOfHandsAndSort(allHandBets, fullHouseIndexes, true)
	determineStrengthOfHandsAndSort(allHandBets, threeOfAKindIndexes, false)
	determineStrengthOfHandsAndSort(allHandBets, twoPairIndexes, false)
	determineStrengthOfHandsAndSort(allHandBets, onePairIndexes, false)
	determineStrengthOfHandsAndSort(allHandBets, highCardIndexes, false)

	for i := 0; i < len(finalRanks); i++ {
		rank := len(finalRanks) - i
		rankTimesBet := rank * finalRanks[i].bet
		totalSumFinito += rankTimesBet
	}

	println("Total sum:", totalSumFinito)
}

func determineStrengthOfHandsAndSort(allHandBets []handBet, indexes []int, shouldPrint bool) {
	bubbleSortHandsAndAddToFinalRanks(allHandBets, indexes, shouldPrint)
}

func removeElement(localIndexes []int, remove int) []int {
	var newIndex []int
	for i := 0; i < len(localIndexes); i++ {
		if localIndexes[i] == remove {
			continue
		}
		newIndex = append(newIndex, localIndexes[i])
	}
	return newIndex
}

func bubbleSortHandsAndAddToFinalRanks(allHandBets []handBet, indexes []int, shouldPrint bool) {
	var handsToBeSorted []handBet
	var tempHandBet handBet

	for i := 0; i < len(indexes); i++ {
		handsToBeSorted = append(handsToBeSorted, allHandBets[indexes[i]])
	}

	for i := 0; i < len(handsToBeSorted); i++ {
		for j := 0; j < len(handsToBeSorted)-1; j++ {
			for card := 0; card < 5; card++ {
				thisCard := handsToBeSorted[j].allCards[card]
				nextCard := handsToBeSorted[j+1].allCards[card]
				if thisCard == nextCard {
					continue
				} else if thisCard > nextCard {
					// fmt.Printf("Swapping %v and %v\n", handsToBeSorted[j].allCards, handsToBeSorted[j+1].allCards)
					tempHandBet = handsToBeSorted[j]
					handsToBeSorted[j] = handsToBeSorted[j+1]
					handsToBeSorted[j+1] = tempHandBet
					break
				} else {
					break
				}

			}
		}
	}

	for i := len(handsToBeSorted) - 1; i > -1; i-- {
		finalRanks = append(finalRanks, handsToBeSorted[i])
		if shouldPrint {
			fmt.Printf("hand: %v\n", handsToBeSorted[i].allCards)
		}
	}
}

func determineHandRank(handBet handBet, idx int) {
	distinctNumbersMap := make(map[int]int)
	distinctNumbers := -1
	for i := 0; i < len(handBet.allCards); i++ {
		distinctNumbersMap[handBet.allCards[i]] = handBet.allCards[i] + 1
	}

	distinctNumbers = len(distinctNumbersMap)

	if _, ok := distinctNumbersMap[1]; ok {
		distinctNumbers--
	}

	if distinctNumbers == 5 {
		highCardIndexes = append(highCardIndexes, idx)
	} else if distinctNumbers == 4 {
		onePairIndexes = append(onePairIndexes, idx)
	} else if distinctNumbers == 3 {
		three := determineThreeOfAKind(handBet.allCards)
		if three {
			threeOfAKindIndexes = append(threeOfAKindIndexes, idx)
		} else {
			twoPairIndexes = append(twoPairIndexes, idx)
		}
	} else if distinctNumbers == 2 {
		four := determineFourOfAKind(handBet.allCards)
		if four {
			fourOfAKindIndexes = append(fourOfAKindIndexes, idx)
		} else {
			fullHouseIndexes = append(fullHouseIndexes, idx)
		}
	} else {
		fiveOfAKindIndexes = append(fiveOfAKindIndexes, idx)
	}
}

func determineThreeOfAKind(entry []int) bool {
	count := make(map[int]int)
	maxCount := -1
	for i := 0; i < len(entry); i++ {
		if _, ok := count[entry[i]]; ok {
			count[entry[i]]++
		} else {
			count[entry[i]] = 1
		}
		if count[entry[i]] > maxCount && entry[i] != 1 {
			maxCount = count[entry[i]]
		}

	}
	if maxCount == 3 ||
		(count[1] == 1 && maxCount == 2) ||
		(count[1] == 2 && maxCount == 1) {
		return true
	}
	return false
}

func determineFourOfAKind(entry []int) bool {
	count := make(map[int]int)
	maxCount := -1
	for i := 0; i < len(entry); i++ {
		if _, ok := count[entry[i]]; ok {
			count[entry[i]]++
		} else {
			count[entry[i]] = 1
		}
		if count[entry[i]] > maxCount && entry[i] != 1 {
			maxCount = count[entry[i]]
		}
	}

	if (maxCount == 3 && count[1] == 1) ||
		(maxCount == 2 && count[1] == 2) ||
		(maxCount == 1 && count[1] == 3) ||
		maxCount == 4 {
		return true
	}
	return false
}

func getInput(scanner *bufio.Scanner) {
	for scanner.Scan() {
		currentLine := scanner.Text()
		handAndBet := strings.Split(currentLine, " ")
		handNumberSlice := splitHandStringToIntSplice(handAndBet)
		betInt, _ := strconv.Atoi(handAndBet[1])
		allHandBets = append(allHandBets, handBet{allCards: handNumberSlice, bet: betInt, rank: -1})
		// println("Hand:", handNumberSlice[0], handNumberSlice[1], handNumberSlice[2], handNumberSlice[3], handNumberSlice[4], "Bet:", betInt)
	}
}

func splitHandStringToIntSplice(handAndBet []string) (slicedStrings []int) {
	for i := 0; i < len(handAndBet[0]); i++ {
		v := string(handAndBet[0][i])
		if isNumber(v) {
			intValue, _ := strconv.Atoi(v)
			slicedStrings = append(slicedStrings, intValue)
			continue
		}
		if v == "T" {
			v = "10"
		}
		if v == "J" {
			v = "1"
		}
		if v == "Q" {
			v = "12"
		}
		if v == "K" {
			v = "13"
		}
		if v == "A" {
			v = "14"
		}
		intValue, _ := strconv.Atoi(v)
		slicedStrings = append(slicedStrings, intValue)
	}
	return slicedStrings
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
