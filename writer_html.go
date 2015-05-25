package tabular

import (
	"bufio"
	"io"
	"strings"
)

// HTMLOpts represents options passed to the HTML writer.
type HTMLOpts struct {
	Caption string
	Indent  int

	TableClass string
	RowClass   string
	HeadClass  string
	DataClass  string
}

// NewHTMLWriter creates a new HTML dataset writer.
func NewHTMLWriter(opts *HTMLOpts) *HTMLWriter {
	w := &HTMLWriter{opts}
	return w
}

// HTMLWriter represents a HTML dataset writer.
type HTMLWriter struct {
	opts *HTMLOpts
}

// Name returns name of the writer.
func (wh *HTMLWriter) Name() string {
	return "html"
}

// NeedsHeaders returns true if headers are required.
func (wh *HTMLWriter) NeedsHeaders() bool {
	return false
}

// Write writes dataset to writer.
func (wh *HTMLWriter) Write(d *Dataset, w io.Writer) error {
	tw := newHTMLTableWriter(d, w, wh.opts)
	return tw.write()
}

func newHTMLTableWriter(d *Dataset, w io.Writer, opts *HTMLOpts) *htmlTableWriter {
	return &htmlTableWriter{
		d:    d,
		w:    bufio.NewWriter(w),
		opts: opts,
	}
}

type htmlTableWriter struct {
	d    *Dataset
	w    *bufio.Writer
	opts *HTMLOpts
	err  error
}

func (h *htmlTableWriter) write() error {
	level := 0
	h.writeStartElem("table", level, h.opts.TableClass, true)

	if h.opts.Caption != "" {
		h.writeInlineElem("caption", h.opts.Caption, "", level+1)
	}

	if h.d.HasHeaders() {
		h.writeStartElem("thead", level+1, "", true)
		h.writeHeaders(level + 2)
		h.writeEndElem("thead", level+1, true)
	}

	h.writeStartElem("tbody", level+1, "", true)
	h.writeRows(level + 2)
	h.writeEndElem("tbody", level+1, true)

	h.writeEndElem("table", level+0, true)

	return h.flush()
}

func (h *htmlTableWriter) writeHeaders(level int) {
	h.writeStartElem("tr", level, h.opts.RowClass, true)
	for _, hdr := range h.d.Headers() {
		h.writeHeader(hdr, level+1)
	}
	h.writeEndElem("tr", level, true)
}

func (h *htmlTableWriter) writeHeader(hdr *Header, level int) {
	h.writeInlineElem("th", hdr.Title, h.opts.HeadClass, level)
}

func (h *htmlTableWriter) writeRows(level int) {
	for _, row := range h.d.Rows() {
		h.writeRow(row, level)
	}
}

func (h *htmlTableWriter) writeRow(row *Row, level int) {
	h.writeStartElem("tr", level, h.opts.RowClass, true)
	for _, item := range row.Items() {
		h.writeRowItem(item, level+1)
	}
	h.writeEndElem("tr", level, true)
}

func (h *htmlTableWriter) writeRowItem(item string, level int) {
	h.writeInlineElem("td", item, h.opts.DataClass, level)
}

func (h *htmlTableWriter) writeElem(name string, val string, class string, level int) {
	h.writeStartElem(name, level, class, false)
	h.writeString(val)
	h.writeEndElem(name, level, true)
}

func (h *htmlTableWriter) writeInlineElem(name string, val string, class string, level int) {
	h.writeStartElem(name, level, class, false)
	h.writeString(val)
	h.writeInlineEndElem(name, level, true)
}

func (h *htmlTableWriter) writeStartElem(name string, level int, class string, newline bool) {
	if class != "" {
		h.writeIndent(`<`+name+` class="`+class+`">`, level, true, false, newline)
	} else {
		h.writeIndent("<"+name+">", level, true, false, newline)
	}
}

func (h *htmlTableWriter) writeEndElem(name string, level int, newline bool) {
	h.writeIndent("</"+name+">", level, false, false, newline)
}

func (h *htmlTableWriter) writeInlineEndElem(name string, level int, newline bool) {
	h.writeIndent("</"+name+">", level, false, true, newline)
}

func (h *htmlTableWriter) writeIndent(val string, level int, start bool, inline bool, newline bool) {
	hasIndent := h.opts.Indent > 0
	if hasIndent && (start || !inline) {
		h.writeString(strings.Repeat(" ", h.opts.Indent*level))
	}

	h.writeString(val)
	if hasIndent && newline {
		h.writeString("\n")
	}
}

func (h *htmlTableWriter) writeString(val string) {
	if h.err != nil {
		return
	}
	_, err := h.w.WriteString(val)
	h.err = err
}

func (h *htmlTableWriter) flush() error {
	if h.err != nil {
		return h.err
	}
	return h.w.Flush()
}
