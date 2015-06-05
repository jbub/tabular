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

func (s *LatexWriterTestSuite) TestWriteTabularX() {
	opts := &LatexOpts{
		TabularX: true,
	}
	w := NewLatexWriter(opts)
	d, err := newTestDataset()
	s.Nil(err)
	out, err := newTestWrite(d, w)
	expected :=
		`\begin{table}[h]
\begin{tabularx}{|X|X|X|}
\hline
First name & Last name & Age \\ \hline
Julia      & Roberts   & 40  \\ \hline
John       & Malkovich & 42  \\ \hline
\end{tabularx}
\end{table}
`

	s.Nil(err)
	s.Equal(expected, out)
}

func (s *LatexWriterTestSuite) TestWriteCaption() {
	opts := &LatexOpts{
		Caption: "Test caption",
	}
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
\caption{Test caption}
\end{table}
`

	s.Nil(err)
	s.Equal(expected, out)
}

func (s *LatexWriterTestSuite) TestWriteCaptionEscape() {
	opts := &LatexOpts{
		Caption: "Te#st capt&ion",
	}
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
\caption{Te\#st capt\&ion}
\end{table}
`

	s.Nil(err)
	s.Equal(expected, out)
}

func (s *LatexWriterTestSuite) TestWriteCenter() {
	opts := &LatexOpts{
		Center: true,
	}
	w := NewLatexWriter(opts)
	d, err := newTestDataset()
	s.Nil(err)
	out, err := newTestWrite(d, w)
	expected :=
		`\begin{table}[h]
\centering
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

func (s *LatexWriterTestSuite) TestWriteCenterTabularX() {
	opts := &LatexOpts{
		Center:   true,
		TabularX: true,
	}
	w := NewLatexWriter(opts)
	d, err := newTestDataset()
	s.Nil(err)
	out, err := newTestWrite(d, w)
	expected :=
		`\begin{table}[h]
\begin{tabularx}{|X|X|X|}
\hline
First name & Last name & Age \\ \hline
Julia      & Roberts   & 40  \\ \hline
John       & Malkovich & 42  \\ \hline
\end{tabularx}
\end{table}
`

	s.Nil(err)
	s.Equal(expected, out)
}

func TestLatexWriterTestSuite(t *testing.T) {
	suite.Run(t, new(LatexWriterTestSuite))
}
