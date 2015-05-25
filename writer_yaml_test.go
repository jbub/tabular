package tabular

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type YAMLWriterTestSuite struct {
	suite.Suite
}

func (s *YAMLWriterTestSuite) TestWrite() {
	opts := &YAMLOpts{}
	w := NewYAMLWriter(opts)
	d, err := newTestDataset()
	s.Nil(err)
	out, err := newTestWrite(d, w)
	expected := `- name: Julia
  surname: Roberts
  age: 40
- name: John
  surname: Malkovich
  age: 42
`

	s.Nil(err)
	s.Equal(expected, out)
}

func TestYAMLWriterTestSuite(t *testing.T) {
	suite.Run(t, new(YAMLWriterTestSuite))
}
