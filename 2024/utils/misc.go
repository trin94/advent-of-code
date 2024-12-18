package utils

func Clone[E comparable](slice []E) []E {
	var dst []E
	return append(dst, slice...)
}
