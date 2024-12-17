package main

import (
	"regexp"
	"strconv"
	"trin94/aoc/2024/inputs"
)

type Coordinate struct {
	x, y int
}

type Robot struct {
	x, y     int
	velocity Velocity
}

type Velocity struct {
	x, y int
}

func solvePuzzle1(path string, spaceWidth int, spaceHeight int) int {
	lines := inputs.ReadLinesFrom(path)
	robots := parseRobotsFrom(lines)

	northWest, northEast, southWest, southEast := 0, 0, 0, 0

	xHalf := spaceWidth / 2
	yHalf := spaceHeight / 2

	seconds := 100

	for _, robot := range robots {
		x, y := moveRobot(robot, seconds, spaceWidth, spaceHeight)

		var east, north bool

		if x < xHalf {
			east = false
		} else if x > xHalf {
			east = true
		} else {
			continue
		}

		if y < yHalf {
			north = true
		} else if y > yHalf {
			north = false
		} else {
			continue
		}

		if east && north {
			northEast++
		} else if !east && north {
			northWest++
		} else if east {
			southEast++
		} else {
			southWest++
		}

	}

	return northWest * northEast * southWest * southEast
}

func solvePuzzle2(path string, spaceWidth int, spaceHeight int) int {
	lines := inputs.ReadLinesFrom(path)
	robots := parseRobotsFrom(lines)

	for times := 1; times < 100000; times++ {
		snapshot := make(map[Coordinate]struct{})

		for i, robot := range robots {
			x, y := moveRobot(robot, 1, spaceWidth, spaceHeight)
			robots[i] = Robot{x, y, robot.velocity}
			coordinate := Coordinate{x, y}
			snapshot[coordinate] = struct{}{}
		}

		for coordinate := range snapshot {
			longRowFound := checkForRobotsInRow(coordinate, snapshot, spaceWidth/4)

			if longRowFound {
				return times
			}
		}

	}

	return -1
}

func parseRobotsFrom(lines []string) []Robot {
	numberPattern, _ := regexp.Compile("-?\\d+")
	result := make([]Robot, len(lines))
	for i, line := range lines {
		numbers := numberPattern.FindAllString(line, -1)
		result[i] = Robot{
			x: toInt(numbers[0]),
			y: toInt(numbers[1]),
			velocity: Velocity{
				x: toInt(numbers[2]),
				y: toInt(numbers[3]),
			},
		}
	}
	return result
}

func toInt(s string) int {
	nr, _ := strconv.Atoi(s)
	return nr
}

func moveRobot(robot Robot, steps int, spaceWidth int, spaceHeight int) (int, int) {
	x := (robot.x + steps*robot.velocity.x) % spaceWidth
	y := (robot.y + steps*robot.velocity.y) % spaceHeight
	if x < 0 {
		x += spaceWidth
	}
	if y < 0 {
		y += spaceHeight
	}
	return x, y
}

func checkForRobotsInRow(coordinate Coordinate, snapshot map[Coordinate]struct{}, minimalRowLength int) bool {
	neighbors := 0

	for l := 1; l <= minimalRowLength; l++ {
		neighborCoordinate := Coordinate{coordinate.x - l, coordinate.y}
		_, found := snapshot[neighborCoordinate]
		if found {
			neighbors++
		} else {
			break
		}
	}

	for r := 1; r <= minimalRowLength; r++ {
		neighborCoordinate := Coordinate{coordinate.x + r, coordinate.y}
		_, found := snapshot[neighborCoordinate]
		if found {
			neighbors++
		} else {
			break
		}
	}

	return neighbors >= minimalRowLength
}
