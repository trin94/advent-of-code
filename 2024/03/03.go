package p03

import (
	"math"
	"regexp"
	"strconv"
	"trin94/aoc/2024/inputs"
)

var MultiPattern, _ = regexp.Compile("mul\\((\\d+),(\\d+)\\)")
var DoPattern, _ = regexp.Compile("do\\(\\)")
var DontPattern, _ = regexp.Compile("don't\\(\\)")

func solvePuzzle1(path string) int {
	lines := inputs.ReadLinesFrom(path)

	var sum int

	for _, line := range lines {
		for _, value := range MultiPattern.FindAllStringSubmatch(line, -1) {
			x, _ := strconv.Atoi(value[1])
			y, _ := strconv.Atoi(value[2])
			sum += x * y
		}
	}

	return sum
}

func solvePuzzle2(path string) int {
	lines := inputs.ReadLinesFrom(path)

	var sum int
	enabled := true

	for _, line := range lines {

		lineLeft := line

		for {

			if enabled {
				var idxDont = math.MaxInt
				var idxMulti = math.MaxInt

				nextDont := DontPattern.FindStringIndex(lineLeft)
				if nextDont != nil {
					idxDont = nextDont[1]
				}

				nextMulti := MultiPattern.FindStringIndex(lineLeft)
				if nextMulti != nil {
					idxMulti = nextMulti[1]
				}

				if idxDont == math.MaxInt && idxMulti == math.MaxInt {
					break
				}

				if idxMulti < idxDont {

					value := MultiPattern.FindStringSubmatch(lineLeft)
					x, _ := strconv.Atoi(value[1])
					y, _ := strconv.Atoi(value[2])
					sum += x * y
					lineLeft = lineLeft[idxMulti:]

				} else {

					enabled = false
					lineLeft = lineLeft[idxDont:]

				}

			} else {

				nextDo := DoPattern.FindStringIndex(lineLeft)
				if nextDo == nil {
					break
				}

				idxDo := nextDo[1]
				enabled = true
				lineLeft = lineLeft[idxDo:]
			}

		}

	}

	return sum
}
