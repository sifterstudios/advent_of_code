package main

import (
	"bufio"
	"log"
	"os"
)

var (
	word         = []rune{'X', 'M', 'A', 'S'}
	lines        = [141]line{}
	sumWordCount = 0
	lineNumber   = 0
	maxXLength   = 0
	lengthOfWord = 4
)

type Coordinate struct {
	x, y int
	rune rune
}

type line struct {
	coords []Coordinate
}

func main() {
	file, err := os.Open("2024/day4/part1/part1_input")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		runeLine := []rune(line)
		// 1. Gather the possible characters with x and y coordinates
		for i, char := range runeLine {
			lines[lineNumber].coords = append(lines[lineNumber].coords, Coordinate{i, lineNumber, char})

			// 2. Search the word horizontally, in both directions
			if char == word[0] {
				checkForwardsSpelling(i, runeLine)
			}

			if char == word[len(word)-1] {
				checkBackwardsSpelling(i, runeLine)
			}
			if i+1 > maxXLength {
				maxXLength = i + 1
			}
		}
		lineNumber++
	}

	// 3. Search the word vertically, in both directions

	for i := 0; i < maxXLength; i++ {
		runeLine := getVerticalRunes(i)
		for j := 0; j < len(runeLine); j++ {

			if runeLine[j] == word[0] {
				checkForwardsSpelling(j, runeLine)
			}

			if runeLine[j] == word[len(word)-1] {
				checkBackwardsSpelling(j, runeLine)
			}
		}
	}

	// 4. Search the word diagonally, in all four directions, but only with rising index so we don't count twice
	for x := 0; x < maxXLength; x++ {
		for y := 0; y < lineNumber; y++ {
			if lines[y].coords[x].rune == word[0] && checkForwardsDiagonals(x, y) {
				// sumWordCount++
			}
			if lines[y].coords[x].rune == word[len(word)-1] && checkBackwardsDiagonals(x, y) {
				// sumWordCount++
			}
		}
	}
	// 5. Add the word count to the sumWordCount
	println("Word count: ", sumWordCount)
}

func checkBackwardsDiagonals(x, y int) bool {
	maxX := maxXLength
	maxY := lineNumber
	localWordCount := 0

	// CheckUpwards
	if x+lengthOfWord < maxX && y-lengthOfWord >= 0 {
		for i := 1; i < len(word); i++ {
			locX, locY := x+i, y-i
			if getCharFromCoordinate(locX, locY) != word[len(word)-1-i] {
				break
			}
			if i == len(word)-1 {
				localWordCount++
			}
		}
	}

	// CheckDownwards
	if x+lengthOfWord < maxX && y+len(word)-1 < maxY {
		for i := 1; i < len(word); i++ {
			locX, locY := x+i, y+i
			if getCharFromCoordinate(locX, locY) != word[len(word)-1-i] {
				break
			}
			if i == len(word)-1 {
				localWordCount++
			}
		}
	}

	if localWordCount > 0 {
		sumWordCount += localWordCount
		return true
	}

	return false
}

func checkForwardsDiagonals(x, y int) bool {
	maxX := maxXLength
	maxY := lineNumber
	localWordCount := 0

	// CheckUpwards
	if x+len(word)-1 < maxX && y-len(word)-1 >= 0 {
		for i := 1; i < len(word); i++ {
			locX, locY := x+i, y-i
			if getCharFromCoordinate(locX, locY) != word[i] {
				break
			}
			if i == len(word)-1 {
				localWordCount++
			}
		}
	}

	// CheckDownwards
	if x+len(word)-1 < maxX && y+len(word)-1 <= maxY {
		for i := 1; i < len(word); i++ {
			locX, locY := x+i, y+i

			if getCharFromCoordinate(locX, locY) != word[i] {
				break
			}

			if i == len(word)-1 {
				localWordCount++
			}
		}
	}

	if localWordCount > 0 {
		sumWordCount += localWordCount
		return true
	}

	return false
}

func checkBackwardsSpelling(i int, runeLine []rune) {
	for j := 0; j < len(word); j++ {
		if i+j >= len(runeLine) || runeLine[i+j] != word[len(word)-1-j] {
			break
		}
		if j == len(word)-1 {
			sumWordCount++
		}
	}
}

func checkForwardsSpelling(indexInRuneLine int, runeLine []rune) {
	for j := 0; j < len(word); j++ {
		if indexInRuneLine+j >= len(runeLine) || runeLine[indexInRuneLine+j] != word[j] {
			break
		}
		if j == len(word)-1 {
			sumWordCount++
		}
	}
}

func getCharFromCoordinate(x, y int) rune {
	for i := 0; i < len(lines[y].coords); i++ {
		if lines[y].coords[i].x == x && lines[y].coords[i].y == y {
			return lines[y].coords[i].rune
		}
	}

	return ' '
}

func getVerticalRunes(xIndex int) []rune {
	var verticalRunes []rune
	for line := 0; line < len(lines); line++ {
		c := getCharFromCoordinate(xIndex, line)
		if c != ' ' {
			verticalRunes = append(verticalRunes, c)
		}
	}
	// println("Row number: ", xIndex, "Vertical runes: ", string(verticalRunes))
	return verticalRunes
}
