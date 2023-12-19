package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type desertMap struct {
	location int
	left     *desertMap
	right    *desertMap
}

var (
	directions []int
	allMaps    []desertMap
)

func main() {
	file, err := os.Open("./2023/input/day8_part1_input")
	if err != nil {
		println("Open file failed")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			println("Could not close file.")
		}
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	getDirections(scanner.Text())
	scanner.Scan()
	for scanner.Scan() {
		//foundLeft, foundRight := false, false
		var currentMap desertMap
		line := scanner.Text()
		println(line)
		locationAndDirections := strings.Split(line, "=")
		foundLocation := false

		locationInt := 0
		for i := 0; i < 3; i++ {
			locationInt += int(locationAndDirections[0][i])
		}
		currentMap.location = locationInt

		directions := strings.Split(locationAndDirections[1], ",")
		left := getDirection(directions[0])
		right := getDirection(directions[1])

		for _, loc := range allMaps {
			if loc.location == locationInt {
				foundLocation = true
			}
		}

		if !foundLocation {
			allMaps = append(allMaps, currentMap)
		}

		for i := 0; i < len(allMaps); i++ {
			if allMaps[i].location == locationInt {
				if locationInt == left && locationInt == right {
					allMaps[i].left = &allMaps[i]
					allMaps[i].right = &allMaps[i]
				}
				if allMaps[i].left != nil && allMaps[i].right != nil {
					continue
				}

				if allMaps[i].left == nil {
					println("Couldn't find left on ", locationInt)
					newLeft := desertMap{location: left}
					allMaps = append(allMaps, newLeft)
					allMaps[i].left = &allMaps[len(allMaps)-1]
				}

				if allMaps[i].right == nil {
					println("Couldn't find right on ", locationInt)
					newRight := desertMap{location: right}
					allMaps = append(allMaps, newRight)
					allMaps[i].right = &allMaps[len(allMaps)-1]
				}
			}
		}
	}
	fmt.Printf("allmaps: %v\n", allMaps)

	findShortestPath(&allMaps[0], 0, 0)
}

func findShortestPath(currentMap *desertMap, steps, directionIndex int) {
	var tempMap *desertMap
	if currentMap.location == 270 {
		println("Found ZZZ, took ", steps, " steps")
		return
	}
	if directionIndex == len(directions)-1 {
		directionIndex = 0
	}

	if directions[directionIndex] == 0 {
		tempMap = currentMap.left
	} else {
		tempMap = currentMap.right
	}
	steps++
	directionIndex++
	findShortestPath(tempMap, steps, directionIndex)
}

func getDirection(s string) int {
	directionInt := 0
	for _, c := range s {
		if c == '(' || c == ')' || c == ' ' {
			continue
		}
		directionInt += int(c)
	}
	return directionInt
}

func getDirections(s string) {
	for _, c := range s {
		if c == 'L' {
			directions = append(directions, 0)
		}
		if c == 'R' {
			directions = append(directions, 1)
		}

	}
}
