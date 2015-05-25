package tabular

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type XMLWriterTestSuite struct {
	suite.Suite
}

func (s *XMLWriterTestSuite) TestWrite() {
	opts := &XMLOpts{
		Indent:     2,
		RowElem:    "row",
		ParentElem: "rows",
	}
	w := NewXMLWriter(opts)
	d, err := newTestDataset()
	s.Nil(err)
	out, err := newTestWrite(d, w)
	expected :=
		`<rows>
  <row>
    <name>Julia</name>
    <surname>Roberts</surname>
    <age>40</age>
  </row>
  <row>
    <name>John</name>
    <surname>Malkovich</surname>
    <age>42</age>
  </row>
</rows>`

	s.Nil(err)
	s.Equal(expected, out)
}

func TestXMLWriterTestSuite(t *testing.T) {
	suite.Run(t, new(XMLWriterTestSuite))
}
