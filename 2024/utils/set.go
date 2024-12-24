package utils

type Set[E comparable] map[E]struct{}

func NewSet[E comparable]() Set[E] {
	return Set[E]{}
}

func (s *Set[E]) Add(e E) {
	(*s)[e] = struct{}{}
}

func (s *Set[E]) Contains(v E) bool {
	_, found := (*s)[v]
	return found
}

func (s *Set[E]) Values() []E {
	result := make([]E, 0)
	for key := range *s {
		result = append(result, key)
	}
	return result
}
