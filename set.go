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

func (s *stringSet) Each(cb func(string)) {
	for elem := range *s {
		cb(elem)
	}
}

func (s *stringSet) Equal(other stringSet) bool {
	if s.Len() != other.Len() {
		return false
	}
	for elem := range *s {
		if !other.Contains(elem) {
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
