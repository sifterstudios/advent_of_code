package main

import (
	"bufio"
	"os"
	"slices"
)

var (
	visitedPositions                           = make(map[Position]struct{})
	positionsGuardWasBlocked                   = make(map[Position][]Direction)
	allPositions                               = []Position{}
	guardPositionStartingSymbol                = '^'
	obstacleSymbol                             = '#'
	currentGuardPosition                       = Position{}
	startingGuardPosition                      = Position{}
	currentGuardDirection                      = UP
	maxLineNumber                              = -1
	maxColumnNumber                            = 0
	sumOfLoopPossibilitiesWithOneExtraObstacle = 0
)

type Position struct {
	x, y                 int
	obstacle             bool
	currentObstacleCheck bool
}
type Direction int

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

func (d Direction) String() string {
	switch d {
	case UP:
		return "UP"
	case RIGHT:
		return "RIGHT"
	case DOWN:
		return "DOWN"
	case LEFT:
		return "LEFT"
	default:
		return "Unknown"
	}
}

func main() {
	file, err := os.Open("2024/day6/part2/part2_input")
	if err != nil {
		println("Couldn't open file")
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		maxLineNumber++
		for x, char := range line {
			if x > maxColumnNumber {
				maxColumnNumber = x
			}

			guard := char == guardPositionStartingSymbol
			if guard {
				currentGuardPosition = Position{x, maxLineNumber, false, false}
				startingGuardPosition = currentGuardPosition
				allPositions = append(allPositions, currentGuardPosition)
				continue
			}

			obstacle := char == obstacleSymbol

			allPositions = append(allPositions, Position{x, maxLineNumber, obstacle, false})
		}
	}

	println("Number of Positions: ", len(allPositions))
	println("Found guard at position: ", currentGuardPosition.x, currentGuardPosition.y)
	println("Max Column Number: ", maxColumnNumber)
	println("Max Line Number: ", maxLineNumber)

	// Add starting position to visited positions
	visitedPositions[currentGuardPosition] = struct{}{}
	getGuardPath()
	// println("Visited positions: ", len(visitedPositions))

	analyzeLoopCreationIfOneObstacleIsAdded()
	println("Sum of loop possibilities with one extra obstacle: ", sumOfLoopPossibilitiesWithOneExtraObstacle)
}

func analyzeLoopCreationIfOneObstacleIsAdded() {
	for i, pos := range allPositions {
		// Skip positions that are already obstacles or the guard's starting position
		if pos.obstacle || (pos.x == startingGuardPosition.x && pos.y == startingGuardPosition.y) {
			continue
		}

		// Temporarily set the position as an obstacle
		allPositions[i].obstacle = true
		// println("Placing temporary obstacle at:", pos.x, pos.y)

		// Simulate the guard's path with the new obstacle
		resetStats()
		getGuardPath()

		// Revert the position back to its original state
		allPositions[i].obstacle = false
	}
}

func resetStats() {
	currentGuardPosition = startingGuardPosition
	currentGuardDirection = UP
	visitedPositions = make(map[Position]struct{})
	positionsGuardWasBlocked = make(map[Position][]Direction)
}

func getGuardPath() {
	if currentPositionIsOutOfBounds() {
		// println("Current Guard Position is out of bounds", currentGuardPosition.x, currentGuardPosition.y)
		return
	}

	if slices.Contains(positionsGuardWasBlocked[currentGuardPosition], currentGuardDirection) {
		println("Loop detected at position", currentGuardPosition.x, currentGuardPosition.y, "in direction", currentGuardDirection)
		sumOfLoopPossibilitiesWithOneExtraObstacle++
		return
	}

	visitedPositions[currentGuardPosition] = struct{}{}
	for isGuardBlocked() {

		positionsGuardWasBlocked[currentGuardPosition] = append(positionsGuardWasBlocked[currentGuardPosition], currentGuardDirection)

		changeCurrentDirection()
	}

	moveGuardInCurrentDirection()
	getGuardPath()
}

func currentPositionIsOutOfBounds() bool {
	return currentGuardPosition.x < 0 || currentGuardPosition.x > maxColumnNumber || currentGuardPosition.y < 0 || currentGuardPosition.y > maxLineNumber
}

func isGuardBlocked() bool {
	nextPosition := getNextPositionInCurrentDirection()

	if nextPosition.obstacle && nextPosition.currentObstacleCheck {
		// println("Guard is blocked at position with currentObstacleCheck: ", nextPosition.x, nextPosition.y)
	}

	return nextPosition.obstacle
}

func getNextPositionInCurrentDirection() Position {
	newPosition := Position{currentGuardPosition.x, currentGuardPosition.y, false, false}
	switch currentGuardDirection {
	case UP:
		newPosition.y--
	case RIGHT:
		newPosition.x++
	case DOWN:
		newPosition.y++
	case LEFT:
		newPosition.x--
	}
	newPosition.obstacle, newPosition.currentObstacleCheck = getObstacleStatus(newPosition.x, newPosition.y)
	return newPosition
}

func getObstacleStatus(x, y int) (bool, bool) {
	for _, pos := range allPositions {
		if pos.x == x && pos.y == y {
			return pos.obstacle, pos.currentObstacleCheck
		}
	}
	return false, false
}

func moveGuardInCurrentDirection() {
	nextPosition := getNextPositionInCurrentDirection()
	// println("Moving guard in direction", currentGuardDirection, "from position", currentGuardPosition.x, currentGuardPosition.y, "to", nextPosition.x, nextPosition.y)
	currentGuardPosition = nextPosition
}

func changeCurrentDirection() {
	currentGuardDirection++
	if currentGuardDirection == 4 {
		currentGuardDirection = UP
	}
}
