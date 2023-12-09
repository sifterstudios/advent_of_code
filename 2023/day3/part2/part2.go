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
	file, err := os.Open("./input/day3_part2_input")
	if err != nil {
		println("Error opening file")
	}

	scanner := bufio.NewScanner(file)
	allLines := []string{}
	defer file.Close()

	for scanner.Scan() {
		allLines = append(allLines, scanner.Text())
	}

	sumSchematics := 0
	sumGears := 0

	for i := 0; i < len(allLines); i++ {
		sumSchematics = getSumOfValidSchematics(allLines, i, err, sumSchematics)
	}
	getSumOfGears(allLines, &sumGears)
	println("Sum of schematics: ", sumSchematics)
	println("Sum of gears: ", sumGears)
}

func getSumOfGears(allLines []string, sumGears *int) {
	numberEntries := []numberEntry{}
	starEntries := []starEntry{}

	for i := 0; i < len(allLines); i++ {
		for j := 0; j < len(allLines[i]); j++ {
			if isNumber(string(allLines[i][j])) {
				numberEntries = append(numberEntries, numberEntry{i, j, -1, -1})
				// Now we've fonud the start of a number, find the end of it
				for j < len(allLines[i]) && isNumber(string(allLines[i][j])) {
					j++
				}
				// Was no longer a number or end of line
				numberEntries[len(numberEntries)-1].end = j - 1 // the -1 is questionable, remember this in test
				numberEntries[len(numberEntries)-1].fullNumber, _ = strconv.Atoi(getFullNumber(allLines[i], j-1, true))
			}

			if string(allLines[i][j]) == "*" {
				println("Found a star at ", i, j)
				starEntries = append(starEntries, starEntry{i, j})

			}
		}
	}
	println("Collected numbers and stars")
	for star := 0; star < len(starEntries); star++ {
		numbersFound := checkForNumberInAllDirections(starEntries[star], numberEntries)
		if len(numbersFound) == 2 {
			// Found a gear
			*sumGears += numbersFound[0] * numbersFound[1]
		}
	}
}

func checkForNumberInAllDirections(star starEntry, numberEntries []numberEntry) (numbersFound []int) {
	idx, line := star.idx, star.line

	// ABOVE
	above := numberInSquare(numberEntries, line-1, idx)
	if above != -1 {
		numbersFound = append(numbersFound, above)
	}

	below := numberInSquare(numberEntries, line+1, idx)
	if below != -1 {
		numbersFound = append(numbersFound, below)
	}

	sides := numberInSquare(numberEntries, line, idx)
	if sides != -1 {
		numbersFound = append(numbersFound, sides)
	}

	if len(numbersFound) > 2 {
		println("Something is terribly wrong, found more than 2 numbers at location ", line, idx)
	}

	return numbersFound
}

func numberInSquare(numberEntries []numberEntry, line int, idx int) (number int) {
	for number := 0; number < len(numberEntries); number++ {
		numberEntry := numberEntries[number]
		if numberEntry.line == line {
			for i := idx - 1; i <= idx+1; i++ {
				if numberEntry.start <= i && i <= numberEntry.end {
					return numberEntry.fullNumber
				}
			}
		}
	}
	return -1
}

func getSumOfValidSchematics(allLines []string, i int, err error, sumSchematics int) int {
	line := allLines[i]
	for j := 0; j < len(line); j++ {
		sideHit, aboveHit, belowHit := false, false, false
		if isPeriod(string(line[j])) {
			continue
		}

		fullNumberString := getFullNumber(line, j, false)

		fullNumber := 0
		if fullNumberString != "" {
			fullNumber, err = strconv.Atoi(fullNumberString)
			if err != nil {
				println("Error converting string to int", err.Error())
			}
		}

		if fullNumberString != "" && !isPeriod(fullNumberString) {
			sideHit = checkForSymbolToTheSides(j, len(fullNumberString), line)

			if i != len(allLines)-1 {
				belowHit, _ = checkForSymbolAboveOrBelow(i, j, len(fullNumberString), allLines[i+1])
			}

			if i != 0 {
				aboveHit, _ = checkForSymbolAboveOrBelow(i, j, len(fullNumberString), allLines[i-1])
			}

			j += len(fullNumberString)
		}

		if sideHit || aboveHit || belowHit {
			sumSchematics += fullNumber
		}
	}
	return sumSchematics
}

func checkForSymbolAboveOrBelow(lineNumber, start, end int, line string) (isSymbol, isStar bool) {
	for i := start - 1; i <= start+end; i++ {
		if i < 0 || i >= len(line) {
			continue
		}

		if !isPeriod(string(line[i])) && !isNumber(string(line[i])) {
			if string(line[i]) == "*" {
				isSymbol = true
				isStar = true
			} else {
				isSymbol = true
			}
		}
	}
	return isSymbol, isStar
}

func checkForSymbolToTheSides(start, end int, line string) bool {
	if start != 0 {
		if !isPeriod(string(line[start-1])) && !isNumber(string(line[start-1])) {
			return true
		}
	}
	if start+end < len(line) {
		if !isPeriod(string(line[start+end])) && !isNumber(string(line[start+end])) {
			return true
		}
	}
	return false
}

func getFullNumber(line string, idx int, backwards bool) (fullNumber string) {
	fullNumber = ""
	for isNumber(string(line[idx])) {
		fullNumber += string(line[idx])

		if backwards && idx <= 0 {
			break
		}
		if !backwards && idx >= len(line) {
			break
		}
		if backwards {
			idx--
		} else {
			idx++
		}
	}
	if backwards {
		// reverse string
		return Reverse(fullNumber)
	} else {
		return fullNumber
	}
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func isPeriod(s string) bool {
	return s == "."
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
