package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./2024/day2/part2/part2_input")
	if err != nil {
		println("Error opening file")
	}

	scanner := bufio.NewScanner(file)
	allLines := []string{}
	sumSafeLevels := 0
	defer file.Close()

	for scanner.Scan() {
		allLines = append(allLines, scanner.Text())
	}

	for i := 0; i < len(allLines); i++ {
		intArray := splitLineIntoNumbers(allLines[i])
		if reportIsSafe(intArray) {
			sumSafeLevels++
		} else {
			if checkProblemDampeners(intArray) {
				sumSafeLevels++
			}
		}
	}

	println("Safe levels: ", sumSafeLevels)
}

func checkProblemDampeners(intArray []int) bool {
	for i := 0; i < len(intArray); i++ {
		if reportIsSafe(getSliceWithoutIndex(intArray, i)) {
			return true
		}
	}
	return false
}

func getSliceWithoutIndex(intArray []int, i int) []int {
	filteredArray := []int{}
	for j := 0; j < len(intArray); j++ {
		if j != i {
			filteredArray = append(filteredArray, intArray[j])
		}
	}
	return filteredArray
}

func reportIsSafe(intArray []int) bool {
	decreasing := intArray[0] > intArray[len(intArray)-1]

	left := 0
	right := 1

	for i := 0; i < len(intArray); i++ {
		// Each number in the line
		for num := 0; num < len(intArray); num++ {
			if right > len(intArray)-1 {
				break
			}

			leftNumber := intArray[left]
			rightNumber := intArray[right]

			if leftNumber == rightNumber {
				return false
			}

			if decreasing {
				if leftNumber < rightNumber {
					return false
				}

				if leftNumber-rightNumber > 3 {
					return false
				}
			} else { // increasing
				if leftNumber > rightNumber {
					return false
				}
				if rightNumber-leftNumber > 3 {
					return false
				}
			}
			left++
			right++
		}
	}

	return true
}

func splitLineIntoNumbers(s string) []int {
	split := strings.Split(s, " ")
	intArray := []int{}

	for i := 0; i < len(split); i++ {
		intArray = append(intArray, convertStringToInt(split[i]))
	}

	return intArray
}

func convertStringToInt(s string) int {
	number, err := strconv.Atoi(s)
	if err != nil {
		println("Error converting string to int")
	}
	return number
}
