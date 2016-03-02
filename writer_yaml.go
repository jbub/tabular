package tabular

import (
	"bufio"
	"io"
	"strings"
)

// YAMLOpts represents options passed to the YAML writer.
type YAMLOpts struct {
}

// NewYAMLWriter creates a new YAML dataset writer.
func NewYAMLWriter(opts *YAMLOpts) *YAMLWriter {
	w := &YAMLWriter{opts}
	return w
}

// YAMLWriter represents a YAML dataset writer.
type YAMLWriter struct {
	opts *YAMLOpts
}

// Name returns name of the writer.
func (wy *YAMLWriter) Name() string {
	return "yaml"
}

// NeedsHeaders returns true if headers are required.
func (wy *YAMLWriter) NeedsHeaders() bool {
	return true
}

// Write writes dataset to writer.
func (wy *YAMLWriter) Write(d *Dataset, w io.Writer) error {
	tw := newYAMLTableWriter(d, w, wy.opts)
	return tw.write()
}

// http://symfony.com/doc/current/components/yaml/yaml_format.html
var yamlReplacements = []string{
	"&", "\\&",
	"%", "\\%",
	"$", "\\$",
	"#", "\\#",
	"{", "\\{",
	"}", "\\}",
}

func newYAMLTableWriter(d *Dataset, w io.Writer, opts *YAMLOpts) *yamlTableWriter {
	return &yamlTableWriter{
		d:        d,
		w:        bufio.NewWriter(w),
		opts:     opts,
		replacer: strings.NewReplacer(yamlReplacements...),
	}
}

type yamlTableWriter struct {
	d    *Dataset
	w    *bufio.Writer
	opts *YAMLOpts
	err  error

	replacer *strings.Replacer
}

func (y *yamlTableWriter) write() error {
	for _, row := range y.d.Rows() {
		y.writeString("- ")
		for idx, hdr := range y.d.Headers() {
			if idx != 0 {
				y.writeString("  ")
			}
			y.writeEscaped(hdr.Key)
			y.writeString(": ")
			y.writeEscaped(row.Get(idx))
			y.writeString("\n")
		}
	}
	return y.flush()
}

func (y *yamlTableWriter) flush() error {
	if y.err != nil {
		return y.err
	}
	return y.w.Flush()
}

func (y *yamlTableWriter) writeString(s string) {
	if y.err != nil {
		return
	}
	_, err := y.w.WriteString(s)
	y.err = err
}

func (l *yamlTableWriter) escapeKey(s string) string {
	return l.escapeString(s)
}

func (l *yamlTableWriter) escapeValue(s string) string {
	return l.escapeString(s)
}

func (y *yamlTableWriter) writeEscaped(s string) {
	y.writeString(y.escapeString(s))
}

func (y *yamlTableWriter) escapeString(s string) string {
	return y.replacer.Replace(s)
}
