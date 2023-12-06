package main

import (
	"fmt"
	"os"
)

var result int

func main() {
	result = 0
	file, err := os.Open("./part1_input")
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}
	input := make([]byte, 1000000)

	_, readError := file.Read(input)
	if readError != nil {
		fmt.Printf("Error reading file: %v", readError)
	}

	for _, c := range input {
		switch c {
		case '(':
			result++
			break
		case ')':
			result--
			break
		}
	}

	fmt.Printf("Result: %d\n", result)
}
