package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

var (
	pageOrderingRules                     = []pageOrderRule{}
	doneParsingRules                      = false
	sumCorrectMiddlePages                 = 0
	sumIncorrectMiddlePagesAfterAdjusting = 0
)

type pageOrderRule struct {
	before, after int
}

func main() {
	file, err := os.Open("2024/day5/part1/part1_input")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			doneParsingRules = true
			continue
		}
		if !doneParsingRules {
			ruleArray := getIntArrayFromLine(line, "|")
			pageOrderingRules = append(pageOrderingRules, pageOrderRule{ruleArray[0], ruleArray[1]})
			continue
		}

		intArray := getIntArrayFromLine(line, ",")
		correctOrder := checkOrderingOfPages(intArray)
		if correctOrder {
			arrayLength := len(intArray)
			middleIndex := divideAndRoundUpEven(arrayLength - 1)
			sumCorrectMiddlePages += intArray[middleIndex]
		} else {
			adjustedArray := reorderPages(intArray)
			printArray(adjustedArray)
			arrayLength := len(adjustedArray)
			middleIndex := divideAndRoundUpEven(arrayLength - 1)
			sumIncorrectMiddlePagesAfterAdjusting += adjustedArray[middleIndex]
		}
	}
	println("Sum of middle pages is: ", sumCorrectMiddlePages)
	println("Sum of middle pages after adjusting is: ", sumIncorrectMiddlePagesAfterAdjusting)
}

func printArray(adjustedArray []int) {
	arrayString := ""
	for _, num := range adjustedArray {
		arrayString += strconv.Itoa(num) + ","
	}
}

func reorderPages(update []int) []int {
	inDegree := make(map[int]int)  // Track incoming edges
	graph := make(map[int][]int)   // Adjacency list graph
	inUpdate := make(map[int]bool) // Pages in the current update

	// Track pages present in this specific update
	for _, page := range update {
		inUpdate[page] = true
	}

	// Build the graph and in-degree map only for relevant rules
	for _, rule := range pageOrderingRules {
		from, to := rule.before, rule.after
		if inUpdate[from] && inUpdate[to] {
			graph[from] = append(graph[from], to)
			inDegree[to]++
			if _, exists := inDegree[from]; !exists {
				inDegree[from] = 0
			}
		}
	}

	// Topological sort using Kahn's algorithm
	var queue []int
	for _, page := range update {
		if inDegree[page] == 0 {
			queue = append(queue, page)
		}
	}

	var sorted []int
	for len(queue) > 0 {
		sort.Ints(queue) // Ensure lexical ordering of nodes with the same priority
		current := queue[0]
		queue = queue[1:]
		sorted = append(sorted, current)
		for _, neighbor := range graph[current] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// If any pages are missing from the sorted result, add them in original order
	visited := make(map[int]bool)
	for _, page := range sorted {
		visited[page] = true
	}

	for _, page := range update {
		if !visited[page] {
			sorted = append(sorted, page)
		}
	}

	fmt.Println("Adjusted array:", sorted)
	return sorted
}

func getIntArrayFromLine(line string, split string) []int {
	stringArray := strings.Split(line, split)
	intArray := []int{}
	for _, str := range stringArray {
		intValue, err := strconv.Atoi(str)
		if err != nil {
			println("Error converting string to int")
		}
		intArray = append(intArray, intValue)
	}
	return intArray
}

func checkOrderingOfPages(pages []int) bool {
	// We need to checx if both the before and after is present in the array
	// if this is the case, THEN we check if the before page is before the after page
	// if this is true for all the rules, we return true
	// Otherwise, return false

	for _, rule := range pageOrderingRules {
		beforePage := rule.before
		afterPage := rule.after
		if !isPagePresent(pages, beforePage) || !isPagePresent(pages, afterPage) {
			continue // it's fine if one rule does not have corresponding pages
		}
		if !isPageBefore(pages, beforePage, afterPage) {
			return false
		}
	}
	return true
}

func isPageBefore(pages []int, beforePage, afterPage int) bool {
	beforeIndex := getIndexOfElementInArray(pages, beforePage)
	afterIndex := getIndexOfElementInArray(pages, afterPage)
	return beforeIndex < afterIndex
}

func isPagePresent(intArray []int, pageTosearchFor int) bool {
	return slices.Contains(intArray, pageTosearchFor)
}

func divideAndRoundUpEven(numerator int) int {
	result := float64(numerator) / float64(2)

	// Check if the result is a half (x.5) and round up, otherwise use ceil for other fractions
	if math.Mod(result, 1) == 0.5 {
		return int(result) + 1
	}
	return int(math.Ceil(result))
}

func getIndexOfElementInArray(array []int, element int) int {
	return slices.Index(array, element)
}
