package main

import (
	"bufio"
	"os"
	"strconv"
)

type numberEntry struct {
	line       int
	start      int
	end        int
	fullNumber int
}

type starEntry struct {
	line int
	idx  int
}

func main() {
	file, err := os.Open("./2023/input/day3_part2_input")
	if err != nil {
		println("Error opening file")
	}

	scanner := bufio.NewScanner(file)
	var allLines []string
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			println("Trouble closing the file!")
		}
	}(file)

	for scanner.Scan() {
		allLines = append(allLines, scanner.Text())
	}

	sumGears := 0

	getSumOfGears(allLines, &sumGears)
	println("Sum of gears: ", sumGears)
}

func getSumOfGears(allLines []string, sumGears *int) {
	var numberEntries []numberEntry
	var starEntries []starEntry

	for i := 0; i < len(allLines); i++ {
		currentLine := allLines[i]
		for j := 0; j < len(currentLine); j++ {
			currentChar := string(allLines[i][j])
			if currentChar == "*" {
				starEntries = append(starEntries, starEntry{i, j})
				continue
			}
			numberEntries = getNumberEntryForChar(allLines, currentChar, i, &j, currentLine, numberEntries)
			if j == len(currentLine)-1 {
				break
			}
		}
	}

	println("NumberEntries: ", len(numberEntries))
	for star := 0; star < len(starEntries); star++ {
		numbersFound := checkForNumberInAllDirections(starEntries[star], numberEntries)
		if len(numbersFound) == 2 {
			product := numbersFound[0] * numbersFound[1]
			*sumGears += product
		} else if len(numbersFound) > 2 {
			println("More than two numbers found for star #", star)
		}
	}
	println("Found ", len(starEntries), " stars")
}

func getNumberEntryForChar(allLines []string, currentChar string, i int, j *int, currentLine string, numberEntries []numberEntry) []numberEntry {
	if isNumber(currentChar) {
		entry := numberEntry{i, *j, -1, -1}
		for *j < len(currentLine) && isNumber(string(currentLine[*j])) {
			if *j < len(currentLine) {
				*j++
			} else {
				break
			}
		}
		*j--
		// Was no longer a number or end of line
		entry.end = *j
		entry.fullNumber, _ = strconv.Atoi(getFullNumber(allLines[i], entry.start))

		// Add to slice
		numberEntries = append(numberEntries, entry)
	}
	return numberEntries
}

func checkForNumberInAllDirections(star starEntry, numberEntries []numberEntry) (numbersFound []int) {
	idx, line := star.idx, star.line

	// ABOVE
	above1, above2 := numberAboveOrBelow(numberEntries, line-1, idx)
	if above1 != -1 {
		numbersFound = append(numbersFound, above1)
	}
	if above2 != -1 && above1 != above2 {
		numbersFound = append(numbersFound, above2)
	}

	below1, below2 := numberAboveOrBelow(numberEntries, line+1, idx)
	if below1 != -1 {
		numbersFound = append(numbersFound, below1)
	}
	if below2 != -1 && below1 != below2 {
		numbersFound = append(numbersFound, below2)
	}

	left, right := numberAdjacent(numberEntries, line, idx)
	if left != -1 {
		numbersFound = append(numbersFound, left)
	}
	if right != -1 {
		numbersFound = append(numbersFound, right)
	}

	if len(numbersFound) > 2 {
		println("Something is terribly wrong, found more than 2 numbers at location ", line, idx)
	}

	return numbersFound
}

func numberAboveOrBelow(numberEntries []numberEntry, line int, idx int) (number1, number2 int) {
	number1 = -1
	number2 = -1
	for number := 0; number < len(numberEntries); number++ {
		numberEntry := numberEntries[number]
		if numberEntry.line == line {
			for i := idx - 1; i <= idx+1; i++ {
				if i >= numberEntry.start && i <= numberEntry.end {
					if number1 != -1 {
						number2 = numberEntry.fullNumber
						continue
					} else {
						number1 = numberEntry.fullNumber
						continue
					}
				}
			}
		}
	}
	return number1, number2
}

func numberAdjacent(numberEntries []numberEntry, line int, idx int) (left, right int) {
	left = -1
	right = -1
	for number := 0; number < len(numberEntries); number++ {
		numberEntry := numberEntries[number]
		if numberEntry.line == line {
			if idx-1 == numberEntry.end {
				left = numberEntry.fullNumber
			}
			if idx+1 == numberEntry.start {
				right = numberEntry.fullNumber
			}
		}
	}
	return left, right
}

func getFullNumber(line string, idx int) (fullNumber string) {
	fullNumber = ""
	for isNumber(string(line[idx])) {
		fullNumber += string(line[idx])

		idx++
		if idx >= len(line) {
			break
		}
	}
	return fullNumber
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
