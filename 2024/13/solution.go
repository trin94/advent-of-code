package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Button struct {
	moveX int
	moveY int
}

type Price struct {
	x int
	y int
}

type Machine struct {
	buttonA Button
	buttonB Button
	price   Price
}

func solvePuzzle1(path string) int {
	lines := readLinesFrom(path)
	machines := parseMachines(lines)

	costs := 0

	for _, machine := range machines {
		costs += play(machine)
	}

	return costs
}

func solvePuzzle2(path string) int {
	lines := readLinesFrom(path)
	machines := parseMachines(lines)
	machines = adjustPriceLocations(machines)

	costs := 0

	for _, machine := range machines {
		costs += play(machine)
	}

	return costs
}

func readLinesFrom(path string) []string {
	inputByteStream, _ := os.ReadFile(path)
	inputString := string(inputByteStream)
	inputString = strings.TrimSpace(inputString)
	return strings.Split(inputString, "\n")
}

func parseMachines(lines []string) []Machine {
	result := make([]Machine, len(lines)/4+1)
	resultIdx := 0

	numberPattern, _ := regexp.Compile("X\\+(\\d+), Y\\+(\\d+)")
	pricePattern, _ := regexp.Compile("X=(\\d+), Y=(\\d+)")

	for i := 0; i < len(lines); i = i + 4 {
		submatch := numberPattern.FindStringSubmatch(lines[i])
		buttonA := Button{moveX: toInt(submatch[1]), moveY: toInt(submatch[2])}

		submatch = numberPattern.FindStringSubmatch(lines[i+1])
		buttonB := Button{moveX: toInt(submatch[1]), moveY: toInt(submatch[2])}

		submatch = pricePattern.FindStringSubmatch(lines[i+2])
		price := Price{x: toInt(submatch[1]), y: toInt(submatch[2])}

		result[resultIdx] = Machine{buttonA, buttonB, price}
		resultIdx++
	}

	return result
}

func toInt(s string) int {
	nr, _ := strconv.Atoi(s)
	return nr
}

func play(machine Machine) (cost int) {
	a1 := float64(machine.buttonA.moveX)
	b1 := float64(machine.buttonB.moveX)
	r1 := float64(machine.price.x)

	a2 := float64(machine.buttonA.moveY)
	b2 := float64(machine.buttonB.moveY)
	r2 := float64(machine.price.y)

	a, b, err := solveEquations(a1, b1, r1, a2, b2, r2)

	if err != nil || a != float64(int64(a)) || b != float64(int64(b)) {
		return
	}

	return int(3*a + b)
}

// Function to solve the system of equations using Cramer's Rule
func solveEquations(a1, b1, c1, a2, b2, c2 float64) (float64, float64, error) {
	// Calculate the determinant of the coefficient matrix
	det := determinant(a1, b1, a2, b2)
	if det == 0 {
		return 0, 0, fmt.Errorf("the system has no unique solution (determinant is zero)")
	}

	// Calculate determinants for x and y
	detX := determinant(c1, b1, c2, b2)
	detY := determinant(a1, c1, a2, c2)

	// Use Cramer's rule to compute x and y
	x := detX / det
	y := detY / det

	return x, y, nil
}

// Function to compute the determinant of a 2x2 matrix
func determinant(a, b, c, d float64) float64 {
	return a*d - b*c
}

func adjustPriceLocations(machines []Machine) []Machine {
	result := make([]Machine, len(machines))
	for i, machine := range machines {
		newMachine := Machine{
			buttonA: machine.buttonA,
			buttonB: machine.buttonB,
			price:   Price{x: machine.price.x + 10000000000000, y: machine.price.y + 10000000000000},
		}
		result[i] = newMachine
	}
	return result
}
