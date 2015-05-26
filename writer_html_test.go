package tabular

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type HTMLWriterTestSuite struct {
	suite.Suite
}

func (s *HTMLWriterTestSuite) TestWriteNoIndent() {
	opts := &HTMLOpts{
		Indent: 0,
	}
	w := NewHTMLWriter(opts)
	d, err := newTestDataset()
	s.Nil(err)
	out, err := newTestWrite(d, w)
	expected :=
		`<table><thead><tr><th>First name</th><th>Last name</th><th>Age</th></tr></thead><tbody><tr><td>Julia</td><td>Roberts</td><td>40</td></tr><tr><td>John</td><td>Malkovich</td><td>42</td></tr></tbody></table>`

	s.Nil(err)
	s.Equal(expected, out)
}

func (s *HTMLWriterTestSuite) TestWriteIndent() {
	opts := &HTMLOpts{
		Indent: 2,
	}
	w := NewHTMLWriter(opts)
	d, err := newTestDataset()
	s.Nil(err)
	out, err := newTestWrite(d, w)
	expected :=
		`<table>
  <thead>
    <tr>
      <th>First name</th>
      <th>Last name</th>
      <th>Age</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>Julia</td>
      <td>Roberts</td>
      <td>40</td>
    </tr>
    <tr>
      <td>John</td>
      <td>Malkovich</td>
      <td>42</td>
    </tr>
  </tbody>
</table>
`

	s.Nil(err)
	s.Equal(expected, out)
}

func (s *HTMLWriterTestSuite) TestWriteDoubleIndent() {
	opts := &HTMLOpts{
		Indent: 4,
	}
	w := NewHTMLWriter(opts)
	d, err := newTestDataset()
	s.Nil(err)
	out, err := newTestWrite(d, w)
	expected :=
		`<table>
    <thead>
        <tr>
            <th>First name</th>
            <th>Last name</th>
            <th>Age</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Julia</td>
            <td>Roberts</td>
            <td>40</td>
        </tr>
        <tr>
            <td>John</td>
            <td>Malkovich</td>
            <td>42</td>
        </tr>
    </tbody>
</table>
`

	s.Nil(err)
	s.Equal(expected, out)
}

func (s *HTMLWriterTestSuite) TestWriteClass() {
	opts := &HTMLOpts{
		Indent:     4,
		TableClass: "table",
		DataClass:  "data",
		HeadClass:  "head",
		RowClass:   "row",
	}
	w := NewHTMLWriter(opts)
	d, err := newTestDataset()
	s.Nil(err)
	out, err := newTestWrite(d, w)
	expected :=
		`<table class="table">
    <thead>
        <tr class="row">
            <th class="head">First name</th>
            <th class="head">Last name</th>
            <th class="head">Age</th>
        </tr>
    </thead>
    <tbody>
        <tr class="row">
            <td class="data">Julia</td>
            <td class="data">Roberts</td>
            <td class="data">40</td>
        </tr>
        <tr class="row">
            <td class="data">John</td>
            <td class="data">Malkovich</td>
            <td class="data">42</td>
        </tr>
    </tbody>
</table>
`

	s.Nil(err)
	s.Equal(expected, out)
}

func (s *HTMLWriterTestSuite) TestWriteCaption() {
	opts := &HTMLOpts{
		Indent:  4,
		Caption: "Oh my!",
	}
	w := NewHTMLWriter(opts)
	d, err := newTestDataset()
	s.Nil(err)
	out, err := newTestWrite(d, w)
	expected :=
		`<table>
    <caption>Oh my!</caption>
    <thead>
        <tr>
            <th>First name</th>
            <th>Last name</th>
            <th>Age</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Julia</td>
            <td>Roberts</td>
            <td>40</td>
        </tr>
        <tr>
            <td>John</td>
            <td>Malkovich</td>
            <td>42</td>
        </tr>
    </tbody>
</table>
`

	s.Nil(err)
	s.Equal(expected, out)
}

func (s *HTMLWriterTestSuite) TestWriteCaptionEscape() {
	opts := &HTMLOpts{
		Indent:     4,
		Caption:    "Oh m\"y!",
		TableClass: "\"dsa",
	}
	w := NewHTMLWriter(opts)
	d, err := newTestDataset()
	s.Nil(err)
	out, err := newTestWrite(d, w)
	expected :=
		`<table class="&#34;dsa">
    <caption>Oh m&#34;y!</caption>
    <thead>
        <tr>
            <th>First name</th>
            <th>Last name</th>
            <th>Age</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Julia</td>
            <td>Roberts</td>
            <td>40</td>
        </tr>
        <tr>
            <td>John</td>
            <td>Malkovich</td>
            <td>42</td>
        </tr>
    </tbody>
</table>
`

	s.Nil(err)
	s.Equal(expected, out)
}

func TestHTMLWriterTestSuite(t *testing.T) {
	suite.Run(t, new(HTMLWriterTestSuite))
}
