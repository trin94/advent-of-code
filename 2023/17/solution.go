package main

import (
	"container/heap"
	"strconv"
	"trin94/aoc/2023/inputs"
	"trin94/aoc/2023/utils"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

type Move struct {
	path             []utils.Coordinate
	current          utils.Coordinate
	currentDirection utils.Direction
	directionSteps   int
	rows, cols       int
	heatLoss         int
}

var InvalidMove = Move{
	path: nil,
}

func (m *Move) North() Move {
	if m.currentDirection == utils.South {
		return InvalidMove
	}
	if m.currentDirection == utils.North && m.directionSteps > 2 {
		return InvalidMove
	}
	if m.current.Row <= 0 {
		return InvalidMove
	}
	next := utils.Coordinate{Col: m.current.Col, Row: m.current.Row - 1}
	path := utils.Clone(m.path)
	path = append(path, next)

	steps := 1
	if m.currentDirection == utils.North {
		steps = m.directionSteps + 1
	}

	return Move{
		path:             path,
		current:          next,
		currentDirection: utils.North,
		directionSteps:   steps,
		rows:             m.rows,
		cols:             m.cols,
		heatLoss:         -1,
	}
}

func (m *Move) UltraNorth() Move {
	if m.current.Row <= 0 {
		return InvalidMove
	}
	if m.currentDirection == utils.South {
		return InvalidMove
	}
	if m.currentDirection == utils.North && m.directionSteps >= 10 {
		return InvalidMove
	}
	if m.currentDirection != utils.North && m.directionSteps < 4 {
		return InvalidMove
	}

	next := utils.Coordinate{Col: m.current.Col, Row: m.current.Row - 1}
	path := utils.Clone(m.path)
	path = append(path, next)

	steps := 1
	if m.currentDirection == utils.North {
		steps = m.directionSteps + 1
	}

	return Move{
		path:             path,
		current:          next,
		currentDirection: utils.North,
		directionSteps:   steps,
		rows:             m.rows,
		cols:             m.cols,
		heatLoss:         -1,
	}
}

func (m *Move) East() Move {
	if m.currentDirection == utils.West {
		return InvalidMove
	}
	if m.currentDirection == utils.East && m.directionSteps > 2 {
		return InvalidMove
	}
	if m.current.Col >= m.cols-1 {
		return InvalidMove
	}
	next := utils.Coordinate{Col: m.current.Col + 1, Row: m.current.Row}
	path := utils.Clone(m.path)
	path = append(path, next)

	steps := 1
	if m.currentDirection == utils.East {
		steps = m.directionSteps + 1
	}

	return Move{
		path:             path,
		current:          next,
		currentDirection: utils.East,
		directionSteps:   steps,
		rows:             m.rows,
		cols:             m.cols,
		heatLoss:         -1,
	}
}

func (m *Move) UltraEast() Move {
	if m.current.Col >= m.cols-1 {
		return InvalidMove
	}
	if m.currentDirection == utils.West {
		return InvalidMove
	}
	if m.currentDirection == utils.East && m.directionSteps >= 10 {
		return InvalidMove
	}
	if m.currentDirection != utils.East && m.directionSteps < 4 {
		return InvalidMove
	}

	next := utils.Coordinate{Col: m.current.Col + 1, Row: m.current.Row}
	path := utils.Clone(m.path)
	path = append(path, next)

	steps := 1
	if m.currentDirection == utils.East {
		steps = m.directionSteps + 1
	}

	return Move{
		path:             path,
		current:          next,
		currentDirection: utils.East,
		directionSteps:   steps,
		rows:             m.rows,
		cols:             m.cols,
		heatLoss:         -1,
	}
}

func (m *Move) South() Move {
	if m.currentDirection == utils.North {
		return InvalidMove
	}
	if m.currentDirection == utils.South && m.directionSteps > 2 {
		return InvalidMove
	}
	if m.current.Row >= m.rows-1 {
		return InvalidMove
	}
	next := utils.Coordinate{Col: m.current.Col, Row: m.current.Row + 1}
	path := utils.Clone(m.path)
	path = append(path, next)

	steps := 1
	if m.currentDirection == utils.South {
		steps = m.directionSteps + 1
	}

	return Move{
		path:             path,
		current:          next,
		currentDirection: utils.South,
		directionSteps:   steps,
		rows:             m.rows,
		cols:             m.cols,
		heatLoss:         -1,
	}
}

func (m *Move) UltraSouth() Move {
	if m.current.Row >= m.rows-1 {
		return InvalidMove
	}
	if m.currentDirection == utils.North {
		return InvalidMove
	}
	if m.currentDirection == utils.South && m.directionSteps >= 10 {
		return InvalidMove
	}
	if m.currentDirection != utils.South && m.directionSteps < 4 {
		return InvalidMove
	}

	next := utils.Coordinate{Col: m.current.Col, Row: m.current.Row + 1}
	path := utils.Clone(m.path)
	path = append(path, next)

	steps := 1
	if m.currentDirection == utils.South {
		steps = m.directionSteps + 1
	}

	return Move{
		path:             path,
		current:          next,
		currentDirection: utils.South,
		directionSteps:   steps,
		rows:             m.rows,
		cols:             m.cols,
		heatLoss:         -1,
	}
}

func (m *Move) West() Move {
	if m.currentDirection == utils.East {
		return InvalidMove
	}
	if m.currentDirection == utils.West && m.directionSteps > 2 {
		return InvalidMove
	}
	if m.current.Col <= 0 {
		return InvalidMove
	}
	next := utils.Coordinate{Col: m.current.Col - 1, Row: m.current.Row}
	path := utils.Clone(m.path)
	path = append(path, next)

	steps := 1
	if m.currentDirection == utils.West {
		steps = m.directionSteps + 1
	}

	return Move{
		path:             path,
		current:          next,
		currentDirection: utils.West,
		directionSteps:   steps,
		rows:             m.rows,
		cols:             m.cols,
		heatLoss:         -1,
	}
}

func (m *Move) UltraWest() Move {
	if m.current.Col <= 0 {
		return InvalidMove
	}
	if m.currentDirection == utils.East {
		return InvalidMove
	}
	if m.currentDirection == utils.West && m.directionSteps >= 10 {
		return InvalidMove
	}
	if m.currentDirection != utils.West && m.directionSteps < 4 {
		return InvalidMove
	}

	next := utils.Coordinate{Col: m.current.Col - 1, Row: m.current.Row}
	path := utils.Clone(m.path)
	path = append(path, next)

	steps := 1
	if m.currentDirection == utils.West {
		steps = m.directionSteps + 1
	}

	return Move{
		path:             path,
		current:          next,
		currentDirection: utils.West,
		directionSteps:   steps,
		rows:             m.rows,
		cols:             m.cols,
		heatLoss:         -1,
	}
}

func (m *Move) CalculateHeatLoss(grid utils.Grid) int {
	if m.heatLoss > -1 {
		return m.heatLoss
	}
	sum := 0
	for _, path := range m.path {
		costs := grid.CharAtCoordinate(path)
		heatLoss, _ := strconv.Atoi(string(costs))
		sum += heatLoss
	}
	m.heatLoss = sum
	return sum
}

func (m *Move) IsValid() bool {
	return m.path != nil
}

type Arrival struct {
	coords    utils.Coordinate
	direction utils.Direction
	steps     int
}

type MoveHeap []Move

func (h MoveHeap) Len() int {
	return len(h)
}

func (h MoveHeap) Less(i, j int) bool {
	return h[i].heatLoss < h[j].heatLoss
}

func (h MoveHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MoveHeap) Push(x any) {
	*h = append(*h, x.(Move))
}

func (h *MoveHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func solvePuzzle1(path string) int {
	lines := inputs.ReadLinesFrom(path)
	grid := utils.NewGrid(lines)

	coord := utils.Coordinate{Col: 0, Row: 0}

	costs := make(map[Arrival]int)
	costs[Arrival{
		coords:    coord,
		direction: utils.East,
		steps:     0,
	}] = 0
	costs[Arrival{
		coords:    coord,
		direction: utils.South,
		steps:     0,
	}] = 0

	visitNext := &MoveHeap{}
	heap.Init(visitNext)

	heap.Push(visitNext, Move{
		path:             make([]utils.Coordinate, 0),
		current:          coord,
		currentDirection: utils.East,
		directionSteps:   0,
		rows:             grid.Rows(),
		cols:             grid.Columns(),
		heatLoss:         0,
	})
	heap.Push(visitNext, Move{
		path:             make([]utils.Coordinate, 0),
		current:          coord,
		currentDirection: utils.South,
		directionSteps:   0,
		rows:             grid.Rows(),
		cols:             grid.Columns(),
		heatLoss:         0,
	})

	process := func(move Move) bool {
		arrival := Arrival{
			coords: utils.Coordinate{
				Col: move.current.Col,
				Row: move.current.Row,
			},
			direction: move.currentDirection,
			steps:     move.directionSteps,
		}
		newCosts := move.CalculateHeatLoss(grid)
		currentCosts, existing := costs[arrival]
		if !existing {
			costs[arrival] = newCosts
			return true
		}
		if newCosts < currentCosts {
			costs[arrival] = newCosts
			return true
		}
		return false
	}

	for visitNext.Len() > 0 {
		next := heap.Pop(visitNext).(Move)

		directions := make([]Move, 0)
		directions = append(directions, next.North())
		directions = append(directions, next.East())
		directions = append(directions, next.South())
		directions = append(directions, next.West())

		for _, direction := range directions {
			if direction.IsValid() {
				store := process(direction)
				if store {
					direction.heatLoss = direction.CalculateHeatLoss(grid)
					heap.Push(visitNext, direction)
				}
			}
		}
	}

	goal := utils.Coordinate{Col: grid.Columns() - 1, Row: grid.Rows() - 1}

	minHeatLoss := MaxInt
	for arrival, cost := range costs {
		if arrival.coords.Col == goal.Col && arrival.coords.Row == goal.Row {
			minHeatLoss = min(minHeatLoss, cost)
		}
	}

	return minHeatLoss
}

func solvePuzzle2(path string) int {
	lines := inputs.ReadLinesFrom(path)
	grid := utils.NewGrid(lines)

	coord := utils.Coordinate{Col: 0, Row: 0}

	costs := make(map[Arrival]int)
	costs[Arrival{
		coords:    coord,
		direction: utils.East,
		steps:     0,
	}] = 0
	costs[Arrival{
		coords:    coord,
		direction: utils.South,
		steps:     0,
	}] = 0

	visitNext := &MoveHeap{}
	heap.Init(visitNext)

	heap.Push(visitNext, Move{
		path:             make([]utils.Coordinate, 0),
		current:          coord,
		currentDirection: utils.East,
		directionSteps:   0,
		rows:             grid.Rows(),
		cols:             grid.Columns(),
		heatLoss:         0,
	})
	heap.Push(visitNext, Move{
		path:             make([]utils.Coordinate, 0),
		current:          coord,
		currentDirection: utils.South,
		directionSteps:   0,
		rows:             grid.Rows(),
		cols:             grid.Columns(),
		heatLoss:         0,
	})

	process := func(move Move) bool {
		arrival := Arrival{
			coords: utils.Coordinate{
				Col: move.current.Col,
				Row: move.current.Row,
			},
			direction: move.currentDirection,
			steps:     move.directionSteps,
		}
		newCosts := move.CalculateHeatLoss(grid)
		currentCosts, existing := costs[arrival]
		if !existing {
			costs[arrival] = newCosts
			return true
		}
		if newCosts < currentCosts {
			costs[arrival] = newCosts
			return true
		}
		return false
	}

	for visitNext.Len() > 0 {
		next := heap.Pop(visitNext).(Move)

		directions := make([]Move, 0)
		directions = append(directions, next.UltraNorth())
		directions = append(directions, next.UltraEast())
		directions = append(directions, next.UltraSouth())
		directions = append(directions, next.UltraWest())

		for _, direction := range directions {
			if direction.IsValid() {
				store := process(direction)
				if store {
					direction.heatLoss = direction.CalculateHeatLoss(grid)
					heap.Push(visitNext, direction)
				}
			}
		}
	}

	goal := utils.Coordinate{Col: grid.Columns() - 1, Row: grid.Rows() - 1}

	minHeatLoss := MaxInt
	for arrival, cost := range costs {
		if arrival.coords.Col == goal.Col && arrival.coords.Row == goal.Row && arrival.steps >= 4 {
			minHeatLoss = min(minHeatLoss, cost)
		}
	}

	return minHeatLoss
}
