package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("part1_input")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			print("Couldn't close test input file")
		}
	}(file)

	scanner := bufio.NewScanner(file)

	sumPossibleGames := 0
	maxBlue, maxRed, maxGreen := 14, 12, 13

	for scanner.Scan() {
		line := scanner.Text()
		gameNumberAndResults := strings.Split(line, ": ")
		gameNumberString := strings.Replace(gameNumberAndResults[0], "Game ", "", -1)
		gameNumber, _ := strconv.Atoi(gameNumberString)

		resultsString := gameNumberAndResults[1]
		allHandfuls := strings.Split(resultsString, "; ")

		for _, handful := range allHandfuls {
			colorsInHandful := strings.Split(handful, ", ")
			for _, color := range colorsInHandful {
				if strings.Contains(color, "blue") {
					numberOfCubes := getNumberOfCubes(color, " blue")

					if numberOfCubes > maxBlue {
						gameNumber = 0
					}
				}
				if strings.Contains(color, "red") {
					numberOfCubes := getNumberOfCubes(color, " red")

					if numberOfCubes > maxRed {
						gameNumber = 0
					}
				}
				if strings.Contains(color, "green") {
					numberOfCubes := getNumberOfCubes(color, " green")

					if numberOfCubes > maxGreen {
						gameNumber = 0
					}
				}
			}
		}

		println(gameNumber)
		sumPossibleGames += gameNumber
	}
	println("Sum possible games", sumPossibleGames)
}

func getNumberOfCubes(numberAndColor, color string) (result int) {
	onlyNumber := strings.Split(numberAndColor, color)
	digit, _ := strconv.Atoi(onlyNumber[0])
	println("Digit: ", digit)
	println("Color: ", numberAndColor)
	return digit
}
