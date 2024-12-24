package utils

import "math"

func Clone[E comparable](slice []E) []E {
	var dst []E
	return append(dst, slice...)
}

// Remove the element at index. Careful: This does not keep the order of elements
func Remove[E comparable](slice []E, index int) []E {
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func AbsInt(a int) int {
	return int(math.Abs(float64(a)))
}
