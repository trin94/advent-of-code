package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"trin94/aoc/2024/inputs"
)

type Equation struct {
	left, operation, right, result string
}

func solvePuzzle1(path string) int {
	lines := inputs.ReadLinesFrom(path)
	gates, equations := parseGatesAndEquations(lines)

	for {
		finished := true

		for _, equation := range equations {
			left, leftFound := gates[equation.left]
			right, rightFound := gates[equation.right]

			if !leftFound || !rightFound {
				finished = false
				continue
			}

			operation, target := equation.operation, equation.result

			var result bool
			switch operation {
			case "XOR":
				result = left != right
			case "OR":
				result = left || right
			case "AND":
				result = left && right
			default:
				panic("Invalid operation: " + operation)
			}

			if _, found := gates[target]; !found {
				gates[target] = result
			}
		}

		if finished {
			break
		}
	}

	return determineBooleanValue(gates)
}

func parseGatesAndEquations(lines []string) (map[string]bool, []Equation) {
	gates := make(map[string]bool)
	equations := make([]Equation, 0)
	isParsingGates := true
	for _, line := range lines {
		if line == "" {
			isParsingGates = false
			continue
		}
		fields := strings.Fields(line)
		if isParsingGates {
			first := fields[0]
			first = first[:len(first)-1]
			gates[first] = fields[1] == "1"
		} else {
			left := fields[0]
			operation := fields[1]
			right := fields[2]
			result := fields[4]
			equation := Equation{left, operation, right, result}
			equations = append(equations, equation)
		}
	}
	return gates, equations
}

func determineBooleanValue(gates map[string]bool) int {
	zGateCount := 0
	for gate := range gates {
		if strings.HasPrefix(gate, "z") {
			zGateCount++
		}
	}

	outputGates := make([]bool, zGateCount)
	for gate, isSet := range gates {
		if strings.HasPrefix(gate, "z") {
			index, _ := strconv.Atoi(gate[1:])
			outputGates[index] = isSet
		}
	}

	return int(boolSliceToBinary(outputGates))
}

func boolSliceToBinary(boolSlice []bool) uint {
	var binaryNumber uint = 0
	for i, val := range boolSlice {
		if val {
			binaryNumber |= 1 << i // Set the bit at position i to 1
		}
	}
	return binaryNumber
}

func solvePuzzle2(path string) string {
	lines := inputs.ReadLinesFrom(path)
	_, equations := parseGatesAndEquations(lines)

	var swapped []string
	var carry0 string

	// Copied from https://github.com/ayoubzulfiqar/advent-of-code/blob/main/2024/Go/Day24/part_2.go

	for i := 0; i < 45; i++ {
		n := fmt.Sprintf("%02d", i)
		var m1, n1, r1, z1, carry1 string

		// Half adder logic
		m1 = find("x"+n, "y"+n, "XOR", equations)
		n1 = find("x"+n, "y"+n, "AND", equations)

		if carry0 != "" {
			r1 = find(carry0, m1, "AND", equations)
			if r1 == "" {
				m1, n1 = n1, m1
				swapped = append(swapped, m1, n1)
				r1 = find(carry0, m1, "AND", equations)
			}

			z1 = find(carry0, m1, "XOR", equations)

			if strings.HasPrefix(m1, "z") {
				m1, z1 = z1, m1
				swapped = append(swapped, m1, z1)
			}

			if strings.HasPrefix(n1, "z") {
				n1, z1 = z1, n1
				swapped = append(swapped, n1, z1)
			}

			if strings.HasPrefix(r1, "z") {
				r1, z1 = z1, r1
				swapped = append(swapped, r1, z1)
			}

			carry1 = find(r1, n1, "OR", equations)
		}

		if strings.HasPrefix(carry1, "z") && carry1 != "z45" {
			carry1, z1 = z1, carry1
			swapped = append(swapped, carry1, z1)
		}

		if carry0 == "" {
			carry0 = n1
		} else {
			carry0 = carry1
		}
	}

	slices.Sort(swapped)
	return strings.Join(swapped, ",")
}

func find(a, b, operator string, equations []Equation) string {
	for _, function := range equations {
		if function.operation != operator {
			continue
		}
		ab := function.left == a && function.right == b
		ba := function.left == b && function.right == a
		if ab || ba {
			return function.result
		}
	}
	return ""
}
