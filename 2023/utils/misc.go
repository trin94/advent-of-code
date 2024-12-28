package utils

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
