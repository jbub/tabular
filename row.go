package tabular

// NewRow creates new row with optional items.
func NewRow(items ...string) *Row {
	r := &Row{}
	r.Add(items...)
	r.tagger = NewTagger()
	return r
}

// NewRowFromSlice creates new row from slice of items.
func NewRowFromSlice(items []string) *Row {
	return NewRow(items...)
}

// Row represents a row of dataset.
type Row struct {
	items  []string
	tagger Tagger
}

// Add appends new items to the row.
func (r *Row) Add(items ...string) {
	r.items = append(r.items, items...)
}

// Get returns tow item on given index .
func (r *Row) Get(idx int) string {
	return r.items[idx]
}

// Items returns slice of row items.
func (r *Row) Items() []string {
	return r.items
}

// Len returns row item count.
func (r *Row) Len() int {
	return len(r.items)
}

// AddTag append new tag to the row.
func (r *Row) AddTag(tag string) {
	r.tagger.Add(tag)
}

// HasTag checks for tag presence of row.
func (r *Row) HasTag(tag string) bool {
	return r.tagger.Has(tag)
}

// HasAllTags checks for presence of all tags in the row.
func (r *Row) HasAllTags(tags ...string) bool {
	return r.tagger.HasAll(tags...)
}

// HasAnyTags checks for presence of any tags in the row.
func (r *Row) HasAnyTags(tags ...string) bool {
	return r.tagger.HasAny(tags...)
}

// Tags returns slice of row tags.
func (r *Row) Tags() []string {
	return r.tagger.Items()
}
