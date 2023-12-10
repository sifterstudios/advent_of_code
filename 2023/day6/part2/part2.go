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
	file, err := os.Open("./2023/input/day6_part2_input")
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

	time := times[0]
	distance := distances[0]
	recordCount := 0

	for msWaited := 0; msWaited < time; msWaited++ {
		newTime := time - msWaited
		if (msWaited * newTime) > distance {
			recordCount++
		}
	}
	possibleRecords = append(possibleRecords, recordCount)

	println("sumPossibleRecords: ", recordCount)
}

func grabInput(scanner *bufio.Scanner) {
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Time") {
			allTimesStringWithoutText := strings.Split(line, ":")
			allTimesString := strings.Split(allTimesStringWithoutText[1], " ")
			combinedTimes := ""
			v := -1
			for _, timeString := range allTimesString {
				_, err := strconv.Atoi(timeString)
				if err == nil {
					combinedTimes += timeString
				}

			}
			v, _ = strconv.Atoi(combinedTimes)
			times = append(times, v)
			println("times: ", times[0])

		}
		if strings.Contains(line, "Distance") {
			allDistanceStringWithoutText := strings.Split(line, ":")
			allDistancesString := strings.Split(allDistanceStringWithoutText[1], " ")
			combinedDistances := ""
			v := -1
			for _, distanceString := range allDistancesString {
				_, err := strconv.Atoi(distanceString)
				if err == nil {
					combinedDistances += distanceString
				}

			}
			v, _ = strconv.Atoi(combinedDistances)
			distances = append(distances, v)
			println("distance: ", distances[0])

		}
	}
}
