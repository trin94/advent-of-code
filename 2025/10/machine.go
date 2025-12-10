package main

import (
	"iter"
	"math"
	"slices"
	"strconv"
	"strings"
	"trin94/aoc/2025/utils"
)

type LightDiagram struct {
	sequenceOn []bool
}

func NewLightDiagram(battery string) LightDiagram {
	sequenceOn := make([]bool, len(battery))
	for idx, value := range battery {
		if value == '.' {
			sequenceOn[idx] = false
		} else {
			sequenceOn[idx] = true
		}
	}
	return LightDiagram{sequenceOn}
}

func (ld *LightDiagram) Length() int {
	return len(ld.sequenceOn)
}

type ButtonWiringSemantic struct {
	pressedOn []int
}

func NewButtonWiringSemantic(semantic string) ButtonWiringSemantic {
	result := parseNumbers(semantic)
	return ButtonWiringSemantic{result}
}

func parseNumbers(semantic string) []int {
	split := strings.Split(semantic, ",")
	result := make([]int, len(split))
	for idx, value := range split {
		result[idx], _ = strconv.Atoi(value)
	}
	return result
}

type JoltageRequirement struct {
	numbers []int
}

func NewJoltageRequirement(numbers string) JoltageRequirement {
	result := parseNumbers(numbers)
	return JoltageRequirement{result}
}

type Machine struct {
	lightDiagram        LightDiagram
	buttonSemantics     []ButtonWiringSemantic
	joltageRequirements JoltageRequirement
}

func NewMachine(line string) Machine {
	split := strings.Split(line, " ")

	batteryString := split[0]
	batteryString = batteryString[1 : len(batteryString)-1]

	joltageString := split[len(split)-1]
	joltageString = joltageString[1 : len(joltageString)-1]

	buttonSemantics := make([]ButtonWiringSemantic, 0)
	buttonString := split[1 : len(split)-1]
	for _, button := range buttonString {
		buttonSemantics = append(buttonSemantics, NewButtonWiringSemantic(button[1:len(button)-1]))
	}

	return Machine{
		lightDiagram:        NewLightDiagram(batteryString),
		buttonSemantics:     buttonSemantics,
		joltageRequirements: NewJoltageRequirement(joltageString),
	}
}

func (m Machine) CalculateFewestButtonPressesToTurnOn() int {
	required := m.lightDiagram.sequenceOn

	for combo := range combinations(m.buttonSemantics) {
		battery := make([]bool, len(required))
		for _, wire := range combo {
			battery = wireUpButton(battery, wire.pressedOn)
		}
		if slices.Equal(required, battery) {
			return len(combo)
		}
	}
	return -1
}

func wireUpButton(battery []bool, toggle []int) []bool {
	result := utils.Clone(battery)
	for _, index := range toggle {
		result[index] = !battery[index]
	}
	return result
}

func combinations[T any](items []T) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		n := len(items)

		for k := 1; k <= n; k++ {
			var backtrack func(start int, current []T) bool
			backtrack = func(start int, current []T) bool {
				if len(current) == k {
					combo := make([]T, k)
					copy(combo, current)
					return yield(combo)
				}
				for i := start; i < n; i++ {
					if !backtrack(i+1, append(current, items[i])) {
						return false
					}
				}
				return true
			}
			if !backtrack(0, []T{}) {
				return
			}
		}
	}
}

func decreaseVoltage(energy []int, toggle []int) []int {
	result := utils.Clone(energy)
	for _, index := range toggle {
		result[index] -= 1
	}
	return result
}

// CalculateFewestButtonPressesToReachVoltage was written by Claude Opus 4.5
func (m Machine) CalculateFewestButtonPressesToReachVoltage() int {
	target := m.joltageRequirements.numbers
	n := len(target)
	numButtons := len(m.buttonSemantics)

	// Build effect matrix
	matrix := make([][]float64, n)
	for i := range matrix {
		matrix[i] = make([]float64, numButtons+1)
		matrix[i][numButtons] = float64(target[i])
	}
	for btnIdx, button := range m.buttonSemantics {
		for _, pos := range button.pressedOn {
			matrix[pos][btnIdx] = 1
		}
	}

	// Gaussian elimination
	pivotCol := make([]int, n)
	for i := range pivotCol {
		pivotCol[i] = -1
	}

	row := 0
	for col := 0; col < numButtons && row < n; col++ {
		// Find pivot
		pivotRow := -1
		for r := row; r < n; r++ {
			if math.Abs(matrix[r][col]) > 1e-9 {
				pivotRow = r
				break
			}
		}
		if pivotRow == -1 {
			continue
		}

		matrix[row], matrix[pivotRow] = matrix[pivotRow], matrix[row]
		pivotCol[row] = col

		// Eliminate column
		for r := 0; r < n; r++ {
			if r != row && math.Abs(matrix[r][col]) > 1e-9 {
				factor := matrix[r][col] / matrix[row][col]
				for k := 0; k <= numButtons; k++ {
					matrix[r][k] -= factor * matrix[row][k]
				}
			}
		}
		row++
	}

	// Identify free variables (columns without pivot)
	isFree := make([]bool, numButtons)
	for col := 0; col < numButtons; col++ {
		isFree[col] = true
	}
	for _, col := range pivotCol {
		if col >= 0 {
			isFree[col] = false
		}
	}

	var freeVars []int
	for col := 0; col < numButtons; col++ {
		if isFree[col] {
			freeVars = append(freeVars, col)
		}
	}

	maxVal := 0
	for _, v := range target {
		if v > maxVal {
			maxVal = v
		}
	}

	best := -1

	var search func(idx int, freeValues []int)
	search = func(idx int, freeValues []int) {
		if idx < len(freeVars) {
			for p := 0; p <= maxVal; p++ {
				search(idx+1, append(freeValues, p))
			}
			return
		}

		// Solve for pivot variables given free values
		presses := make([]int, numButtons)
		for i, col := range freeVars {
			presses[col] = freeValues[i]
		}

		valid := true
		for r := n - 1; r >= 0; r-- {
			col := pivotCol[r]
			if col < 0 {
				continue
			}

			rhs := matrix[r][numButtons]
			for c := 0; c < numButtons; c++ {
				if c != col {
					rhs -= matrix[r][c] * float64(presses[c])
				}
			}
			val := rhs / matrix[r][col]
			rounded := int(math.Round(val))
			if rounded < 0 || math.Abs(val-float64(rounded)) > 1e-6 {
				valid = false
				break
			}
			presses[col] = rounded
		}

		if !valid {
			return
		}

		total := 0
		for _, p := range presses {
			total += p
		}
		if best == -1 || total < best {
			best = total
		}
	}

	search(0, []int{})
	return best
}
