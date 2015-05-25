package tabular

import (
	"encoding/xml"
	"io"
	"strings"
)

// XMLOpts represents options passed to the XML writer.
type XMLOpts struct {
	Indent int

	RowElem    string
	ParentElem string
}

// NewXMLWriter creates a new XML dataset writer.
func NewXMLWriter(opts *XMLOpts) *XMLWriter {
	w := &XMLWriter{opts}
	return w
}

// XMLWriter represents a XML dataset writer.
type XMLWriter struct {
	opts *XMLOpts
}

// Name returns name of the writer.
func (wx *XMLWriter) Name() string {
	return "xml"
}

// NeedsHeaders returns true if headers are required.
func (wx *XMLWriter) NeedsHeaders() bool {
	return true
}

// Write writes dataset to writer.
func (wx *XMLWriter) Write(d *Dataset, w io.Writer) error {
	wr := wx.newXMLWrapper(d)
	enc := xml.NewEncoder(w)
	enc.Indent("", strings.Repeat(" ", wx.opts.Indent))
	return enc.Encode(wr)
}

func (wx *XMLWriter) newXMLWrapper(d *Dataset) *xmlWrapper {
	wrapper := &xmlWrapper{
		writer: wx,
		d:      d,
	}
	return wrapper
}

type xmlWrapper struct {
	writer *XMLWriter
	d      *Dataset
}

func (xw xmlWrapper) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{
		Space: "",
		Local: xw.writer.opts.ParentElem,
	}
	if err := e.EncodeToken(start); err != nil {
		return err
	}

	for _, row := range xw.d.Rows() {
		if err := xw.encodeRow(e, row); err != nil {
			return err
		}
	}

	return e.EncodeToken(start.End())
}

func (xw xmlWrapper) encodeRow(e *xml.Encoder, row *Row) error {
	elem := xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: xw.writer.opts.RowElem,
		},
		Attr: nil,
	}
	if err := e.EncodeToken(elem); err != nil {
		return err
	}

	for idx, val := range row.Items() {
		if err := xw.encodeItem(e, idx, val); err != nil {
			return err
		}
	}

	return e.EncodeToken(elem.End())
}

func (xw xmlWrapper) encodeItem(e *xml.Encoder, idx int, val string) error {
	h, ok := xw.d.GetHeader(idx)
	if !ok {
		return ErrInvalidHeaderIndex{idx}
	}

	elem := xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: h.Key,
		},
		Attr: nil,
	}
	return e.EncodeElement(val, elem)
}
