package tabular

// NewTagger returns a new Tagger.
func NewTagger() Tagger {
	return &SetTagger{
		tags: newStringSet(),
	}
}

// Tagger represents set of tags.
type Tagger interface {
	Add(tag string)
	Has(tag string) bool
	HasAll(tags ...string) bool
	HasAny(tags ...string) bool
	Items() []string
	Len() int
}

// SetTagger implements a Tagger using string map set.
type SetTagger struct {
	tags stringSet
}

// Add adds new tag.
func (t *SetTagger) Add(tag string) {
	t.tags.Add(tag)
}

// Has checks if tag in present in the set.
func (t *SetTagger) Has(tag string) bool {
	return t.tags.Contains(tag)
}

// HasAll checks if set has all of the tags.
func (t *SetTagger) HasAll(tags ...string) bool {
	if len(tags) == 0 || t.Len() == 0 {
		return false
	}
	other := newStringSet()
	for _, tag := range tags {
		other.Add(tag)
	}
	return t.tags.Equal(other)
}

// HasAny checks if at least one of the tags is present.
func (t *SetTagger) HasAny(tags ...string) bool {
	if len(tags) == 0 || t.Len() == 0 {
		return false
	}
	other := newStringSet()
	for _, tag := range tags {
		other.Add(tag)
	}
	return t.tags.IsSubset(other)
}

// Items returns all tags as a slice of strings.
func (t *SetTagger) Items() []string {
	tags := make([]string, 0, t.tags.Len())
	t.tags.Each(func(tag string) {
		tags = append(tags, tag)
	})
	return tags
}

// Len returns tag count.
func (t *SetTagger) Len() int {
	return t.tags.Len()
}
