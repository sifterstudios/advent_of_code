package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("./2024/day3/part1/part1_input")
	if err != nil {
		println("Error opening file", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	isMulEnabled := true // Start with mul instructions enabled

	reMul := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	// Iterate over each line
	for scanner.Scan() {
		line := scanner.Text()

		// Combine all matches (do, don't, mul) and sort them by their position in the line
		combinedPattern := `do\(\)|don't\(\)|mul\(\d{1,3},\d{1,3}\)`
		reCombined := regexp.MustCompile(combinedPattern)
		matches := reCombined.FindAllStringIndex(line, -1)

		// Iterate over each match in order
		for _, match := range matches {
			token := line[match[0]:match[1]] // Extract matched substring

			switch {
			case token == "do()":
				isMulEnabled = true
			case token == "don't()":
				isMulEnabled = false
			case reMul.MatchString(token) && isMulEnabled:
				submatches := reMul.FindStringSubmatch(token)
				num1, _ := strconv.Atoi(submatches[1])
				num2, _ := strconv.Atoi(submatches[2])
				sum += (num1 * num2)
			}
		}
	}

	println("Sum: ", sum)
}
