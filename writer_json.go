package tabular

import (
	"bufio"
	"io"
	"strings"
)

// JSONOpts represents options passed to the JSON writer.
type JSONOpts struct {
	Indent int
}

// NewJSONWriter creates a new JSON dataset writer.
func NewJSONWriter(opts *JSONOpts) *JSONWriter {
	w := &JSONWriter{opts}
	return w
}

// JSONWriter represents a JSON dataset writer.
type JSONWriter struct {
	opts *JSONOpts
}

// Name returns name of the writer.
func (wj *JSONWriter) Name() string {
	return "json"
}

// NeedsHeaders returns true if headers are required.
func (wj *JSONWriter) NeedsHeaders() bool {
	return true
}

// Write writes dataset to writer.
func (wj *JSONWriter) Write(d *Dataset, w io.Writer) error {
	tw := newJSONTableWriter(d, w, wj.opts)
	return tw.write()
}

func newJSONTableWriter(d *Dataset, w io.Writer, opts *JSONOpts) *jsonTableWriter {
	return &jsonTableWriter{
		d:    d,
		w:    bufio.NewWriter(w),
		opts: opts,
	}
}

type jsonTableWriter struct {
	d    *Dataset
	w    *bufio.Writer
	opts *JSONOpts
	err  error
}

func (j *jsonTableWriter) write() error {
	level := 0
	j.writeIndent("[", level)

	for ridx, row := range j.d.Rows() {
		j.writeIndent("{", level+1)

		for hidx, hdr := range j.d.Headers() {
			j.writeInlineIndent("\"", level+2)
			j.writeEscaped(hdr.Key)

			if j.opts.Indent > 0 {
				j.writeString("\": ")
			} else {
				j.writeString("\":")
			}

			j.writeString("\"")
			j.writeEscaped(row.Get(hidx))
			j.writeString("\"")

			if hidx+1 != j.d.HeaderCount() {
				j.writeString(",")
			}

			j.writeOnIndent("\n")
		}

		j.writeInlineIndent("}", level+1)

		if ridx+1 != j.d.Len() {
			j.writeString(",")
			j.writeOnIndent("\n")
		}
	}

	j.writeOnIndent("\n")
	j.writeString("]")

	return j.flush()
}

func (j *jsonTableWriter) flush() error {
	if j.err != nil {
		return j.err
	}
	return j.w.Flush()
}

func (j *jsonTableWriter) writeOnIndent(s string) {
	if j.opts.Indent > 0 {
		j.writeString(s)
	}
}

func (j *jsonTableWriter) writeIndent(s string, level int) {
	if j.opts.Indent > 0 {
		j.writeString(strings.Repeat(" ", level*j.opts.Indent) + s)
		j.writeString("\n")
	} else {
		j.writeString(s)
	}
}

func (j *jsonTableWriter) writeInlineIndent(s string, level int) {
	if j.opts.Indent > 0 {
		j.writeString(strings.Repeat(" ", level*j.opts.Indent) + s)
	} else {
		j.writeString(s)
	}
}

func (j *jsonTableWriter) writeString(s string) {
	if j.err != nil {
		return
	}
	_, err := j.w.WriteString(s)
	j.err = err
}

func (j *jsonTableWriter) writeEscaped(s string) {
	j.writeString(j.escapeString(s))
}

func (j *jsonTableWriter) escapeString(s string) string {
	return s
}
