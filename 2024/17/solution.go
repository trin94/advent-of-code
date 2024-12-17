package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"trin94/aoc/2024/inputs"
)

type Instruction struct {
	opcode  int
	operand int
}

type Computer struct {
	a                  int
	b                  int
	c                  int
	instructionPointer int
	output             []string
	outputInt          []int
}

func NewComputer(a, b, c int) Computer {
	return Computer{
		a:         a,
		b:         b,
		c:         c,
		output:    make([]string, 0),
		outputInt: make([]int, 0),
	}
}

func (c *Computer) Write(value int) {
	c.output = append(c.output, strconv.Itoa(value))
	c.outputInt = append(c.outputInt, value)
}

// adv 0, combo operand
func (c *Computer) adv(inst Instruction) {
	numerator := float64(c.a)
	denominator := math.Pow(2, float64(c.comboOperand(inst.operand)))
	c.a = int(numerator / denominator)
}

// bxl 1, literal operand
func (c *Computer) bxl(inst Instruction) {
	first := c.b
	second := inst.operand
	c.b = first ^ second
}

// bst 2, combo operand
func (c *Computer) bst(inst Instruction) {
	c.b = c.comboOperand(inst.operand) % 8
}

// jnz 3, literal operand
func (c *Computer) jnz(inst Instruction) bool {
	if c.a == 0 {
		return false
	}
	c.instructionPointer = inst.operand
	return true
}

// bxc 4, ignores operand
func (c *Computer) bxc() {
	c.b = c.b ^ c.c
}

// out 5, combo operand
func (c *Computer) out(inst Instruction) {
	value := c.comboOperand(inst.operand) % 8
	c.Write(value)
}

// bdv 6, combo operand
func (c *Computer) bdv(inst Instruction) {
	numerator := float64(c.a)
	denominator := math.Pow(2, float64(c.comboOperand(inst.operand)))
	result := int(numerator / denominator)
	c.b = result
}

// cdv 7, combo operand
func (c *Computer) cdv(inst Instruction) {
	numerator := float64(c.a)
	denominator := math.Pow(2, float64(c.comboOperand(inst.operand)))
	result := int(numerator / denominator)
	c.c = result
}

func (c *Computer) comboOperand(operand int) int {
	switch operand {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 3
	case 4:
		return c.a
	case 5:
		return c.b
	case 6:
		return c.c
	default:
		panic(fmt.Sprintf("invalid combo operand: %d", operand))
	}
}

func solvePuzzle1(path string) string {
	lines := inputs.ReadLinesFrom(path)
	computer, instructions := readProgram(lines)
	return runProgram(&computer, instructions)
}

func runProgram(computer *Computer, instructions []int) string {
	for computer.instructionPointer < len(instructions) {

		instruction := Instruction{
			opcode:  instructions[computer.instructionPointer],
			operand: instructions[computer.instructionPointer+1],
		}

		switch instruction.opcode {
		case 0:
			computer.adv(instruction)
		case 1:
			computer.bxl(instruction)
		case 2:
			computer.bst(instruction)
		case 3:
			jumped := computer.jnz(instruction)
			if jumped {
				continue
			}
		case 4:
			computer.bxc()
		case 5:
			computer.out(instruction)
		case 6:
			computer.bdv(instruction)
		case 7:
			computer.cdv(instruction)
		default:
			panic(fmt.Sprintf("unknown opcode %d", instruction.opcode))
		}

		computer.instructionPointer += 2
	}

	return strings.Join(computer.output, ",")
}

func solvePuzzle2(path string) int {
	// copied and adapted from: https://github.com/Meez25/AOC2024/blob/main/day17/main.go

	lines := inputs.ReadLinesFrom(path)
	_, instructions := readProgram(lines)

	a := 0
	for i := len(instructions) - 1; i >= 0; i-- {
		a *= 8

		for j := 0; j < 64; j++ {
			candidate := a + j
			c := NewComputer(candidate, 0, 0)

			runProgram(&c, instructions)

			if compareSlices(c.outputInt, instructions[i:]) {
				a = candidate
				break
			}
		}
	}

	return a
}

func readProgram(lines []string) (Computer, []int) {
	numberPattern, _ := regexp.Compile("(\\d+)")

	a, _ := strconv.Atoi(numberPattern.FindString(lines[0]))
	b, _ := strconv.Atoi(numberPattern.FindString(lines[1]))
	c, _ := strconv.Atoi(numberPattern.FindString(lines[2]))
	computer := NewComputer(a, b, c)

	program := numberPattern.FindAllString(lines[4], -1)
	if len(program)%2 != 0 {
		panic("program length must be even")
	}

	instructions := make([]int, 0)

	for i := 0; i < len(program); i++ {
		value, _ := strconv.Atoi(program[i])
		instructions = append(instructions, value)
	}

	return computer, instructions
}

func compareSlices[E comparable](a, b []E) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
