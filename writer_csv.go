package tabular

import (
	"encoding/csv"
	"io"
)

// CSVOpts represents options passed to the CSV writer.
type CSVOpts struct {
	Comma   rune
	UseCRLF bool
}

// NewCSVWriter creates a new CSV dataset writer.
func NewCSVWriter(opts *CSVOpts) *CSVWriter {
	w := &CSVWriter{opts}
	return w
}

// CSVWriter represents a CSV dataset writer.
type CSVWriter struct {
	opts *CSVOpts
}

// Name returns name of the writer.
func (wc *CSVWriter) Name() string {
	return "csv"
}

// NeedsHeaders returns true if headers are required.
func (wc *CSVWriter) NeedsHeaders() bool {
	return false
}

// Write writes dataset to writer.
func (wc *CSVWriter) Write(d *Dataset, w io.Writer) error {
	cw := csv.NewWriter(w)
	cw.Comma = wc.opts.Comma
	cw.UseCRLF = wc.opts.UseCRLF

	if d.HasHeaders() {
		var hdrs []string
		for _, hdr := range d.Headers() {
			hdrs = append(hdrs, hdr.Title)
		}
		if err := cw.Write(hdrs); err != nil {
			return err
		}
	}

	for _, row := range d.rows {
		err := cw.Write(row.Items())
		if err != nil {
			return err
		}
	}

	cw.Flush()
	return cw.Error()
}
