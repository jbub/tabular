package tabular

import (
	"sort"
)

// RowSorter implements a row sorter.
type RowSorter struct {
	idx     int
	reverse bool

	rows []*Row
}

// Len returns the row count.
func (r RowSorter) Len() int {
	return len(r.rows)
}

// Swap swaps two rows.
func (r RowSorter) Swap(i, j int) {
	r.rows[i], r.rows[j] = r.rows[j], r.rows[i]
}

// Less compares two rows.
func (r RowSorter) Less(i, j int) bool {
	if r.reverse {
		return r.rows[i].Get(r.idx) > r.rows[j].Get(r.idx)
	}
	return r.rows[i].Get(r.idx) < r.rows[j].Get(r.idx)
}

// Sort sorts rows and returns sorted slice.
func (r RowSorter) Sort() []*Row {
	sort.Sort(r)
	return r.rows
}
