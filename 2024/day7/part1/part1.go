package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	testValues = []int{}
	equations  = make(map[int][]int)
)

func main() {
	file, err := os.Open("2024/day7/part1/part1_input")
	if err != nil {
		println("Couldn't open file")
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		firstSplit := strings.Split(line, ": ")
		testValues = append(testValues, toInt(firstSplit[0]))
		secondSplit := strings.Split(firstSplit[1], " ")
		equations[toInt(firstSplit[0])] = toIntArray(secondSplit)

		println(line)
	}
	println(len(testValues))
	printIntArray(testValues)
	println(len(equations))
	printIntArray(equations[190])

	// Calculate the total calibration result
	total := 0
	for _, testValue := range testValues {
		numbers := equations[testValue]
		if canProduceTarget(numbers, testValue) {
			total += testValue
		}
	}

	println("Total calibration result:", total)
}

// Recursive function to check if a target can be produced
func canProduceTarget(numbers []int, target int) bool {
	// Start the recursion from the first number
	return evaluate(numbers, 1, numbers[0], target)
}

// Recursive function to try all operator combinations
func evaluate(numbers []int, index, currentValue, target int) bool {
	// Base case: all numbers have been processed
	if index == len(numbers) {
		return currentValue == target
	}

	// Recursive case: try both `+` and `*`
	nextNum := numbers[index]
	return evaluate(numbers, index+1, currentValue+nextNum, target) ||
		evaluate(numbers, index+1, currentValue*nextNum, target)
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		println("Couldn't convert", s, "to int")
	}
	return i
}

func toIntArray(s []string) []int {
	arr := []int{}
	for _, i := range s {
		arr = append(arr, toInt(i))
	}
	return arr
}

func printIntArray(arr []int) {
	for _, i := range arr {
		print(i, ", ")
	}
	println()
}
