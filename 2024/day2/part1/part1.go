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
	file, err := os.Open("part1_input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	pattern := regexp.MustCompile(`\d`)
	scanner := bufio.NewScanner(file)

	sumSafeReports := 0
	sumReallySafeReports := 0

	for scanner.Scan() {
		line := scanner.Text()

		if pattern.MatchString(line) {

			stringArray := strings.Split(line, " ")
			intArray := stringsToIntArray(stringArray)

			decreasing := false
			safe := true
			reallySafe := true

			decreasing = intArray[0] > intArray[len(intArray)-1]

			safe = checkIfSafe(intArray, decreasing)
			reallySafe = checkIfReallySafe(intArray)

			if reallySafe {
				sumReallySafeReports++
			}
			if safe {
				sumSafeReports++
			} else {

				problemDampenerSafe := false
				for i := 0; i < len(intArray); i++ {
					filteredArray := append([]int{}, append(intArray[:i], intArray[i+1:]...)...)
					decreasing = filteredArray[0] > filteredArray[len(filteredArray)-1]
					if checkIfSafe(filteredArray, decreasing) {
						problemDampenerSafe = true
						break
					} else {
					}
				}
				if problemDampenerSafe {
					sumSafeReports++
				}
			}
		}
	}

	println("Safe reports:", sumSafeReports)
	println("Really safe reports:", sumReallySafeReports)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func stringifyIntArray(filteredArray []int) string {
	stringifiedArray := ""
	for _, element := range filteredArray {
		stringifiedArray += strconv.Itoa(element)
		stringifiedArray += " "
	}

	return stringifiedArray
}

func checkIfSafe(intArray []int, decreasing bool) bool {
	for i := 1; i < len(intArray); i++ {
		difference := -1

		if decreasing {
			if intArray[i-1] < intArray[i] {
				return false
			}

			difference = intArray[i-1] - intArray[i]
		} else { // increasing
			if intArray[i-1] > intArray[i] {
				return false
			}
			difference = intArray[i] - intArray[i-1]
		}

		if difference > 3 || difference == 0 {
			return false
		}
	}
	return true
}

func checkIfReallySafe(intArray []int) bool {
	left := 0
	right := 1
	problems := 0
	decreasing := -1
	for i := 1; i < len(intArray); i++ {
		if right-left > 2 || problems > 1 {
			return false
		}

		difference := intArray[left] - intArray[right]

		if difference == 0 {
			problems++
			right++
			continue
		}

		if difference > 0 {
			if decreasing == -1 {
				decreasing = 1
			} else if decreasing == 0 {
				problems++
				right++
				continue
			}
		}

		if difference < 0 {
			if decreasing == -1 {
				decreasing = 0
			} else if decreasing == 1 {
				problems++
				right++
				continue
			}
		}

		left++
	}
	return true
}

func intAbs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}

func stringsToIntArray(stringArray []string) []int {
	intArray := []int{}
	for _, element := range stringArray {
		number, _ := strconv.Atoi(element)
		intArray = append(intArray, number)
	}
	return intArray
}
