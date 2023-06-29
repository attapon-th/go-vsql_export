package pkg

import (
	"database/sql"
	"errors"
	"io"

	"github.com/joho/sqltocsv"
)

const (
	// TimeFormat is the default time format for the logger.
	TimeFormat = "2006-01-02 15:04:05.999"
)

// ToCsv convert sql.Rows to csv
func ToCsv(rows *sql.Rows, w io.Writer) error {
	if rows == nil {
		return errors.New("rows is nil")
	}
	if w == nil {
		return errors.New("writer is nil")
	}

	stoc := sqltocsv.New(rows)
	stoc.WriteHeaders = true
	stoc.TimeFormat = TimeFormat
	return stoc.Write(w)
}
