package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var lineA = "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}"
var machineA = Machine{
	lightDiagram: LightDiagram{
		[]bool{false, true, true, false},
	},
	buttonSemantics: []ButtonWiringSemantic{
		{pressedOn: []int{3}},
		{pressedOn: []int{1, 3}},
		{pressedOn: []int{2}},
		{pressedOn: []int{2, 3}},
		{pressedOn: []int{0, 2}},
		{pressedOn: []int{0, 1}},
	},
	joltageRequirements: JoltageRequirement{numbers: []int{3, 5, 4, 7}},
}

var lineB = "[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}"
var machineB = Machine{
	lightDiagram: LightDiagram{
		[]bool{false, false, false, true, false},
	},
	buttonSemantics: []ButtonWiringSemantic{
		{pressedOn: []int{0, 2, 3, 4}},
		{pressedOn: []int{2, 3}},
		{pressedOn: []int{0, 4}},
		{pressedOn: []int{0, 1, 2}},
		{pressedOn: []int{1, 2, 3, 4}},
	},
	joltageRequirements: JoltageRequirement{numbers: []int{7, 5, 12, 7, 2}},
}

func TestParseMachines(t *testing.T) {
	var testParseLineParams = []struct {
		line     string
		expected Machine
	}{
		{line: lineA, expected: machineA},
		{line: lineB, expected: machineB},
	}

	for _, tt := range testParseLineParams {
		assert.Equal(t, tt.expected, NewMachine(tt.line))
	}

}

func TestMachine_CalculateFewestButtonPressesToTurnOn(t *testing.T) {
	var params = []struct {
		machine  Machine
		expected int
	}{
		{machine: machineA, expected: 2},
		{machine: machineB, expected: 3},
	}

	for _, tt := range params {
		assert.Equal(t, tt.expected, tt.machine.CalculateFewestButtonPressesToTurnOn())
	}
}

func TestWireUpButton(t *testing.T) {
	current := []bool{false, false, false, false}

	current = wireUpButton(current, []int{0, 2})
	current = wireUpButton(current, []int{0, 1})

	expected := []bool{false, true, true, false}

	assert.Equal(t, expected, current)
}

func TestVoltageIncrease(t *testing.T) {
	current := []int{3, 5, 4, 7}

	current = decreaseVoltage(current, []int{3})
	current = decreaseVoltage(current, []int{1, 3})
	current = decreaseVoltage(current, []int{1, 3})
	current = decreaseVoltage(current, []int{1, 3})
	current = decreaseVoltage(current, []int{2, 3})
	current = decreaseVoltage(current, []int{2, 3})
	current = decreaseVoltage(current, []int{2, 3})
	current = decreaseVoltage(current, []int{0, 2})
	current = decreaseVoltage(current, []int{0, 1})
	current = decreaseVoltage(current, []int{0, 1})

	expected := []int{0, 0, 0, 0}
	assert.Equal(t, expected, current)
}

func TestMachine_CalculateFewestButtonPressesToReachVoltage(t *testing.T) {
	var params = []struct {
		machine  Machine
		expected int
	}{
		{machine: machineA, expected: 10},
		{machine: machineB, expected: 12},
	}

	for _, tt := range params {
		assert.Equal(t, tt.expected, tt.machine.CalculateFewestButtonPressesToReachVoltage())
	}
}
