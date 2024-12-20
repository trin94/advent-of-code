package utils

func Clone[E comparable](slice []E) []E {
	var dst []E
	return append(dst, slice...)
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
