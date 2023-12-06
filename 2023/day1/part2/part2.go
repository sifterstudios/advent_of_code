package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("part2_input")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			print("Couldn't close test input file")
		}
	}(file)

	pattern := regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)
	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		allMatches := pattern.FindAllString(line, -1)

		if allMatches != nil {
			leftmostMatch := allMatches[0]
			rightmostMatch := allMatches[len(allMatches)-1]

			updatedRightMatch := checkForDudusRightSide(rightmostMatch, line)

			leftNumberString := convertToDigitString(leftmostMatch)
			rightNumberString := convertToDigitString(updatedRightMatch)

			combined := leftNumberString + rightNumberString

			combinedNumber, err := strconv.Atoi(combined)
			if err != nil {
				println("Error converting ", combined, " to int")
			}
			sum += combinedNumber

		} else {
			println("No matches found")
		}
	}

	println("Final sum: ", sum)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func convertToDigitString(match string) (result string) {
	newResult := match
	switch match {
	case "one":
		newResult = "1"
		break
	case "two":
		newResult = "2"
		break
	case "three":
		newResult = "3"
		break
	case "four":
		newResult = "4"
		break
	case "five":
		newResult = "5"
		break
	case "six":
		newResult = "6"
		break
	case "seven":
		newResult = "7"
		break
	case "eight":
		newResult = "8"
		break
	case "nine":
		newResult = "9"
		break
	}
	return newResult
}

func checkForDudusRightSide(rightmostMatch, line string) (updatedRightMatch string) {
	if _, err := strconv.Atoi(rightmostMatch); err == nil {
		return rightmostMatch
	}

	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	matchString := ""
	for i := len(line) - 1; i >= 0; i-- {
		matchString = string(line[i]) + matchString

		for _, number := range numbers {
			if strings.Contains(matchString, number) {
				return number
			}
		}

	}
	return "!"
}
