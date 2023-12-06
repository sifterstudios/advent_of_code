package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("part1_input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	pattern := regexp.MustCompile(`\d`)
	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		if pattern.MatchString(line) {
			allMatches := pattern.FindAllString(line, -1)
			if allMatches != nil {
				leftNumber := allMatches[0]
				rightNumber := allMatches[len(allMatches)-1]

				combined := leftNumber + rightNumber

				combinedNumber, _ := strconv.Atoi(combined)
				sum += combinedNumber
				println("Sum so far: ", sum)
			}
		}
	}

	println("Final sum: ", sum)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func atoi(combined string) {
	panic("unimplemented")
}
