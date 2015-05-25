package tabular

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type JSONWriterTestSuite struct {
	suite.Suite
}

func (s *JSONWriterTestSuite) TestWriteDoubleIndent() {
	opts := &JSONOpts{
		Indent: 4,
	}
	w := NewJSONWriter(opts)
	d, err := newTestDataset()
	s.Nil(err)
	out, err := newTestWrite(d, w)
	expected := `[
    {
        "name": "Julia",
        "surname": "Roberts",
        "age": "40"
    },
    {
        "name": "John",
        "surname": "Malkovich",
        "age": "42"
    }
]`

	s.Nil(err)
	s.Equal(expected, out)
}

func (s *JSONWriterTestSuite) TestWriteIndent() {
	opts := &JSONOpts{
		Indent: 2,
	}
	w := NewJSONWriter(opts)
	d, err := newTestDataset()
	s.Nil(err)
	out, err := newTestWrite(d, w)
	expected := `[
  {
    "name": "Julia",
    "surname": "Roberts",
    "age": "40"
  },
  {
    "name": "John",
    "surname": "Malkovich",
    "age": "42"
  }
]`

	s.Nil(err)
	s.Equal(expected, out)
}

func (s *JSONWriterTestSuite) TestWriteNoIndent() {
	opts := &JSONOpts{
		Indent: 0,
	}
	w := NewJSONWriter(opts)
	d, err := newTestDataset()
	s.Nil(err)
	out, err := newTestWrite(d, w)
	expected := `[{"name":"Julia","surname":"Roberts","age":"40"},{"name":"John","surname":"Malkovich","age":"42"}]`

	s.Nil(err)
	s.Equal(expected, out)
}

func TestJSONWriterTestSuite(t *testing.T) {
	suite.Run(t, new(JSONWriterTestSuite))
}
