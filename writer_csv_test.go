package tabular

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CSVWriterTestSuite struct {
	suite.Suite
}

func (s *CSVWriterTestSuite) TestWrite() {
	opts := &CSVOpts{
		Comma: ';',
	}
	w := NewCSVWriter(opts)
	d, err := newTestDataset()
	s.Nil(err)
	out, err := newTestWrite(d, w)
	expected := `First name;Last name;Age
Julia;Roberts;40
John;Malkovich;42
`

	s.Nil(err)
	s.Equal(expected, out)
}

func (s *CSVWriterTestSuite) TestWriteComma() {
	opts := &CSVOpts{
		Comma: ',',
	}
	w := NewCSVWriter(opts)
	d, err := newTestDataset()
	s.Nil(err)
	out, err := newTestWrite(d, w)
	expected := `First name,Last name,Age
Julia,Roberts,40
John,Malkovich,42
`

	s.Nil(err)
	s.Equal(expected, out)
}

func (s *CSVWriterTestSuite) TestWriteCRLF() {
	opts := &CSVOpts{
		Comma:   ',',
		UseCRLF: true,
	}
	w := NewCSVWriter(opts)
	d, err := newTestDataset()
	s.Nil(err)
	out, err := newTestWrite(d, w)
	expected := "First name,Last name,Age\r\nJulia,Roberts,40\r\nJohn,Malkovich,42\r\n"

	s.Nil(err)
	s.Equal(expected, out)
}

func TestCSVWriterTestSuite(t *testing.T) {
	suite.Run(t, new(CSVWriterTestSuite))
}
