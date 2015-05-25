package tabular

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type LatexWriterTestSuite struct {
	suite.Suite
}

func (s *LatexWriterTestSuite) TestWrite() {
	opts := &LatexOpts{}
	w := NewLatexWriter(opts)
	d, err := newTestDataset()
	s.Nil(err)
	out, err := newTestWrite(d, w)
	expected :=
		`\begin{table}[h]
\begin{tabular}{|l|l|l|}
\hline
First name & Last name & Age \\ \hline
Julia      & Roberts   & 40  \\ \hline
John       & Malkovich & 42  \\ \hline
\end{tabular}
\end{table}
`

	s.Nil(err)
	s.Equal(expected, out)
}

func TestLatexWriterTestSuite(t *testing.T) {
	suite.Run(t, new(LatexWriterTestSuite))
}
