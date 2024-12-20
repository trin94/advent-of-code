package main

import (
	"math"
	"trin94/aoc/2024/inputs"
	"trin94/aoc/2024/utils"
)

type Shortcut struct {
	a utils.Coordinate
	b utils.Coordinate
}

func solvePuzzle1(path string) (int, map[int]int) {
	lines := inputs.ReadLinesFrom(path)

	grid := utils.NewGrid(lines)
	start := grid.FindCoordinateWithChar('S')
	end := grid.FindCoordinateWithChar('E')

	_, distances := race(start, end, grid)
	walls := findShortcutWalls(grid)

	savings := determineTimeSaves(walls, distances)

	result := 0

	for saving, amount := range savings {
		if saving >= 100 {
			result += amount
		}
	}

	return result, savings
}

func solvePuzzle2(path string) (int, map[int]int) {
	lines := inputs.ReadLinesFrom(path)

	grid := utils.NewGrid(lines)
	start := grid.FindCoordinateWithChar('S')
	end := grid.FindCoordinateWithChar('E')

	raceTrack, distances := race(start, end, grid)

	savings := determineTimeSavesWithinSeconds(20, raceTrack, grid, distances)

	result := 0

	for saving, amount := range savings {
		if saving >= 100 {
			result += amount
		}
	}

	return result, savings

}

func race(start, end utils.Coordinate, grid utils.Grid) ([]utils.Coordinate, map[utils.Coordinate]int) {

	directions := []utils.Coordinate{
		{Row: -1, Col: 0}, // north
		{Row: 0, Col: 1},  // east
		{Row: 1, Col: 0},  // south
		{Row: 0, Col: -1}, // west
	}

	path := make([]utils.Coordinate, 0)
	distances := make(map[utils.Coordinate]int)

	for distance, previous, next := 0, start, start; ; distance++ {

		path = append(path, next)
		distances[next] = distance

		for _, direction := range directions {
			adjacent := utils.Coordinate{Row: next.Row + direction.Row, Col: next.Col + direction.Col}
			if adjacent == previous {
				continue
			}
			char := grid.CharAtCoordinate(adjacent)
			if char == '.' || char == 'E' {
				previous = next
				next = adjacent
				break
			}
		}

		if next == end {
			path = append(path, next)
			distances[next] = distance + 1
			return path, distances
		}
	}

}

func findShortcutWalls(grid utils.Grid) []Shortcut {
	walls := make([]Shortcut, 0)

	for r := 1; r < grid.Rows-1; r++ {
		for c := 1; c < grid.Columns-1; c++ {
			if grid.CharAt(r, c) != '#' {
				continue
			}

			charLeft := grid.CharAt(r, c-1)
			charRight := grid.CharAt(r, c+1)
			if isRaceTrack(charLeft) && isRaceTrack(charRight) {
				shortcut := Shortcut{
					a: utils.Coordinate{Row: r, Col: c - 1},
					b: utils.Coordinate{Row: r, Col: c + 1},
				}
				walls = append(walls, shortcut)
			}

			charTop := grid.CharAt(r-1, c)
			charBottom := grid.CharAt(r+1, c)
			if isRaceTrack(charTop) && isRaceTrack(charBottom) {
				shortcut := Shortcut{
					a: utils.Coordinate{Row: r - 1, Col: c},
					b: utils.Coordinate{Row: r + 1, Col: c},
				}
				walls = append(walls, shortcut)
			}
		}
	}

	return walls
}

func isRaceTrack(char rune) bool {
	return char == '.' || char == 'S' || char == 'E'
}

func determineTimeSaves(shortcuts []Shortcut, distances map[utils.Coordinate]int) map[int]int {

	savings := make(map[int]int)

	for _, shortcut := range shortcuts {

		distanceA := float64(distances[shortcut.a])
		distanceB := float64(distances[shortcut.b])

		saving := int(math.Abs(distanceA-distanceB)) - utils.ManhattanDistanceBetween(shortcut.a, shortcut.b)

		value, found := savings[saving]
		if !found {
			savings[saving] = 1
		} else {
			savings[saving] = value + 1
		}
	}

	return savings
}

func determineTimeSavesWithinSeconds(maxSeconds int, raceTrack []utils.Coordinate, grid utils.Grid, distances map[utils.Coordinate]int) map[int]int {

	savings := make(map[int]int)

	for _, track := range raceTrack {

		farthestLeft := utils.MaxInt(1, track.Col-maxSeconds)
		farthestRight := utils.MinInt(track.Col+maxSeconds, grid.Columns)

		farthestTop := utils.MaxInt(1, track.Row-maxSeconds)
		farthestBottom := utils.MinInt(track.Row+maxSeconds, grid.Rows)

		for r := farthestTop; r <= farthestBottom; r++ {
			for c := farthestLeft; c <= farthestRight; c++ {

				// skip if not racetrack
				if !isRaceTrack(grid.CharAt(r, c)) {
					continue
				}

				current := utils.Coordinate{Row: r, Col: c}

				// skip if track is same as current
				if current == track {
					continue
				}

				// only check coordinates that are in front
				if distances[track] > distances[current] {
					continue
				}

				invest := utils.ManhattanDistanceBetween(current, track)

				// reachable within time interval
				if invest > maxSeconds {
					continue
				}

				saving := distances[current] - distances[track] - invest

				value, found := savings[saving]
				if !found {
					savings[saving] = 1
				} else {
					savings[saving] = value + 1
				}

			}
		}

	}

	return savings
}
