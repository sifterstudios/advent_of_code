package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("part1_input")
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
		line := allLines[i]
		sumSchematics = sumOfValidSchematics(line, err, i, allLines, sumSchematics)
		sumGears = sumOfGears(line, allLines)
	}
	println("Sum of schematics: ", sumSchematics)
	println("Sum of gears: ", sumGears)
}

func sumOfGears(line string, allLines []string) int {
	panic("unimplemented")
}

func sumOfValidSchematics(line string, err error, i int, allLines []string, sumSchematics int) int {
	for j := 0; j < len(line); j++ {
		sideHit, aboveHit, belowHit := false, false, false
		if isPeriod(string(line[j])) {
			continue
		}
		fullNumberString := getFullNumber(line, j)

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
			println(fullNumberString + " is a schematic!")
			sumSchematics += fullNumber
		} else if fullNumberString != "" {
			println(fullNumberString + " is NOT a schematic!")
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
				isStar = true
				isSymbol = true
			} else {
				isSymbol = true
			}
		}
	}
	return isStar, isSymbol
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

func getFullNumber(line string, j int) (fullNumber string) {
	fullNumber = ""
	for isNumber(string(line[j])) {
		fullNumber += string(line[j])
		j++
		if j >= len(line) {
			break
		}
	}
	return fullNumber
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func isPeriod(s string) bool {
	return s == "."
}
