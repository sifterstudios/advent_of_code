package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type almanacEntry struct {
	seed        int
	soil        int
	fertilizer  int
	water       int
	light       int
	temperature int
	humidity    int
	location    int
}

type seedEntry struct {
	rangeStart  int
	rangeLength int
}

type gardenMap struct {
	destinationRangeStart int
	sourceRangeStart      int
	rangeLength           int
}

var (
	almanac                  []almanacEntry
	seeds                    []seedEntry
	seedToSoilMap            []gardenMap
	soilToFertilizerMap      []gardenMap
	fertilizerToWaterMap     []gardenMap
	waterToLightMap          []gardenMap
	lightToTemperatureMap    []gardenMap
	temperatureToHumidityMap []gardenMap
	humidityToLocationMap    []gardenMap
	lowestLocation           int
	allInputLines            []string
)

func main() {
	file, err := os.Open("./2023/input/day5_part1_input")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	getInputLines(scanner)
	lowestLocation = 1000000000

	for i := 0; i < len(allInputLines); i++ {
		currentLine := allInputLines[i]
		grabSeeds(i, currentLine)
		// Empty line
		grabMap(currentLine, "seed-to-soil map:", i, &seedToSoilMap)
		grabMap(currentLine, "soil-to-fertilizer map:", i, &soilToFertilizerMap)
		grabMap(currentLine, "fertilizer-to-water map:", i, &fertilizerToWaterMap)
		grabMap(currentLine, "water-to-light map:", i, &waterToLightMap)
		grabMap(currentLine, "light-to-temperature map:", i, &lightToTemperatureMap)
		grabMap(currentLine, "temperature-to-humidity map:", i, &temperatureToHumidityMap)
		grabMap(currentLine, "humidity-to-location map:", i, &humidityToLocationMap)
	}

	for i := 0; i < len(seeds); i++ {
		seedStart := seeds[i].rangeStart
		startPlusRange := seedStart + seeds[i].rangeLength
		println("Seed start: ", seedStart)
		for i := seedStart; i < startPlusRange; i++ {
			soil := getNextPosition(i, seedToSoilMap)
			fertilizer := getNextPosition(soil, soilToFertilizerMap)
			water := getNextPosition(fertilizer, fertilizerToWaterMap)
			light := getNextPosition(water, waterToLightMap)
			temp := getNextPosition(light, lightToTemperatureMap)
			humidity := getNextPosition(temp, temperatureToHumidityMap)
			location := getNextPosition(humidity, humidityToLocationMap)

			//almanac = append(almanac, almanacEntry{
			//	seed: i, soil: soil, fertilizer: fertilizer, water: water,
			//	light: light, temperature: temp, humidity: humidity, location: location,
			//})
			//print("Added entry to almanac: Seed: ", seedStart, " Soil: ", soil, " Fertilizer: ", fertilizer, " Water: ", water,
			//	" Light: ", light, "\n Temperature: ", temp, " Humidity: ", humidity, " Location: ", location, "\n")
			if location < lowestLocation {
				lowestLocation = location
			}
		}
		println("Done with seed number ", i)
	}

	println("Lowest location: ", lowestLocation)
}

func getNextPosition(sourceIdx int, mapToSearch []gardenMap) int {
	for i := 0; i < len(mapToSearch); i++ {
		entry := mapToSearch[i]
		if entry.sourceRangeStart <= sourceIdx && entry.sourceRangeStart+entry.rangeLength > sourceIdx { // inclusive or not???
			return entry.destinationRangeStart + (sourceIdx - entry.sourceRangeStart)
		}
	}
	return sourceIdx
}

func grabMap(currentLine, searchString string, i int, gardenMapToGrab *[]gardenMap) {
	if currentLine == searchString {
		i++
		for i < len(allInputLines) && allInputLines[i] != "" {
			addToMap(allInputLines[i], *&gardenMapToGrab)
			i++
		}

		//println("Added "+searchString+" with length: ", len(*gardenMapToGrab))
	}
}

func addToMap(currentLine string, gardenMapToAddTo *[]gardenMap) {
	mapSlice := strings.Split(currentLine, " ")
	mapToAppend := gardenMap{}
	mapToAppend.destinationRangeStart, _ = strconv.Atoi(mapSlice[0])
	mapToAppend.sourceRangeStart, _ = strconv.Atoi(mapSlice[1])
	mapToAppend.rangeLength, _ = strconv.Atoi(mapSlice[2])
	*gardenMapToAddTo = append(*gardenMapToAddTo, mapToAppend)
}

func grabSeeds(i int, currentLine string) {
	if i == 0 {

		seedsFullString := strings.Split(currentLine, ": ")
		seedsString := strings.Split(seedsFullString[1], " ")
		for j := 0; j < len(seedsString)-1; j++ {
			seedStart, _ := strconv.Atoi(seedsString[j])
			j++
			seedRange, _ := strconv.Atoi(seedsString[j])
			seed := seedEntry{rangeStart: seedStart, rangeLength: seedRange}
			seeds = append(seeds, seed)
			//println("Seed: ", seedStart, " Range: ", seedRange)
		}

	}
}

func getInputLines(scanner *bufio.Scanner) {
	for scanner.Scan() {
		allInputLines = append(allInputLines, scanner.Text())
	}
}
