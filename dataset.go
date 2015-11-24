package tabular

import (
	"errors"
	"fmt"
	"io"
)

var (
	// ErrEmptyDataset is returned when operations are applied to empty dataset.
	ErrEmptyDataset = errors.New("Dataset is empty.")
)

// ErrInvalidRowWidth is error returned when adding new row with invalid width.
type ErrInvalidRowWidth struct {
	actual   int
	expected int
}

func (e ErrInvalidRowWidth) Error() string {
	return fmt.Sprintf("Invalid row width = %d, expected = %d.", e.actual, e.expected)
}

// ErrHeadersRequired is error returned when writer is missing the headers.
type ErrHeadersRequired struct {
	w Writer
}

func (e ErrHeadersRequired) Error() string {
	return fmt.Sprintf("Writer %s needs headers.", e.w.Name())
}

// NewDataSet creates new dataset.
func NewDataSet() *Dataset {
	d := &Dataset{
		headers: &Headers{},
	}
	return d
}

// Dataset represents a set of data.
type Dataset struct {
	headers *Headers
	rows    []*Row

	columns int
	lengths map[int]int
}

// AddHeader adds new header.
func (d *Dataset) AddHeader(key string, title string) {
	d.headers.Add(key, title)
	d.updateHeaders()
}

// GetHeader returns header by its index.
func (d *Dataset) GetHeader(idx int) (*Header, bool) {
	return d.headers.Get(idx)
}

// HeaderCount returns header count.
func (d *Dataset) HeaderCount() int {
	return d.headers.Len()
}

// Headers returns a slice of headers.
func (d *Dataset) Headers() []*Header {
	return d.headers.Items()
}

// HasHeaders returns true if headers was set.
func (d *Dataset) HasHeaders() bool {
	return d.headers.Len() > 0
}

// Append appends new rows to the dataset.
func (d *Dataset) Append(rows ...*Row) error {
	for _, r := range rows {
		if err := d.validateRow(r); err != nil {
			return err
		}
		d.rows = append(d.rows, r)
		d.updateLengths(r)
	}
	return nil
}

// HasCol checks the presence of given column.
func (d *Dataset) HasCol(key string) bool {
	_, ok := d.getColumnIndex(key)
	return ok
}

// GetColValues returns values of given column.
func (d *Dataset) GetColValues(key string) []string {
	if idx, ok := d.getColumnIndex(key); ok {
		var col []string
		for _, row := range d.rows {
			col = append(col, row.Get(idx))
		}
		return col
	}
	return nil
}

// GetKeyWidth returns maximum column width.
func (d *Dataset) GetKeyWidth(key string) int {
	if idx, ok := d.getColumnIndex(key); ok {
		return d.getIndexWidth(idx)
	}
	return 0
}

// GetIdxWidth returns maximum column width.
func (d *Dataset) GetIdxWidth(idx int) int {
	return d.getIndexWidth(idx)
}

// Find filters dataset for tag.
func (d *Dataset) Find(tag string) *Dataset {
	var rows []*Row
	for _, row := range d.rows {
		if row.HasTag(tag) {
			rows = append(rows, row)
		}
	}
	d.rows = rows
	return d
}

// FindAny filters dataset for any tags.
func (d *Dataset) FindAny(tags ...string) *Dataset {
	var rows []*Row
	for _, row := range d.rows {
		if row.HasAnyTags(tags...) {
			rows = append(rows, row)
		}
	}
	d.rows = rows
	return d
}

// FindAll filters dataset for all tags.
func (d *Dataset) FindAll(tags ...string) *Dataset {
	var rows []*Row
	for _, row := range d.rows {
		if row.HasAllTags(tags...) {
			rows = append(rows, row)
		}
	}
	d.rows = rows
	return d
}

// Slice returns sliced dataset.
func (d *Dataset) Slice(start int, end int) *Dataset {
	d.rows = d.rows[start:end]
	return d
}

// Sort sorts dataset by key, set reverse to inverse the order direction.
func (d *Dataset) Sort(key string, reverse bool) *Dataset {
	if idx, ok := d.getColumnIndex(key); ok {
		sorter := &RowSorter{
			rows:    d.rows,
			idx:     idx,
			reverse: reverse,
		}
		d.rows = sorter.Sort()
	}
	return d
}

// Write writes dataset using dataset writer to writer.
func (d *Dataset) Write(dw Writer, w io.Writer) error {
	if d.Len() == 0 {
		return ErrEmptyDataset
	}

	if dw.NeedsHeaders() && d.headers.Empty() {
		return ErrHeadersRequired{dw}
	}

	return dw.Write(d, w)
}

// Get returns a row on given index.
func (d *Dataset) Get(idx int) (*Row, bool) {
	if d.isValidIndex(idx) {
		return d.rows[idx], true
	}
	return nil, false
}

// Rows returns all rows of dataset.
func (d *Dataset) Rows() []*Row {
	return d.rows
}

// Len returns the row count.
func (d *Dataset) Len() int {
	return len(d.rows)
}

func (d *Dataset) validateRow(r *Row) error {
	columns := r.Len()
	if columns < 1 {
		return ErrInvalidRowWidth{
			actual:   columns,
			expected: 1,
		}
	}

	if d.Len() > 0 && d.columns != columns {
		return ErrInvalidRowWidth{
			actual:   columns,
			expected: d.columns,
		}
	}

	return nil
}

func (d *Dataset) updateHeaders() {
	d.columns = d.HeaderCount()
}

func (d *Dataset) updateLengths(r *Row) {
	if !d.HasHeaders() {
		d.columns = r.Len()
	}

	if d.lengths == nil {
		d.lengths = make(map[int]int)
	}

	for idx, item := range r.Items() {
		l := len(item)
		if l > d.lengths[idx] {
			d.lengths[idx] = l
		}
	}
}

func (d *Dataset) getColumnIndex(key string) (int, bool) {
	for i, col := range d.headers.Items() {
		if col.Key == key {
			return i, true
		}
	}
	return 0, false
}

func (d *Dataset) getIndexWidth(idx int) int {
	if d.isValidIndex(idx) {
		length := d.lengths[idx]
		if d.HasHeaders() {
			if h, ok := d.GetHeader(idx); ok {
				hdrWidth := len(h.Title)
				if hdrWidth > length {
					return hdrWidth
				}
			}
		}
		return length
	}
	return d.lengths[idx]
}

func (d *Dataset) isValidIndex(idx int) bool {
	if d.Len() == 0 {
		return false
	}
	return idx >= 0 && idx <= d.Len()
}
