package tabular

type stringSet map[string]struct{}

func newStringSet() stringSet {
	return make(stringSet)
}

func (s *stringSet) Add(i string) bool {
	if _, found := (*s)[i]; found {
		return false
	}

	(*s)[i] = struct{}{}
	return true
}

func (s *stringSet) Contains(i ...string) bool {
	for _, val := range i {
		if _, ok := (*s)[val]; !ok {
			return false
		}
	}
	return true
}

func (s *stringSet) Items() []string {
	elems := make([]string, 0, len(*s))
	for elem := range *s {
		elems = append(elems, elem)
	}
	return elems
}

func (s *stringSet) Len() int {
	return len(*s)
}
