package main

import (
	"bufio"
	"os"
	"strconv"
)

type gear struct {
	starLocation       []int
	isNumberAbove      bool
	isNumberBelow      bool
	numberAbove        int
	numberBelow        int
	lastSearchedIndex  int
	goingBackwards     bool
	wasLastIndexNumber bool
}

func newGear() gear {
	return gear{
		starLocation: []int{-1, -1}, isNumberAbove: false,
		isNumberBelow: false, numberAbove: -1, numberBelow: -1,
		lastSearchedIndex: 0, goingBackwards: false, wasLastIndexNumber: false,
	}
}

func main() {
	file, err := os.Open("part2_input")
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
	getSumOfGears(allLines, 0, &sumGears)
	println("Sum of schematics: ", sumSchematics)
	println("Sum of gears: ", sumGears)
}

func getSumOfGears(allLines []string, lineNumber int, sumGears *int, gearOpt ...gear) {
	gear := gear{}
	if gearOpt == nil {
		gear = newGear()
	} else {
		gear = gearOpt[0]
	}

	// End case
	if (gear.lastSearchedIndex > len(allLines[lineNumber])) &&
		lineNumber == len(allLines)-1 {
		println("End of input, no more gears to find")
		return
	}

	// Handle end of line
	if gear.lastSearchedIndex == len(allLines[lineNumber])-2 {
		println("End of line, moving to next line")
		gear.lastSearchedIndex = 0
		getSumOfGears(allLines, lineNumber+1, sumGears, gear)
		return
	} else {
	}

	gear.lastSearchedIndex++ // I cheated, I know the first index is never a star
	// Search for star
	char := string(allLines[lineNumber][gear.lastSearchedIndex])

	// Handle if I don't find star
	if char != "*" {
		getSumOfGears(allLines, lineNumber, sumGears, gear)
		return
	}

	println("Found star at line", lineNumber, "index", gear.lastSearchedIndex)
	lineAbove := allLines[lineNumber-1]
	lineBelow := allLines[lineNumber+1]
	i := gear.lastSearchedIndex

	// Search for number above and one to both sides
	if !isNumber(string(lineAbove[i])) &&
		!isNumber(string(lineAbove[i-1])) &&
		!isNumber(string(lineAbove[i+1])) {
		println("No number above, moving on...")
		getSumOfGears(allLines, lineNumber, sumGears, gear)
		return
	}
	println("Found number above at line", lineNumber-1, "index", gear.lastSearchedIndex)
	// Search for number below and one to both sides
	if !isNumber(string(lineAbove[i])) &&
		!isNumber(string(lineAbove[i-1])) &&
		!isNumber(string(lineAbove[i+1])) {
		println("No number below, moving on...")
		getSumOfGears(allLines, lineNumber, sumGears, gear)
		return
	}
	println("Found number both above and below")
	// If number above and below, find the whole number by going first forwards, then backwards until i hit a symbol
	i1, i2 := findWhereToStartSearching(lineAbove, i)
	numberAbove := getFullNumber(lineAbove, i1, false) + getFullNumber(lineAbove, i2, true)
	println("Number above is", numberAbove)

	i1, i2 = findWhereToStartSearching(lineBelow, i)
	numberBelow := getFullNumber(lineBelow, i1, false) + getFullNumber(lineBelow, i2, true)
	println("Number below is", numberBelow)

	// If there's a number both above and below, multiply them and add to sumGears and continue searching
	return
}

func findWhereToStartSearching(lineAbove string, i int) (i1 int, reverse bool) {
	if isNumber(string(lineAbove[i-1])) {
		return i, i - 1
	}
	if isNumber(string(lineAbove[i])) {
		return i, i - 1
	}
	if isNumber(string(lineAbove[i+1])) {
		return i + 1, i
	}
	return -1, -1 // This will cause intended panic
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
