package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMove1(t *testing.T) {
	// given
	linesBefore := []string{
		"#######",
		"#@....#",
		"#######",
	}
	grid := NewGrid(clone(linesBefore))
	robot := grid.robotPosition()

	// when
	robotAfter := grid.move(robot, '<')

	// then
	linesAfter := []string{
		"#######",
		"#@....#",
		"#######",
	}
	gridAfter := NewGrid(linesAfter)
	assert.Equal(t, grid.lines, gridAfter.lines)
	assert.Equal(t, gridAfter.robotPosition(), robotAfter)
}

func TestMove2(t *testing.T) {
	// given
	linesBefore := []string{
		"#######",
		"#.@...#",
		"#######",
	}
	grid := NewGrid(clone(linesBefore))
	robot := grid.robotPosition()

	// when
	robotAfter := grid.move(robot, '<')

	// then
	linesAfter := []string{
		"#######",
		"#@....#",
		"#######",
	}
	gridAfter := NewGrid(linesAfter)
	assert.Equal(t, grid.lines, gridAfter.lines)
	assert.Equal(t, gridAfter.robotPosition(), robotAfter)
}

func TestMove3(t *testing.T) {
	// given
	linesBefore := []string{
		"#######",
		"#.@O..#",
		"#######",
	}
	grid := NewGrid(clone(linesBefore))
	robot := grid.robotPosition()

	// when
	robotAfter := grid.move(robot, '>')

	// then
	linesAfter := []string{
		"#######",
		"#..@O.#",
		"#######",
	}
	gridAfter := NewGrid(linesAfter)
	assert.Equal(t, grid.lines, gridAfter.lines)
	assert.Equal(t, gridAfter.robotPosition(), robotAfter)
}

func TestMove4(t *testing.T) {
	// given
	linesBefore := []string{
		"#######",
		"#.@OO.#",
		"#######",
	}
	grid := NewGrid(clone(linesBefore))
	robot := grid.robotPosition()

	// when
	robotAfter := grid.move(robot, '>')

	// then
	linesAfter := []string{
		"#######",
		"#..@OO#",
		"#######",
	}
	gridAfter := NewGrid(linesAfter)
	assert.Equal(t, grid.lines, gridAfter.lines)
	assert.Equal(t, gridAfter.robotPosition(), robotAfter)
}

func TestMove5(t *testing.T) {
	// given
	linesBefore := []string{
		"#######",
		"#..@OO#",
		"#######",
	}
	grid := NewGrid(clone(linesBefore))
	robot := grid.robotPosition()

	// when
	robotAfter := grid.move(robot, '>')

	// then
	linesAfter := []string{
		"#######",
		"#..@OO#",
		"#######",
	}
	gridAfter := NewGrid(linesAfter)
	assert.Equal(t, grid.lines, gridAfter.lines)
	assert.Equal(t, gridAfter.robotPosition(), robotAfter)
}

func TestSample1(t *testing.T) {
	input := "sample.1.txt"
	expected := 2028
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestSample2(t *testing.T) {
	input := "sample.2.txt"
	expected := 10092
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := "input.txt"
	expected := 1451928
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestDetermineCoordinatesToMove1(t *testing.T) {
	linesBefore := []string{
		"##############",
		"##......##..##",
		"##........@.##",
		"##...[][]...##",
		"##....[]....##",
		"##..........##",
		"##############",
	}
	grid := NewGrid(linesBefore)
	robot := grid.robotPosition()

	actualCoordinates, canMove := grid.determineCoordinatesToMove(robot, 1)
	expectedCoordinates := []Coordinate{
		{2, 10},
	}

	assert.Equal(t, expectedCoordinates, actualCoordinates)
	assert.Equal(t, true, canMove)
}

func TestDetermineCoordinatesToMove2(t *testing.T) {
	linesBefore := []string{
		"##############",
		"##......##..##",
		"##......@...##",
		"##...[][]...##",
		"##....[]....##",
		"##..........##",
		"##############",
	}
	grid := NewGrid(linesBefore)
	robot := grid.robotPosition()

	actualCoordinates, canMove := grid.determineCoordinatesToMove(robot, 1)
	expectedCoordinates := []Coordinate{
		{2, 8},
		{3, 8},
		{3, 7},
		{4, 7},
		{4, 6},
	}

	assert.Equal(t, expectedCoordinates, actualCoordinates)
	assert.Equal(t, true, canMove)
}

func TestDetermineCoordinatesToMove3(t *testing.T) {
	linesBefore := []string{
		"##############",
		"##......##..##",
		"##.....@....##",
		"##...[][]...##",
		"##....[]....##",
		"##..........##",
		"##############",
	}
	grid := NewGrid(linesBefore)
	robot := grid.robotPosition()

	actualCoordinates, canMove := grid.determineCoordinatesToMove(robot, 1)
	expectedCoordinates := []Coordinate{
		{2, 7},
		{3, 7},
		{3, 8},
		{4, 7},
		{4, 6},
	}

	assert.Equal(t, expectedCoordinates, actualCoordinates)
	assert.Equal(t, true, canMove)
}

func TestDetermineCoordinatesToMove4(t *testing.T) {
	linesBefore := []string{
		"##############",
		"##......##..##",
		"##....@.....##",
		"##...[][]...##",
		"##....[]....##",
		"##..........##",
		"##############",
	}
	grid := NewGrid(linesBefore)
	robot := grid.robotPosition()

	actualCoordinates, canMove := grid.determineCoordinatesToMove(robot, 1)
	expectedCoordinates := []Coordinate{
		{2, 6},
		{3, 6},
		{3, 5},
		{4, 6},
		{4, 7},
	}

	assert.Equal(t, expectedCoordinates, actualCoordinates)
	assert.Equal(t, true, canMove)
}

func TestDetermineCoordinatesToMove5(t *testing.T) {
	linesBefore := []string{
		"##############",
		"##..........##",
		"##.[][][][].##",
		"##..[][][]..##",
		"##...[][]...##",
		"##....[]....##",
		"##....@.....##",
		"##############",
	}
	grid := NewGrid(linesBefore)
	robot := grid.robotPosition()

	actualCoordinates, canMove := grid.determineCoordinatesToMove(robot, -1)
	expectedCoordinates := []Coordinate{
		{6, 6},
		{5, 6},
		{5, 7},
		{4, 6},
		{4, 5},
		{4, 7},
		{4, 8},
		{3, 6},
		{3, 7},
		{3, 5},
		{3, 4},
		{3, 8},
		{3, 9},
		{2, 6},
		{2, 5},
		{2, 7},
		{2, 8},
		{2, 4},
		{2, 3},
		{2, 9},
		{2, 10},
	}

	assert.Equal(t, expectedCoordinates, actualCoordinates)
	assert.Equal(t, true, canMove)
}

func TestDetermineCoordinatesToMove6(t *testing.T) {
	linesBefore := []string{
		"##############",
		"##......#...##",
		"##.[][][][].##",
		"##..[][][]..##",
		"##...[][]...##",
		"##....[]....##",
		"##....@.....##",
		"##############",
	}
	grid := NewGrid(linesBefore)
	robot := grid.robotPosition()

	_, canMove := grid.determineCoordinatesToMove(robot, -1)

	assert.Equal(t, false, canMove)
}

func TestDetermineCoordinatesToMove7(t *testing.T) {
	linesBefore := []string{
		"##############",
		"##......#...##",
		"##.[][][][].##",
		"##..[][][]..##",
		"##...[][]...##",
		"##...@[]....##",
		"##..........##",
		"##############",
	}
	grid := NewGrid(linesBefore)
	robot := grid.robotPosition()

	_, canMove := grid.determineCoordinatesToMove(robot, -1)

	assert.Equal(t, false, canMove)
}

func TestDetermineCoordinatesToMove8(t *testing.T) {
	linesBefore := []string{
		"##############",
		"##......#...##",
		"##.[][][][].##",
		"##..[][][]..##",
		"##..@[][]...##",
		"##....[]....##",
		"##..........##",
		"##############",
	}
	grid := NewGrid(linesBefore)
	robot := grid.robotPosition()

	actualCoordinates, canMove := grid.determineCoordinatesToMove(robot, -1)
	expectedCoordinates := []Coordinate{
		{4, 4},
		{3, 4},
		{3, 5},
		{2, 4},
		{2, 3},
		{2, 5},
		{2, 6},
	}

	assert.Equal(t, expectedCoordinates, actualCoordinates)
	assert.Equal(t, true, canMove)
}

func TestMoveScaled1(t *testing.T) {
	// given
	linesBefore := []string{
		"##############",
		"##......##..##",
		"##..........##",
		"##....[][]@.##",
		"##....[]....##",
		"##..........##",
		"##############",
	}
	grid := NewGrid(clone(linesBefore))
	robot := grid.robotPosition()

	// when
	robotAfter := grid.moveScaled(robot, '<')

	// then
	linesAfter := []string{
		"##############",
		"##......##..##",
		"##..........##",
		"##...[][]@..##",
		"##....[]....##",
		"##..........##",
		"##############",
	}
	gridAfter := NewGrid(linesAfter)
	assert.Equal(t, grid.lines, gridAfter.lines)
	assert.Equal(t, gridAfter.robotPosition(), robotAfter)
}

func TestMoveScaled2(t *testing.T) {
	// given
	linesBefore := []string{
		"##############",
		"##......##..##",
		"##..........##",
		"##[][]....@.##",
		"##....[]....##",
		"##..........##",
		"##############",
	}
	grid := NewGrid(clone(linesBefore))
	robot := grid.robotPosition()

	// when
	robotAfter := grid.moveScaled(robot, '<')

	// then
	linesAfter := []string{
		"##############",
		"##......##..##",
		"##..........##",
		"##[][]...@..##",
		"##....[]....##",
		"##..........##",
		"##############",
	}
	gridAfter := NewGrid(linesAfter)
	assert.Equal(t, grid.lines, gridAfter.lines)
	assert.Equal(t, gridAfter.robotPosition(), robotAfter)
}

func TestMoveScaled3(t *testing.T) {
	// given
	linesBefore := []string{
		"##############",
		"##......##..##",
		"##..........##",
		"##[][]@.....##",
		"##....[]....##",
		"##..........##",
		"##############",
	}
	grid := NewGrid(clone(linesBefore))
	robot := grid.robotPosition()

	// when
	robotAfter := grid.moveScaled(robot, '<')

	// then
	linesAfter := []string{
		"##############",
		"##......##..##",
		"##..........##",
		"##[][]@.....##",
		"##....[]....##",
		"##..........##",
		"##############",
	}
	gridAfter := NewGrid(linesAfter)
	assert.Equal(t, grid.lines, gridAfter.lines)
	assert.Equal(t, gridAfter.robotPosition(), robotAfter)
}

func TestMoveScaled4(t *testing.T) {
	// given
	linesBefore := []string{
		"##############",
		"##......##..##",
		"##........@.##",
		"##...[][]...##",
		"##....[]....##",
		"##..........##",
		"##############",
	}
	grid := NewGrid(clone(linesBefore))
	robot := grid.robotPosition()

	// when
	robotAfter := grid.moveScaled(robot, 'v')

	// then
	linesAfter := []string{
		"##############",
		"##......##..##",
		"##..........##",
		"##...[][].@.##",
		"##....[]....##",
		"##..........##",
		"##############",
	}
	gridAfter := NewGrid(linesAfter)
	assert.Equal(t, grid.lines, gridAfter.lines)
	assert.Equal(t, gridAfter.robotPosition(), robotAfter)
}

func TestMoveScaled5(t *testing.T) {
	// given
	linesBefore := []string{
		"##############",
		"##......##..##",
		"##......@...##",
		"##...[][]...##",
		"##....[]....##",
		"##..........##",
		"##############",
	}
	grid := NewGrid(linesBefore)
	robot := grid.robotPosition()

	// when
	robotAfter := grid.moveScaled(robot, 'v')

	// then
	linesAfter := []string{
		"##############",
		"##......##..##",
		"##..........##",
		"##...[].@...##",
		"##.....[]...##",
		"##....[]....##",
		"##############",
	}
	gridAfter := NewGrid(linesAfter)
	assert.Equal(t, grid.lines, gridAfter.lines)
	assert.Equal(t, gridAfter.robotPosition(), robotAfter)
}

func TestMoveScaled6(t *testing.T) {
	// given
	linesBefore := []string{
		"##############",
		"##......##..##",
		"##......@...##",
		"##...[][]...##",
		"##..........##",
		"##..........##",
		"##############",
	}
	grid := NewGrid(linesBefore)
	robot := grid.robotPosition()

	// when
	robotAfter := grid.moveScaled(robot, 'v')

	// then
	linesAfter := []string{
		"##############",
		"##......##..##",
		"##..........##",
		"##...[].@...##",
		"##.....[]...##",
		"##..........##",
		"##############",
	}
	gridAfter := NewGrid(linesAfter)
	assert.Equal(t, grid.lines, gridAfter.lines)
	assert.Equal(t, gridAfter.robotPosition(), robotAfter)
}

func TestMoveScaled7(t *testing.T) {
	// given
	linesBefore := []string{
		"##############",
		"##......##..##",
		"##......@...##",
		"##.....[]...##",
		"##.....[]...##",
		"##..........##",
		"##############",
	}
	grid := NewGrid(linesBefore)
	robot := grid.robotPosition()

	// when
	robotAfter := grid.moveScaled(robot, 'v')

	// then
	linesAfter := []string{
		"##############",
		"##......##..##",
		"##..........##",
		"##......@...##",
		"##.....[]...##",
		"##.....[]...##",
		"##############",
	}
	gridAfter := NewGrid(linesAfter)
	assert.Equal(t, grid.lines, gridAfter.lines)
	assert.Equal(t, gridAfter.robotPosition(), robotAfter)
}

func TestMoveScaled8(t *testing.T) {
	// given
	linesBefore := []string{
		"##############",
		"##......##..##",
		"##..........##",
		"##.....[]...##",
		"##.....[]...##",
		"##......@...##",
		"##############",
	}
	grid := NewGrid(linesBefore)
	robot := grid.robotPosition()

	// when
	robotAfter := grid.moveScaled(robot, '^')

	// then
	linesAfter := []string{
		"##############",
		"##......##..##",
		"##.....[]...##",
		"##.....[]...##",
		"##......@...##",
		"##..........##",
		"##############",
	}
	gridAfter := NewGrid(linesAfter)
	assert.Equal(t, grid.lines, gridAfter.lines)
	assert.Equal(t, gridAfter.robotPosition(), robotAfter)
}

func TestMoveScaled9(t *testing.T) {
	// given
	linesBefore := []string{
		"##############",
		"##......##..##",
		"##..........##",
		"##.....[]...##",
		"##.....[]...##",
		"##.....@....##",
		"##############",
	}
	grid := NewGrid(linesBefore)
	robot := grid.robotPosition()

	// when
	robotAfter := grid.moveScaled(robot, '^')

	// then
	linesAfter := []string{
		"##############",
		"##......##..##",
		"##.....[]...##",
		"##.....[]...##",
		"##.....@....##",
		"##..........##",
		"##############",
	}
	gridAfter := NewGrid(linesAfter)
	assert.Equal(t, grid.lines, gridAfter.lines)
	assert.Equal(t, gridAfter.robotPosition(), robotAfter)
}

func TestSample3(t *testing.T) {
	input := "sample.2.txt"
	expected := 9021
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestPart2Puzzle(t *testing.T) {
	input := "input.txt"
	expected := 1462788
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func clone[E comparable](slice []E) []E {
	into := make([]E, len(slice))
	copy(into, slice)
	return into
}
