package tabular

import (
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

type SQLWriterTestSuite struct {
	suite.Suite
}

func (s *SQLWriterTestSuite) TestWrite() {
	s.writeTest("actors", nil)
}

func (s *SQLWriterTestSuite) TestWriteMapping() {
	mapping := map[string]string{
		"age":     "renamed_age",
		"invalid": "column",
	}
	table := "actors"
	s.writeTest(table, mapping)
}

func (s *SQLWriterTestSuite) writeTest(table string, mapping map[string]string) {
	db, mock, err := sqlmock.New()
	s.NoError(err)
	defer db.Close()

	opts := &SQLOpts{
		DB:    db,
		Table: table,
	}
	if mapping != nil {
		opts.ColMapping = mapping
	}

	w := NewSQLWriter(opts)
	d, err := newTestDataset()
	s.NoError(err)

	mock.ExpectBegin()

	for _, row := range d.Rows() {
		var vals []driver.Value
		for _, v := range row.Items() {
			vals = append(vals, v)
		}

		mock.ExpectExec("INSERT INTO " + table).
			WithArgs(vals...).
			WillReturnResult(sqlmock.NewResult(1, 1))
	}

	mock.ExpectCommit()

	_, err = newTestWrite(d, w)
	s.NoError(err)
}

func TestSQLWriterTestSuite(t *testing.T) {
	suite.Run(t, new(SQLWriterTestSuite))
}
