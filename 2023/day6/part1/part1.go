package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	times              []int
	distances          []int
	possibleRecords    []int
	sumPossibleRecords int
)

func main() {
	file, err := os.Open("./2023/input/day6_part1_input")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)

	scanner := bufio.NewScanner(file)

	grabInput(scanner)

	for i, time := range times {
		distance := distances[i]
		recordCount := 0

		for msWaited := 0; msWaited < time; msWaited++ {
			newTime := time - msWaited
			if (msWaited * newTime) > distance {
				recordCount++
				// println("msWaited: ", msWaited, "newTime: ", newTime, "distance: ", distance, "recordCount: ", recordCount)
			}
		}
		possibleRecords = append(possibleRecords, recordCount)

	}
	sumPossibleRecords = (possibleRecords[0] * possibleRecords[1] * possibleRecords[2] * possibleRecords[3])
	println("sumPossibleRecords: ", sumPossibleRecords)
}

func grabInput(scanner *bufio.Scanner) {
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Time") {
			allTimesStringWithoutText := strings.Split(line, ":")
			allTimesString := strings.Split(allTimesStringWithoutText[1], " ")
			for _, timeString := range allTimesString {
				v, err := strconv.Atoi(timeString)
				if err == nil {
					times = append(times, v)
				}

			}

		}
		if strings.Contains(line, "Distance") {
			allDistanceStringWithoutText := strings.Split(line, ":")
			allDistanceString := strings.Split(allDistanceStringWithoutText[1], " ")
			for _, distanceString := range allDistanceString {
				v, err := strconv.Atoi(distanceString)
				if err == nil {
					distances = append(distances, v)
				}

			}
		}
	}
}
