package utils

import "strings"

func Clone[E comparable](slice []E) []E {
	var dst []E
	return append(dst, slice...)
}

func Filter[E any](slice []E, predicate func(E) bool) (ret []E) {
	for _, s := range slice {
		if predicate(s) {
			ret = append(ret, s)
		}
	}
	return
}

func RotateClockwise(lines []string) []string {
	maxRows, maxColumns := len(lines), len(lines[0])
	result := make([]string, maxColumns)
	for col := 0; col < maxColumns; col++ {
		var sb strings.Builder
		for row := maxRows - 1; row >= 0; row-- {
			sb.WriteByte(lines[row][col])
		}
		result[col] = sb.String()
	}
	return result
}
