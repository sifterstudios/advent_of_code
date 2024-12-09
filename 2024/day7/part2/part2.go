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
	// Open the input file
	file, err := os.Open("2024/day7/part2/part2_input")
	if err != nil {
		println("Couldn't open file")
		return
	}
	defer file.Close()

	// Read the input file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		firstSplit := strings.Split(line, ": ")
		testValue := toInt(firstSplit[0])
		testValues = append(testValues, testValue)
		secondSplit := strings.Split(firstSplit[1], " ")
		equations[testValue] = toIntArray(secondSplit)
	}

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

	// Recursive case: try all operators
	nextNum := numbers[index]
	// Addition
	if evaluate(numbers, index+1, currentValue+nextNum, target) {
		return true
	}
	// Multiplication
	if evaluate(numbers, index+1, currentValue*nextNum, target) {
		return true
	}
	// Concatenation
	concatenatedValue := concat(currentValue, nextNum)
	if evaluate(numbers, index+1, concatenatedValue, target) {
		return true
	}
	return false
}

// Concatenation function
func concat(a, b int) int {
	// Convert integers to strings and concatenate them
	concatenatedStr := strconv.Itoa(a) + strconv.Itoa(b)
	// Convert back to integer
	concatenatedValue, err := strconv.Atoi(concatenatedStr)
	if err != nil {
		println("Error concatenating", a, "and", b)
	}
	return concatenatedValue
}

// Utility to convert a string to an integer
func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		println("Couldn't convert", s, "to int")
	}
	return i
}

// Utility to convert a slice of strings to a slice of integers
func toIntArray(s []string) []int {
	arr := []int{}
	for _, i := range s {
		arr = append(arr, toInt(i))
	}
	return arr
}
