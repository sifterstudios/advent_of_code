package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("/2024/day2/part2/part2_input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	pattern := regexp.MustCompile(`\d`)
	scanner := bufio.NewScanner(file)

	var leftNumbers []int
	var rightNumbers []int
	countInRightSum := 0

	for scanner.Scan() {
		line := scanner.Text()

		if pattern.MatchString(line) {
			// find first number before a space

			stringArray := strings.Split(line, " ")
			leftNumber := stringArray[0]
			rightNumber := stringArray[len(stringArray)-1]

			l, _ := strconv.Atoi(leftNumber)
			r, _ := strconv.Atoi(rightNumber)

			leftNumbers = append(leftNumbers, l)
			rightNumbers = append(rightNumbers, r)
		}
	}

	sort.Ints(leftNumbers)
	sort.Ints(rightNumbers)

	for i := 0; i < len(leftNumbers); i++ {
		leftnum := leftNumbers[i]
		countInRightSlice := 0
		for _, num := range rightNumbers {
			if num == leftnum {
				countInRightSlice++
			}
		}

		countInRightSum += leftnum * countInRightSlice

	}

	println("Count in right slice sum:", countInRightSum)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
