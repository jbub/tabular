package tabular

import (
	"fmt"
	"io"
	"strings"
)

// ErrInvalidHeaderIndex is error returned when writer is missing the headers.
type ErrInvalidHeaderIndex struct {
	idx int
}

func (e ErrInvalidHeaderIndex) Error() string {
	return fmt.Sprintf("Invalid header index %d.", e.idx)
}

// Writer represents a dataset writer.
type Writer interface {
	// Name returns name of the writer.
	Name() string

	// NeedsHeaders returns true if headers are required.
	NeedsHeaders() bool

	// Write writes dataset to writer.
	Write(d *Dataset, w io.Writer) error
}

func padString(s string, total int) string {
	length := len(s)
	if length >= total {
		return s
	}
	return s + strings.Repeat(" ", total-length)
}
