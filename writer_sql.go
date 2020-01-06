package tabular

import (
	"database/sql"
	"io"

	"github.com/Masterminds/squirrel"
)

// SQLOpts represents options passed to the SQL writer.
type SQLOpts struct {
	DB         *sql.DB
	Driver     string
	Table      string
	ColMapping map[string]string
}

// NewSQLWriter creates a new SQL dataset writer.
func NewSQLWriter(opts *SQLOpts) *SQLWriter {
	w := &SQLWriter{opts}
	return w
}

// SQLWriter represents a SQL dataset writer.
type SQLWriter struct {
	opts *SQLOpts
}

// Name returns name of the writer.
func (sw *SQLWriter) Name() string {
	return "sql"
}

// NeedsHeaders returns true if headers are required.
func (sw *SQLWriter) NeedsHeaders() bool {
	return true
}

// Write writes dataset to writer.
func (sw *SQLWriter) Write(d *Dataset, w io.Writer) error {
	wr := newSQLTableWriter(d, sw.opts)
	return wr.write()
}

func newSQLTableWriter(d *Dataset, opts *SQLOpts) *sqlTableWriter {
	return &sqlTableWriter{
		d:    d,
		opts: opts,
	}
}

type sqlTableWriter struct {
	d    *Dataset
	opts *SQLOpts
}

func (stw *sqlTableWriter) placeholder() squirrel.PlaceholderFormat {
	if stw.opts.Driver == "postgres" || stw.opts.Driver == "postgresql" {
		return squirrel.Dollar
	}
	return squirrel.Question
}

func (stw *sqlTableWriter) cols() []string {
	cols := make([]string, 0, stw.d.HeaderCount())
	for _, hdr := range stw.d.Headers() {
		if v, ok := stw.opts.ColMapping[hdr.Key]; ok {
			cols = append(cols, v)
		} else {
			cols = append(cols, hdr.Key)
		}
	}
	return cols
}

func (stw *sqlTableWriter) vals(row *Row) []interface{} {
	res := make([]interface{}, 0, row.Len())
	for _, item := range row.Items() {
		res = append(res, item)
	}
	return res
}

func (stw *sqlTableWriter) query(tx *sql.Tx, row *Row) (sql.Result, error) {
	return squirrel.
		Insert(stw.opts.Table).
		Columns(stw.cols()...).
		Values(stw.vals(row)...).
		RunWith(tx).
		PlaceholderFormat(stw.placeholder()).
		Exec()
}

func (stw *sqlTableWriter) write() error {
	tx, err := stw.opts.DB.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	for _, row := range stw.d.Rows() {
		if _, err = stw.query(tx, row); err != nil {
			return err
		}
	}

	return err
}
