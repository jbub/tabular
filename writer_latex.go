package tabular

import (
	"bufio"
	"io"
)

// LatexOpts represents options passed to the LaTeX writer.
type LatexOpts struct {
	Caption  string
	Center   bool
	TabularX bool
}

// NewLatexWriter creates a new LaTeX dataset writer.
func NewLatexWriter(opts *LatexOpts) *LatexWriter {
	w := &LatexWriter{opts}
	return w
}

// LatexWriter represents a LaTeX dataset writer.
type LatexWriter struct {
	opts *LatexOpts
}

// Name returns name of the writer.
func (wl *LatexWriter) Name() string {
	return "latex"
}

// NeedsHeaders returns true if headers are required.
func (wl *LatexWriter) NeedsHeaders() bool {
	return false
}

// Write writes dataset to writer.
func (wl *LatexWriter) Write(d *Dataset, w io.Writer) error {
	tw := newLatexTableWriter(d, w, wl.opts)
	return tw.write()
}

func newLatexTableWriter(d *Dataset, w io.Writer, opts *LatexOpts) *latexTableWriter {
	return &latexTableWriter{
		d:    d,
		w:    bufio.NewWriter(w),
		opts: opts,
	}
}

type latexTableWriter struct {
	d    *Dataset
	w    *bufio.Writer
	opts *LatexOpts
	err  error
}

func (l *latexTableWriter) write() error {
	l.writeString("\\begin{table}[h]\n")
	l.writeString("\\begin{tabular}{|l|l|l|}\n")
	l.writeString("\\hline\n")

	if l.d.HasHeaders() {
		l.writeHeaders()
	}

	l.writeRows()

	l.writeString("\\end{tabular}\n")
	l.writeString("\\end{table}\n")

	return l.flush()
}

func (l *latexTableWriter) flush() error {
	if l.err != nil {
		return l.err
	}
	return l.w.Flush()
}

func (l *latexTableWriter) writeHeaders() {
	for idx, hdr := range l.d.Headers() {
		l.writeHeader(idx, hdr)
	}
	l.writeString(" \\\\ \\hline\n")
}

func (l *latexTableWriter) writeHeader(idx int, hdr *Header) {
	width := l.d.GetIdxWidth(idx)
	l.writeString(padString(hdr.Title, width))

	if l.d.columns > idx+1 {
		l.writeString(" & ")
	}
}

func (l *latexTableWriter) writeRows() {
	for _, row := range l.d.Rows() {
		l.writeRow(row)
	}
}

func (l *latexTableWriter) writeRow(r *Row) {
	for idx, item := range r.Items() {
		l.writeItem(idx, item)
	}
	l.writeString(" \\\\ \\hline\n")
}

func (l *latexTableWriter) writeItem(idx int, item string) {
	width := l.d.GetIdxWidth(idx)
	padded := padString(item, width)
	l.writeString(padded)

	if l.d.columns > idx+1 {
		l.writeString(" & ")
	}
}

func (l *latexTableWriter) writeString(s string) {
	if l.err != nil {
		return
	}
	_, err := l.w.WriteString(s)
	l.err = err
}

func (l *latexTableWriter) escapeString(s string) string {
	// TODO escape
	// & % $ # _ { } => with backslashes
	// ~ ^ \ => \textasciitilde, \textasciicircum, and \textbackslash.
	return s
}
