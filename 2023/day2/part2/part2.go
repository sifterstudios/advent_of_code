package main

import (
	"bufio"
	"os"
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

	scanner := bufio.NewScanner(file)

	sumPossibleGames := 0
	maxBlue, maxRed, maxGreen := 999, 999, 999
	minPossibleBlue, minPossibleRed, minPossibleGreen := 0, 0, 0
	currentBlue, currentRed, currentGreen := 0, 0, 0
	sumPowerCubes := 0

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
					} else {
						if numberOfCubes > minPossibleBlue {
							minPossibleBlue = numberOfCubes
							currentBlue = numberOfCubes
						}
					}
				}
				if strings.Contains(color, "red") {
					numberOfCubes := getNumberOfCubes(color, " red")

					if numberOfCubes > maxRed {
						gameNumber = 0
					} else {
						if numberOfCubes > minPossibleRed {
							minPossibleRed = numberOfCubes
							currentRed = numberOfCubes
						}
					}
				}
				if strings.Contains(color, "green") {
					numberOfCubes := getNumberOfCubes(color, " green")

					if numberOfCubes > maxGreen {
						gameNumber = 0
					} else {
						if numberOfCubes > minPossibleGreen {
							minPossibleGreen = numberOfCubes
							currentGreen = numberOfCubes
						}
					}
				}
			}
		}

		println(gameNumber)
		if gameNumber != 0 {
			sumPowerCubes += currentBlue * currentRed * currentGreen
		}
		currentBlue = 0
		currentRed = 0
		currentGreen = 0
		minPossibleBlue = 0
		minPossibleRed = 0
		minPossibleGreen = 0
	}
	println("Sum possible games", sumPossibleGames)
	println("Sum power cubes", sumPowerCubes)
}

func getNumberOfCubes(numberAndColor, color string) (result int) {
	onlyNumber := strings.Split(numberAndColor, color)
	digit, _ := strconv.Atoi(onlyNumber[0])
	println("Digit: ", digit)
	println("Color: ", numberAndColor)
	return digit
}
