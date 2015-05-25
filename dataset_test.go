package tabular

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	testHeaders = []struct {
		Key   string
		Title string
	}{
		{
			Key:   "name",
			Title: "First name",
		},
		{
			Key:   "surname",
			Title: "Last name",
		},
		{
			Key:   "age",
			Title: "Age",
		},
	}
	testRows = [][]string{
		{"Julia", "Roberts", "40"},
		{"John", "Malkovich", "42"},
	}
)

func newTestDataset() (*Dataset, error) {
	d := NewDataSet()
	for _, hdr := range testHeaders {
		d.AddHeader(hdr.Key, hdr.Title)
	}
	for _, row := range testRows {
		r := NewRowFromSlice(row)
		err := d.Append(r)
		if err != nil {
			return nil, err
		}
	}
	return d, nil
}

func newTestWrite(d *Dataset, w Writer) (string, error) {
	var buf bytes.Buffer
	bufw := bufio.NewWriter(&buf)
	err := d.Write(w, bufw)
	if err != nil {
		return "", err
	}
	err = bufw.Flush()
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

type DatasetTestSuite struct {
	suite.Suite
}

func (s *DatasetTestSuite) TestRowWidthWithHeaders() {
	d := NewDataSet()
	d.AddHeader("name", "Name")

	r1 := NewRow("john")
	err := d.Append(r1)
	s.Nil(err)

	r2 := NewRow("julia", "mitchell")
	err = d.Append(r2)
	s.Error(err)
	s.True(d.HasHeaders())
}

func (s *DatasetTestSuite) TestRowWidthWithoutHeaders() {
	d := NewDataSet()

	r1 := NewRow("julia", "mitchell")
	err := d.Append(r1)
	s.Nil(err)

	r2 := NewRow("john")
	err = d.Append(r2)
	s.Error(err)
	s.False(d.HasHeaders())
}

func (s *DatasetTestSuite) TestSort() {
	d := NewDataSet()
	d.AddHeader("name", "Name")
	d.AddHeader("surname", "Surname")

	r1 := NewRow("julia", "mitchell")
	r2 := NewRow("martin", "brown")
	r3 := NewRow("peter", "kafka")

	d.Append(r1, r2, r3)
	d.Sort("name", false)

	e1, _ := d.Get(0)
	e2, _ := d.Get(1)
	e3, _ := d.Get(2)

	s.Equal(e1, r1)
	s.Equal(e2, r2)
	s.Equal(e3, r3)
}

func (s *DatasetTestSuite) TestSortReverse() {
	d := NewDataSet()
	d.AddHeader("name", "Name")
	d.AddHeader("surname", "Surname")

	r1 := NewRow("julia", "mitchell")
	r2 := NewRow("martin", "brown")
	r3 := NewRow("peter", "kafka")

	d.Append(r1, r2, r3)
	d.Sort("name", true)

	e1, _ := d.Get(0)
	e2, _ := d.Get(1)
	e3, _ := d.Get(2)

	s.Equal(e1, r3)
	s.Equal(e2, r2)
	s.Equal(e3, r1)
}

func (s *DatasetTestSuite) TestHasColumns() {
	d := NewDataSet()
	d.AddHeader("name", "Name")
	d.AddHeader("surname", "Surname")

	s.True(d.HasCol("name"))
	s.True(d.HasCol("surname"))
	s.False(d.HasCol("not"))
	s.Equal(2, d.HeaderCount())
}

func (s *DatasetTestSuite) TestColValues() {
	d := NewDataSet()
	d.AddHeader("name", "Name")
	d.AddHeader("surname", "Surname")

	r1 := NewRow("julia", "mitchell")
	r2 := NewRow("martin", "brown")
	r3 := NewRow("peter", "kafka")
	d.Append(r1, r2, r3)

	s.Equal([]string{"julia", "martin", "peter"}, d.GetColValues("name"))
	s.Equal([]string{"mitchell", "brown", "kafka"}, d.GetColValues("surname"))
}

func (s *DatasetTestSuite) TestColWidth() {
	d := NewDataSet()
	d.AddHeader("name", "Name")
	d.AddHeader("surname", "Surname")

	r1 := NewRow("julia", "mitchell")
	r2 := NewRow("martin", "brown")
	r3 := NewRow("peter", "kafka")
	d.Append(r1, r2, r3)

	s.Equal(6, d.GetKeyWidth("name"))
	s.Equal(8, d.GetKeyWidth("surname"))
	s.Equal(0, d.GetKeyWidth("not"))

	s.Equal(6, d.GetIdxWidth(0))
	s.Equal(8, d.GetIdxWidth(1))
	s.Equal(0, d.GetIdxWidth(23))
}

func TestDatasetTestSuite(t *testing.T) {
	suite.Run(t, new(DatasetTestSuite))
}
